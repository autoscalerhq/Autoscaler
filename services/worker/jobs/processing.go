package jobs

import (
	"context"
	"github.com/autoscalerhq/autoscaler/internal/bootstrap"
	"github.com/nats-io/nats.go"
	"github.com/nats-io/nats.go/jetstream"
	"strings"
	"time"
)

func StartConsumer(ctx context.Context, subject string) {

	nc, err := bootstrap.GetNatsConn()

	if err != nil {
		println("Failed to connect to NATS: %v", err)
	}

	sub, err := nc.QueueSubscribe("job-queue", "job-worker-group",
		func(m *nats.Msg) {
			// Process each message asynchronously
			processJob(m.Subject, m.Data)

			// Acknowledge the message
			if err := m.Ack(); err != nil {
				println("Failed to acknowledge message: %v", err)
			}
		},
	)

	if err != nil {
		println("Failed to subscribe to subject %s: %v", subject, err)
	}

	defer func(sub *nats.Subscription) {
		err := sub.Unsubscribe()
		if err != nil {
			println("Failed to unsubscribe from subject %s: %v", subject, err)
		}
	}(sub)

	println("Started consumer for subject: %s", subject)

	// Wait until context is canceled
	<-ctx.Done()
}

func processJob(subject string, data []byte) {
	// Extract a job type from a subject
	jobType := strings.TrimPrefix(subject, "jobs.")

	// Implement job processing logic based on a job type
	switch jobType {
	case "init":
		AddIdempotency(subject, data)
	case "client":
		SyncClientJob()
	case "sync.aws":
		fetchAWSPricing()
	case "sync.azure":
		fetchAzurePricing()
	case "sync.gcp":
		fetchGCPPricing()
	default:
		println("Unknown job type: %s", jobType)
	}
}

func AddIdempotency(subject string, data []byte) {
	idempotencykv, _, err := bootstrap.GetKVStore(jetstream.KeyValueConfig{
		Bucket:   "idempotency",
		TTL:      time.Hour * 6,
		MaxBytes: 1024 * 1024,
	})

	if err != nil {
		println("Failed to get idempotency kv store: %v", err)
	}

	idkey, err := GetIdempotencyKey(idempotencykv)

	if err != nil {
		println("Failed to get idempotency key: %v", err)
	}

	newData, newSubject, err := addIdempotencyKey(data, idkey)

	if err != nil {
		println("Failed to add idempotency key: %v", err)
	}

	js, err := bootstrap.GetJetStream()

	if err != nil {
		println("Failed to connect to NATS Jetstream: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	// Publish the message to the JetStream subject.
	// Send the message to NATS with the message as the subject
	if _, err := js.Publish(ctx, newSubject, newData); err != nil {
		println("error publishing", err)
	}
}

func PipeSubjectsToJetstream() {
	nc, err := bootstrap.GetNatsConn()

	if err != nil {
		println("Failed to connect to NATS: %v", err)
	}

	js, err := bootstrap.GetJetStream()

	if err != nil {
		println("Failed to connect to NATS Jetstream: %v", err)
	}

	// Define the standard NATS subject we're subscribed to.
	standardSubject := "jobs"
	// Define the JetStream subject.
	jetStreamSubject := "jobs.init"

	// Subscribe to the standard NATS subject.
	_, err = nc.Subscribe(standardSubject, func(msg *nats.Msg) {
		println("Received a message on %s: %s", standardSubject, string(msg.Data))

		ctx := context.Background()
		// Publish the message to the JetStream subject.
		_, err := js.Publish(ctx, jetStreamSubject, msg.Data)
		if err != nil {
			println("Failed to publish to JetStream: %v", err)
		} else {
			println("Message republished to JetStream subject: %s", jetStreamSubject)
		}
	})
	if err != nil {
		println("Failed to subscribe to subject %s: %v", standardSubject, err)
	}
}

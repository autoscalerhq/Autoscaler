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

	case "sync.azure":

	case "sync.gcp":

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

	idkey, err := GetIdempotencyKey(&idempotencykv)

	if err != nil {
		println("Failed to get idempotency key: %v", err)
	}

	newData, err := addIdempotencyKey(data, idkey)

	if err != nil {
		println("Failed to add idempotency key: %v", err)
	}

	nc, err := bootstrap.GetNatsConn()

	if err != nil {
		println("Failed to connect to NATS: %v", err)
	}

	// Send the message to NATS with the message as the subject
	if err := nc.Publish(string(data), newData); err != nil {
		println("error publishing", err)
	}
}

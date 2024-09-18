package jobs

import (
	"fmt"
	"github.com/autoscalerhq/autoscaler/internal/bootstrap"
	"time"

	"github.com/nats-io/nats.go"
)

var (
	QUEUES    map[string]nats.StreamConfig
	CONSUMERS map[string]nats.ConsumerConfig
)

func initQueues() {
	QUEUES = map[string]nats.StreamConfig{
		"JOB_STREAM": {
			Name:        "JobStream",
			Description: "A stream for job processing",
			Subjects:    []string{"jobs.*"},
			Retention:   nats.WorkQueuePolicy,
			MaxBytes:    1024 * 1024 * 10, // 10 MB
			Discard:     nats.DiscardOld,
			MaxAge:      24 * time.Hour, // 1 day
			MaxMsgSize:  1024,           // 1 KB
			Storage:     nats.FileStorage,
			Replicas:    1,
			NoAck:       false,
			Duplicates:  2 * time.Minute,
		},
	}

	CONSUMERS = map[string]nats.ConsumerConfig{
		"JOB_CONSUMER": {
			Durable:        "JobConsumer",
			DeliverSubject: "job-queue",
			DeliverGroup:   "job-worker-group",
			AckPolicy:      nats.AckExplicitPolicy,
			MaxAckPending:  1000,
			ReplayPolicy:   nats.ReplayInstantPolicy,
		},
	}
}

func CreateQueues() error {

	initQueues()

	nc, err := bootstrap.GetNatsConn()

	if err != nil {
		println("Failed to connect to NATS: %v", err)
	}

	// Access JetStream context
	js, err := nc.JetStream()
	if err != nil {
		return fmt.Errorf("error accessing JetStream context: %w", err)
	}

	for name, config := range QUEUES {
		println("Creating or keeping queue:", name)
		_, err := js.StreamInfo(config.Name)
		if err != nil {
			_, err = js.AddStream(&config)
			if err != nil {
				fmt.Println("Error adding stream:", err)
			} else {
				fmt.Println("Stream created:", config.Name)
			}
		} else {
			fmt.Println("Stream exists:", config.Name)
		}

		// Create consumers for the stream
		for cname, cconfig := range CONSUMERS {
			println("Creating consumer:", cname)
			cconfig.FilterSubject = "jobs.*" // Ensure the consumer is filtering the correct subject

			if _, err := js.AddConsumer(config.Name, &cconfig); err != nil {
				fmt.Println("Could not add consumer:", err)
			} else {
				fmt.Println("Consumer created for stream:", config.Name)
			}
		}
	}

	return nil
}

func InitializeAppJobs() {
	client := bootstrap.GetDkronClient()
	err := CreateClientSyncCron(client)
	if err != nil {
		println("Error creating client sync cron")
	}
	err = CreatePricePullCron(client)
	if err != nil {
		println("Error creating price pull cron")
	}

	err = CreateQueues()
	if err != nil {
		println("Error creating queues ")
	}
}

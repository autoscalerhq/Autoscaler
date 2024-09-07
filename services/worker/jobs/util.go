package jobs

import (
	"context"
	"errors"
	"github.com/autoscalerhq/autoscaler/lib/dkron"
	"github.com/nats-io/nats.go"
	"github.com/nats-io/nats.go/jetstream"
	"log"
)

type Job func() error

// IdempotentJobWrapper is the higher-order function to create idempotent jobs
func IdempotentJobWrapper(job Job, kvStore jetstream.KeyValue) func(ctx context.Context, msg *nats.Msg) {
	return func(ctx context.Context, msg *nats.Msg) {
		// Extract the idempotency key from NATS headers
		idempotencyKey := msg.Header.Get("Idempotency-Key")
		if idempotencyKey == "" {
			log.Println("No Idempotency-Key header provided")
			return
		}

		// Check if the key already exists in the KV store
		_, err := kvStore.Get(ctx, idempotencyKey)

		if errors.Is(err, jetstream.ErrKeyNotFound) {
			log.Printf("Idempotency-Key %s has already been processed", idempotencyKey)
			return
		}

		// Execute the job
		if err := job(); err != nil {
			log.Printf("Job execution failed: %v", err)
			return
		}

		// Log the Idempotency-Key in the KV store
		_, err = kvStore.Put(ctx, idempotencyKey, []byte("done"))

		if err != nil {
			log.Printf("Failed to set Idempotency-Key %s in KV store: %v", idempotencyKey, err)
		}
	}
}

func DeepCopyJob(j *dkron.Job) *dkron.Job {
	newJob := &dkron.Job{
		Name:           j.Name,
		DisplayName:    j.DisplayName,
		Owner:          j.Owner,
		Retries:        j.Retries,
		Schedule:       j.Schedule,
		Executor:       j.Executor,
		Tags:           make(map[string]string),
		ExecutorConfig: make(map[string]string),
	}

	// Deep copying Tags map
	for key, value := range j.Tags {
		newJob.Tags[key] = value
	}

	// Deep copying ExecutorConfig map
	for key, value := range j.ExecutorConfig {
		newJob.ExecutorConfig[key] = value
	}

	return newJob
}

package jobs

import (
	"fmt"
	"github.com/autoscalerhq/autoscaler/lib/dkron"
	"github.com/nats-io/nats.go"
	"time"
)

var QUEUES map[string]nats.StreamConfig = nil

func initQueues() {
	QUEUES = map[string]nats.StreamConfig{
		"JOB_STREAM": {
			Name:        "JobStream",
			Description: "A stream for job processing",
			Subjects:    []string{"jobs.*"},
			Retention:   0, // Customize as per retention policy
			//MaxConsumers: 10,
			//MaxMsgs:      1000,
			MaxBytes:   1024 * 1024 * 10, // 10 MB
			Discard:    0,                // Customize as per discard policy
			MaxAge:     24 * time.Hour,   // 1 day
			MaxMsgSize: 1024,             // 1 KB
			Storage:    0,                // Customize as per storage type
			Replicas:   1,
			NoAck:      false,
			Duplicates: 2 * time.Minute,
			// Add other fields as needed
		},
	}
}

func CreateQueues(nc *nats.Conn) error {
	initQueues()
	// Access JetStream context
	js, err := nc.JetStream()
	if err != nil {
		return fmt.Errorf("error accessing JetStream context: %w", err)
	}

	for name, config := range QUEUES {
		println("Creating or keeping queue:", name)
		_, err := js.StreamInfo(config.Name)
		if err != nil {
			_, err = js.AddStream(&nats.StreamConfig{
				Name:                 config.Name,
				Description:          config.Description,
				Subjects:             config.Subjects,
				Retention:            nats.RetentionPolicy(config.Retention),
				MaxConsumers:         config.MaxConsumers,
				MaxMsgs:              config.MaxMsgs,
				MaxBytes:             config.MaxBytes,
				Discard:              nats.DiscardPolicy(config.Discard),
				DiscardNewPerSubject: config.DiscardNewPerSubject,
				MaxAge:               config.MaxAge,
				MaxMsgsPerSubject:    config.MaxMsgsPerSubject,
				MaxMsgSize:           config.MaxMsgSize,
				Storage:              nats.StorageType(config.Storage),
				Replicas:             config.Replicas,
				NoAck:                config.NoAck,
				Duplicates:           config.Duplicates,
				Placement:            config.Placement,
				Mirror:               config.Mirror,
				Sources:              config.Sources,
				Sealed:               config.Sealed,
				DenyDelete:           config.DenyDelete,
				DenyPurge:            config.DenyPurge,
				AllowRollup:          config.AllowRollup,
				Compression:          nats.StoreCompression(config.Compression),
				FirstSeq:             config.FirstSeq,
				SubjectTransform:     config.SubjectTransform,
				RePublish:            config.RePublish,
				AllowDirect:          config.AllowDirect,
				MirrorDirect:         config.MirrorDirect,
				ConsumerLimits:       config.ConsumerLimits,
				Metadata:             config.Metadata,
				Template:             config.Template,
			})
			if err != nil {
				fmt.Println("Error adding stream:", err)
			} else {
				fmt.Println("Stream created:", config.Name)
			}
		} else {
			fmt.Println("Stream exists:", config.Name)
		}
	}

	return nil
}

func CreateJobs(client *dkron.Client) {
	CreateClientSyncCron(client)
	CreatePricePullCron(client)
}

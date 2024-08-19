package natutils

import (
	"context"
	"errors"
	"fmt"
	"github.com/nats-io/nats.go"
	"github.com/nats-io/nats.go/jetstream"
	"os"
)

func GetNatsConn() (*nats.Conn, error) {
	url := os.Getenv("NATS_URL")
	if url == "" {
		url = nats.DefaultURL
	}
	nc, err := nats.Connect(url)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to nats: %w", err)
	}

	return nc, nil
}

func NewJetStream() (jetstream.JetStream, error) {
	nc, err := GetNatsConn()
	js, err := jetstream.New(nc)
	if err != nil {
		return nil, fmt.Errorf("failed to create new jetstream: %w", err)
	}
	return js, nil
}

// NewKeyValueStore creates a new key value store in JetStream if it doesn't exist
// otherwise it returns the existing one.
func NewKeyValueStore(config jetstream.KeyValueConfig) (jetstream.KeyValue, context.Context, error) {
	js, err := NewJetStream()
	ctx := context.Background()
	if err != nil {
		return nil, nil, fmt.Errorf("error creating jetstream: %w", err)
	}
	kv, err := js.CreateKeyValue(ctx, config)
	if err != nil {
		if !errors.Is(err, jetstream.ErrBucketExists) {
			return nil, nil, fmt.Errorf("error creating key value store: %w", err)
		}
		kv, err = js.KeyValue(ctx, config.Bucket)
	}
	return kv, ctx, nil
}

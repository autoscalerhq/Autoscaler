package natutils

import (
	"context"
	"errors"
	"fmt"
	"github.com/nats-io/nats.go"
	"github.com/nats-io/nats.go/jetstream"
	"os"
	"time"
)

func GetNatsConn() (*nats.Conn, error) {
	url := os.Getenv("NATS_URL")
	if url == "" {
		url = nats.DefaultURL
	}
	nc, err := nats.Connect(url,
		nats.DisconnectErrHandler(func(nc *nats.Conn, err error) {
			fmt.Printf("Got disconnected! Reason: %q\n", err)
		}),
		nats.ReconnectHandler(func(nc *nats.Conn) {
			fmt.Printf("Got reconnected to %v!\n", nc.ConnectedUrl())
		}),
		nats.ClosedHandler(func(nc *nats.Conn) {
			fmt.Printf("Connection closed. Reason: %q\n", nc.LastError())
		}),
		nats.ReconnectJitter(500*time.Millisecond, 2*time.Second),
		nats.MaxReconnects(5),
		nats.ReconnectWait(2*time.Second),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to nats: %w", err)
	}

	return nc, nil
}

func NewJetStream(nc *nats.Conn, opts ...jetstream.JetStreamOpt) (jetstream.JetStream, error) {
	js, err := jetstream.New(nc, opts...)
	if err != nil {
		return nil, fmt.Errorf("failed to create new jetstream: %w", err)
	}
	return js, nil
}

// NewKeyValueStore creates a new key value store in JetStream if it doesn't exist
// otherwise it returns the existing one.
func NewKeyValueStore(nc *nats.Conn, config jetstream.KeyValueConfig) (jetstream.KeyValue, context.Context, error) {
	ctx := context.Background()
	js, err := NewJetStream(nc)
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

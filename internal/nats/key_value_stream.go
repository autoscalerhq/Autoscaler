package nats

import (
	"context"
	"github.com/nats-io/nats.go"
	"github.com/nats-io/nats.go/jetstream"
	"os"
	"time"
)

func NewKeyValueStream(config jetstream.KeyValueConfig) (jetstream.KeyValue, context.Context) {
	url := os.Getenv("NATS_URL")
	if url == "" {
		url = nats.DefaultURL
	}
	nc, _ := nats.Connect(url)
	defer func(nc *nats.Conn) {
		err := nc.Drain()
		if err != nil {
		}
	}(nc)
	js, _ := jetstream.New(nc)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	kv, _ := js.CreateKeyValue(ctx, config)

	return kv, ctx
}

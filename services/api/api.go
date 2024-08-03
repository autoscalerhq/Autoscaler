package main

import (
	"context"
	"fmt"
	"github.com/nats-io/nats.go"
	"github.com/nats-io/nats.go/jetstream"
	"time"
)

func main() {

	var servers = "nats://localhost:4222, nats://localhost:4223"

	nc, err := nats.Connect(servers,
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
		println("nats err", err.Error())
		panic("Failed to connect to NATS")
	}

	defer func(nc *nats.Conn) {
		err := nc.Drain()
		if err != nil {
			println(err.Error(), "Nats drain err")
		}
	}(nc)

	js, err := jetstream.New(nc)
	if err != nil {
		println("JetStream err", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	kv, _ := js.CreateKeyValue(ctx, jetstream.KeyValueConfig{
		Bucket: "profiles",
	})

	kv.Put(ctx, "sue.color", []byte("blue"))
	entry, _ := kv.Get(ctx, "sue.color")
	fmt.Printf("%s @ %d -> %q\n", entry.Key(), entry.Revision(), string(entry.Value()))

}

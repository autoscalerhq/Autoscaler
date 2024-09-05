package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/autoscalerhq/autoscaler/lib/dkron"
	"github.com/nats-io/nats.go"
	"github.com/nats-io/nats.go/jetstream"
	"time"
)

func main() {

	// Create a new Dkron client
	client := dkron.NewClient("http://localhost:8080/v1")

	name, err := client.ShowJobByName("example_job_nats")
	if err != nil {
		fmt.Println("Error showing job:", err)
		return
	}

	fmt.Println("Job name", name)

	// Example job definition
	job2 := dkron.Job{
		Name:     "example_job_nats",
		Schedule: "@every 1s",
		Executor: "nats",
		ExecutorConfig: map[string]string{
			"url":     "nats://host.docker.internal:4222",
			"message": "job id",
			"subject": "events.cron",
		},
	}

	// Create or update the job
	createdJob2, err := client.CreateOrUpdateJob(job2, false)
	if err != nil {
		fmt.Println("Error creating or updating job:", err)
		return
	}
	fmt.Println("Created/Updated Job:", createdJob2)

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
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Minute)
	defer cancel()

	// Getting work jobs
	cfg := jetstream.StreamConfig{
		Name:      "EVENTS",
		Retention: jetstream.WorkQueuePolicy,
		Subjects:  []string{"events.>"},
	}

	ctx, cancel = context.WithTimeout(context.Background(), 2*time.Minute)
	defer cancel()

	stream, err := js.CreateOrUpdateStream(ctx, cfg)
	if err != nil {
		println("Create stream err", err.Error())
		return
	}
	fmt.Println("created the stream")
	printStreamState(ctx, stream)

	cons, err := stream.CreateOrUpdateConsumer(ctx, jetstream.ConsumerConfig{
		Name: "processor-3",
	})

	if err != nil {
		println("Create Consumer err", err.Error())
		return
	}

	var messageHandler jetstream.MessageHandler = func(msg jetstream.Msg) {
		println("Msg: ", string(msg.Data()))
		println("Msg sub: ", msg.Subject())
		println("Msg headers: ", msg.Headers().Values(""))
		err := msg.Ack()
		if err != nil {
			println(err.Error(), "ack err")
		}
	}

	consumeCtx, err := cons.Consume(messageHandler, jetstream.PullMaxMessages(10))
	defer consumeCtx.Drain()

	if err != nil {
		println("Consuming messages err", err.Error())
		return
	}

	println("Pausing main thread")

	fmt.Println("\n# Stream info with one consumer")
	printStreamState(ctx, stream)
}

func printStreamState(ctx context.Context, stream jetstream.Stream) {
	info, err := stream.Info(ctx)
	if err != nil {
		println(err.Error(), "Cant get stream info")
		return
	}

	b, err := json.MarshalIndent(info.State, "", " ")
	if err != nil {
		println(err.Error(), "could not marshal state info")
		return
	}
	fmt.Println(string(b))
}

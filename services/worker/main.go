package main

import (
	"fmt"
	natutils "github.com/autoscalerhq/autoscaler/internal/nats"
	"github.com/autoscalerhq/autoscaler/lib/dkron"
	"github.com/autoscalerhq/autoscaler/services/worker/jobs"
	"github.com/joho/godotenv"
	"github.com/nats-io/nats.go"
	"github.com/nats-io/nats.go/jetstream"
	"time"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		// Todo this should be handled better before going to production
		//log.Fatal("Error loading .env file")
	}
	// Envs to get Dkron and Nats

	// Set up the NATS DKronclient first
	// This is to ensure that DKron has something to send to
	nc, err := natutils.GetNatsConn()

	defer func(nc *nats.Conn) {
		err := nc.Drain()
		if err != nil {
			println(err.Error(), "error draining nats connection")
		}
	}(nc)

	//js, err := natutils.NewJetStream(
	//	nc,
	//	jetstream.JetStreamOpt{},
	//)

	ldkv, _, err := natutils.NewKeyValueStore(
		nc,
		jetstream.KeyValueConfig{
			Bucket:   "leadership",
			TTL:      time.Hour * 24,
			MaxBytes: 1024 * 1024,
		},
	)

	if err != nil {
		println(err)
	}

	// Create a new Dkron client
	DKronclient := dkron.NewClient("http://localhost:8080/v1")

	// Create app jobs to keep sync of state
	name, err := DKronclient.ShowJobByName("example_job_nats")
	if err != nil {
		fmt.Println("Error showing job:", err)
		return
	}

	fmt.Println("Job name", name)

	jobs.CreateClientSyncCron(DKronclient)
	jobs.CreatePricePullCron(DKronclient)

	//ctx, cancel := context.WithTimeout(context.Background(), 2*time.Minute)
	//defer cancel()

	//// Getting work jobs
	//cfg := jetstream.StreamConfig{
	//	Name:      "EVENTS",
	//	Retention: jetstream.WorkQueuePolicy,
	//	Subjects:  []string{"events.>"},
	//}
	//
	//ctx, cancel = context.WithTimeout(context.Background(), 2*time.Minute)
	//defer cancel()
	//
	//stream, err := js.CreateOrUpdateStream(ctx, cfg)
	//if err != nil {
	//	println("Create stream err", err.Error())
	//	return
	//}
	//fmt.Println("created the stream")
	//printStreamState(ctx, stream)
	//
	//cons, err := stream.CreateOrUpdateConsumer(ctx, jetstream.ConsumerConfig{
	//	Name: "processor-3",
	//})
	//
	//if err != nil {
	//	println("Create Consumer err", err.Error())
	//	return
	//}
	//
	//var messageHandler jetstream.MessageHandler = func(msg jetstream.Msg) {
	//	println("Msg: ", string(msg.Data()))
	//	println("Msg sub: ", msg.Subject())
	//	println("Msg headers: ", msg.Headers().Values(""))
	//	err := msg.Ack()
	//	if err != nil {
	//		println(err.Error(), "ack err")
	//	}
	//}
	//
	//consumeCtx, err := cons.Consume(messageHandler, jetstream.PullMaxMessages(10))
	//defer consumeCtx.Drain()
	//
	//if err != nil {
	//	println("Consuming messages err", err.Error())
	//	return
	//}
	//
	//println("Pausing main thread")
	//
	//fmt.Println("\n# Stream info with one consumer")
	//printStreamState(ctx, stream)
}

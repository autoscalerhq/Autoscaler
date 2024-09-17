package main

import (
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

	err = jobs.CreateQueues(nc)
	if err != nil {
		return
	}

	clusterkv, _, err := natutils.NewKeyValueStore(
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
	jobs.CreateJobs(DKronclient)

}

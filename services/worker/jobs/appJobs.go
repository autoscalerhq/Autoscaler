package jobs

import (
	"fmt"
	"github.com/autoscalerhq/autoscaler/lib/dkron"
)

func SyncClientJob() {
	// TODO
	// get every job in the database in batch
	// per batch put into a stream
	// workers then resync to dkron

	println("TODO Syncing Client Jobs")
}

func CreateClientSyncCron(client *dkron.Client) error {

	retries := 1
	job := dkron.Job{
		Name:        "crit_client_job_sync",
		DisplayName: "CRIT Client Job Sync",
		Ephemeral:   false,
		Owner:       "Autoscaler",
		Retries:     &retries,
		Tags: map[string]string{
			"critical": "true",
		},
		Schedule:    "0 */15 * * * *",
		Executor:    "nats",
		Concurrency: "forbid",
		ExecutorConfig: map[string]string{
			"url":     "nats://host.docker.internal:4222",
			"message": "{\"newSubject\": \"job.client\"}",
			"subject": "job.init",
		},
	}

	createdJob, err := client.CreateOrUpdateJob(job, true)

	if err != nil {
		return err
	}

	// TODO replace with logging system
	fmt.Println("Created/Updated Job:", createdJob)

	return nil
}

func CreatePricePullCron(client *dkron.Client) error {

	var clouds = []string{"aws", "azure", "gcp"}

	retries := 1
	job := dkron.Job{
		Name:        "crit_price_sync",
		DisplayName: "CRIT Price Sync",
		Owner:       "Autoscaler",
		Retries:     &retries,
		Tags: map[string]string{
			"critical": "true",
		},
		Schedule:    "@daily",
		Executor:    "nats",
		Concurrency: "forbid",
		ExecutorConfig: map[string]string{
			"url":     "nats://host.docker.internal:4222",
			"subject": "job.init",
		},
	}

	for _, cloud := range clouds {

		sjob := DeepCopyJob(&job)

		sjob.Name = fmt.Sprintf("crit_price_sync_%s", cloud)
		sjob.DisplayName = fmt.Sprintf("CRIT Price Sync %s", cloud)
		sjob.ExecutorConfig["message"] = fmt.Sprintf("\"{\\\"newSubject\\\": \\\"job.price.%s\\\"}\",", cloud)

		// Create or update the job
		createdjob, err := client.CreateOrUpdateJob(*sjob, true)

		// FIXME should continue and report error
		if err != nil {
			return err
		}

		// TODO replace with logging system
		fmt.Println("Created/Updated Job:", createdjob)
	}
	return nil
}

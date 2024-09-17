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

func CreateClientSyncCron(client *dkron.Client) {

	job := dkron.Job{
		Name:        "CRIT_client_job_sync",
		DisplayName: "CRIT Client Job Sync",
		Owner:       "Autoscaler",
		Retries:     1,
		Tags: map[string]string{
			"critical": "true",
		},
		Schedule:    "0 */15 * * * *",
		Executor:    "nats",
		Concurrency: "forbid",
		ExecutorConfig: map[string]string{
			"url":     "nats://host.docker.internal:4222",
			"message": "job.client",
			"subject": "job.init",
		},
	}

	// Create or update the job
	createdJob, err := client.CreateOrUpdateJob(job, true)

	if err != nil {
		fmt.Println("Error creating or updating job:", err)
		return
	}

	// TODO replace with logging system
	fmt.Println("Created/Updated Job:", createdJob)
}

func CreatePricePullCron(client *dkron.Client) {

	var clouds = []string{"aws", "azure", "gcp"}

	job := dkron.Job{
		Name:        "CRIT_Price_Sync",
		DisplayName: "CRIT Price Sync",
		Owner:       "Autoscaler",
		Retries:     1,
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

		sjob.Name = fmt.Sprintf("CRIT_Price_Sync_%s", cloud)
		sjob.DisplayName = fmt.Sprintf("CRIT Price Sync %s", cloud)
		sjob.ExecutorConfig["subject"] = fmt.Sprintf("sync.price.%s", cloud)

		// Create or update the job
		createdjob, err := client.CreateOrUpdateJob(*sjob, true)

		if err != nil {
			fmt.Println("Error creating or updating job:", err)
			return
		}

		// TODO replace with logging system
		fmt.Println("Created/Updated Job:", createdjob)
	}
}

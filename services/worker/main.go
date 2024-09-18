package main

import (
	"context"
	"github.com/autoscalerhq/autoscaler/services/worker/jobs"
	"github.com/joho/godotenv"
	"strings"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		// Todo this should be handled better before going to production
		//log.Fatal("Error loading .env file")
	}

	jobs.InitializeAppJobs()

	// TODO Make this based on a env/refreshable varible
	// TODO make this able to be configured by env
	// Determine the subjects to subscribe to
	// Set up with actual env config
	subjects := parseJobTypes("*")
	println("Subscribing to subjects: %v", subjects)

	ctx := context.Background()

	// Set up subscriptions for each subject
	for _, subject := range subjects {
		go jobs.StartConsumer(ctx, subject)
	}

	// Wait indefinitely
	select {}

}

func parseJobTypes(jobTypes string) []string {
	if jobTypes == "*" {
		return []string{"jobs.>"}
	}

	types := strings.Split(jobTypes, ",")
	subjects := make([]string, len(types))
	for i, t := range types {
		t = strings.TrimSpace(t)
		subjects[i] = "jobs." + t
	}
	return subjects
}

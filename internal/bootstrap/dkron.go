package bootstrap

import (
	"errors"
	"fmt"
	"github.com/autoscalerhq/autoscaler/lib/dkron"
	"sync"
)

var dlock = &sync.Mutex{}
var dkronClient *dkron.Client

// GetDkronClient returns a singleton instance of a Dkron client
func GetDkronClient() (*dkron.Client, error) {

	if shuttingDown {
		return nil, errors.New("sys shutdown")
	}

	if dkronClient == nil {
		dlock.Lock()
		defer dlock.Unlock()
		if dkronClient == nil {
			fmt.Println("Creating single Dkron client instance now.")
			// TODO fix url with env manager
			dkronClient = dkron.NewClient("http://localhost:8090/v1")

			_, err := dkronClient.GetStatus()

			if err != nil {
				panic("Could not connect to Dkron server.")
			}

			// Register cleanup function if necessary for session
			RegisterCleanup(func() {
				fmt.Println("Cleanup AWS session resources if needed.")
				// Add any cleanup logic here for the session if required
				dkronClient = nil
			})

		} else {
			fmt.Println("Dkron client instance already created.")
		}
	} else {
		fmt.Println("Dkron client instance already created.")
	}

	return dkronClient, nil
}

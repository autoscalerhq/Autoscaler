package bootstrap

import (
	"fmt"
	"github.com/autoscalerhq/autoscaler/lib/dkron"
	"sync"
)

var lock = &sync.Mutex{}
var dkronClient *dkron.Client

// GetDkronClient returns a singleton instance of a Dkron client
func GetDkronClient() *dkron.Client {
	if dkronClient == nil {
		lock.Lock()
		defer lock.Unlock()
		if dkronClient == nil {
			fmt.Println("Creating single Dkron client instance now.")
			// TODO fix url with env manager
			dkronClient = dkron.NewClient("http://localhost:8080/v1")
		} else {
			fmt.Println("Dkron client instance already created.")
		}
	} else {
		fmt.Println("Dkron client instance already created.")
	}

	return dkronClient
}

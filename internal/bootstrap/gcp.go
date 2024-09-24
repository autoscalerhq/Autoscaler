package bootstrap

import (
	"context"
	"errors"
	"fmt"
	"sync"

	"cloud.google.com/go/storage"
	"google.golang.org/api/option"
)

var gcpLock = &sync.Mutex{}
var gcpClient *storage.Client

// GetGCPClient returns a singleton instance of a GCP storage client
func GetGCPClient() (*storage.Client, error) {

	if shuttingDown {
		return nil, errors.New("sys shutdown")
	}

	if gcpClient == nil {
		gcpLock.Lock()
		defer gcpLock.Unlock()
		if gcpClient == nil {
			fmt.Println("Creating single GCP client instance now.")
			ctx := context.Background()
			client, err := storage.NewClient(ctx, option.WithCredentialsFile("path/to/your/credentials/file.json"))
			if err != nil {
				fmt.Println("Error creating GCP client:", err)
				return nil, errors.New("error creating GCP client")
			}
			gcpClient = client
		} else {
			fmt.Println("GCP client instance already created.")
		}

		// Register cleanup function if necessary for session
		RegisterCleanup(func() {
			fmt.Println("Cleanup FF session resources if needed.")
			// Add any cleanup logic here for the session if required
			err := gcpClient.Close()
			if err != nil {
				println("Error closing GCP client:", err)
			}
			gcpClient = nil
		})
	} else {
		fmt.Println("GCP client instance already created.")
	}

	return gcpClient, nil
}

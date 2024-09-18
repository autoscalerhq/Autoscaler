package bootstrap

import (
	"context"
	"fmt"
	"sync"

	"cloud.google.com/go/storage"
	"google.golang.org/api/option"
)

var gcpLock = &sync.Mutex{}
var gcpClient *storage.Client

// GetGCPClient returns a singleton instance of a GCP storage client
func GetGCPClient() *storage.Client {
	if gcpClient == nil {
		gcpLock.Lock()
		defer gcpLock.Unlock()
		if gcpClient == nil {
			fmt.Println("Creating single GCP client instance now.")
			ctx := context.Background()
			client, err := storage.NewClient(ctx, option.WithCredentialsFile("path/to/your/credentials/file.json"))
			if err != nil {
				fmt.Println("Error creating GCP client:", err)
				return nil
			}
			gcpClient = client
		} else {
			fmt.Println("GCP client instance already created.")
		}
	} else {
		fmt.Println("GCP client instance already created.")
	}

	return gcpClient
}

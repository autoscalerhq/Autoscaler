package bootstrap

import (
	"errors"
	"fmt"
	"sync"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
)

var azureLock = &sync.Mutex{}
var azureCredential azcore.TokenCredential

// GetAzureCredential returns a singleton instance of an Azure credential
func GetAzureCredential() (azcore.TokenCredential, error) {

	if shuttingDown {
		return nil, errors.New("sys shutdown")
	}

	if azureCredential == nil {
		azureLock.Lock()
		defer azureLock.Unlock()
		if azureCredential == nil {
			fmt.Println("Creating single Azure credential instance now.")
			cred, err := azidentity.NewDefaultAzureCredential(nil)
			if err != nil {
				fmt.Println("Error creating Azure credential:", err)
				return nil, errors.New("error creating credential")
			}
			azureCredential = cred

			RegisterCleanup(func() {
				fmt.Println("Cleanup Azure session resources if needed.")
				// Add any cleanup logic here for the session if required
				azureCredential = nil
			})
		} else {
			fmt.Println("Azure credential instance already created.")
		}
	} else {
		fmt.Println("Azure credential instance already created.")
	}

	return azureCredential, nil
}

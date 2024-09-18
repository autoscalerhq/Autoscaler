package bootstrap

import (
	"fmt"
	"sync"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
)

var azureLock = &sync.Mutex{}
var azureCredential azcore.TokenCredential

// GetAzureCredential returns a singleton instance of an Azure credential
func CetAzureCredential() azcore.TokenCredential {
	if azureCredential == nil {
		azureLock.Lock()
		defer azureLock.Unlock()
		if azureCredential == nil {
			fmt.Println("Creating single Azure credential instance now.")
			cred, err := azidentity.NewDefaultAzureCredential(nil)
			if err != nil {
				fmt.Println("Error creating Azure credential:", err)
				return nil
			}
			azureCredential = cred
		} else {
			fmt.Println("Azure credential instance already created.")
		}
	} else {
		fmt.Println("Azure credential instance already created.")
	}

	return azureCredential
}

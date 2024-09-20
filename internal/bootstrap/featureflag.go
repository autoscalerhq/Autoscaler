package bootstrap

import (
	"fmt"
	"sync"
	"time"

	flagd "github.com/open-feature/go-sdk-contrib/providers/flagd/pkg"
	"github.com/open-feature/go-sdk/pkg/openfeature"
)

var flagdProviderLock = &sync.Mutex{}
var flagdProviderInstance *flagd.Provider

// GetFlagdProviderInstance returns a singleton instance of the flagd provider
func GetFeatureFlagInstance() (*flagd.Provider, error) {
	if flagdProviderInstance == nil {
		flagdProviderLock.Lock()
		defer flagdProviderLock.Unlock()
		if flagdProviderInstance == nil {
			fmt.Println("Creating single flagd provider instance now.")
			providerOptions := []flagd.ProviderOption{
				flagd.WithBasicInMemoryCache(),
				flagd.WithRPCResolver(),
				flagd.WithHost("localhost"),
				flagd.WithPort(8013),
			}
			provider := flagd.NewProvider(providerOptions...)

			// Create an empty evaluation context
			evalContext := openfeature.NewEvaluationContext("key", map[string]interface{}{})

			// Initialize the provider with the context
			err := provider.Init(evalContext)
			if err != nil {
				return nil, fmt.Errorf("unable to initialize flagd provider: %w", err)
			}

			// Wait for the provider to be ready
			ready := waitForProvider(provider, 10*time.Second, 500*time.Millisecond)
			if !ready {
				return nil, fmt.Errorf("flagd provider not ready after waiting")
			}

			flagdProviderInstance = provider
		} else {
			fmt.Println("Flagd provider instance already created.")
		}
	} else {
		fmt.Println("Flagd provider instance already created.")
	}

	return flagdProviderInstance, nil
}

// waitForProvider waits for the provider to be ready, with a maximum wait time and retry interval.
func waitForProvider(provider *flagd.Provider, maxWait time.Duration, interval time.Duration) bool {
	start := time.Now()
	for {
		println("status", provider.Status())
		if provider.Status() == "READY" {
			return true
		}
		if time.Since(start) > maxWait {
			return false
		}
		time.Sleep(interval)
	}
}

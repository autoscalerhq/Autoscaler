package bootstrap

import (
	"errors"
	"fmt"
	"sync"

	"github.com/nats-io/nats.go/jetstream"
)

var jsLock = &sync.Mutex{}
var jsConn jetstream.JetStream

// GetJetStream returns a singleton instance of a NATS Jetstream connection
func GetJetStream(opts ...jetstream.JetStreamOpt) (jetstream.JetStream, error) {

	if shuttingDown {
		return nil, errors.New("sys shutdown")
	}

	nc, err := GetNatsConn()

	if err != nil {
		return nil, err
	}

	if jsConn == nil {
		jsLock.Lock()
		defer jsLock.Unlock()
		if jsConn == nil {
			fmt.Println("Creating single NATS Jetstream connection instance now.")
			js, err := jetstream.New(nc, opts...)
			if err != nil {
				return nil, fmt.Errorf("failed to create new jetstream: %w", err)
			}

			jsConn = js
		} else {
			fmt.Println("NATS connection instance already created.")
		}

		// Register cleanup function if necessary for session
		RegisterCleanup(func() {
			fmt.Println("Cleanup jetstream  if needed.")
			// Add any cleanup logic here for the session if required
			jsConn = nil
		})
	} else {
		fmt.Println("NATS connection instance already created.")
	}

	return jsConn, nil
}

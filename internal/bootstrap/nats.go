package bootstrap

import (
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/nats-io/nats.go"
)

var natsLock = &sync.Mutex{}
var natsConn *nats.Conn

// GetNatsConn returns a singleton instance of a NATS connection
func GetNatsConn() (*nats.Conn, error) {
	if natsConn == nil {
		natsLock.Lock()
		defer natsLock.Unlock()
		if natsConn == nil {
			fmt.Println("Creating single NATS connection instance now.")
			url := os.Getenv("NATS_URL")
			if url == "" {
				url = nats.DefaultURL
			}
			nc, err := nats.Connect(url,
				nats.DisconnectErrHandler(func(nc *nats.Conn, err error) {
					fmt.Printf("Got disconnected! Reason: %q\n", err)
				}),
				nats.ReconnectHandler(func(nc *nats.Conn) {
					fmt.Printf("Got reconnected to %v!\n", nc.ConnectedUrl())
				}),
				nats.ClosedHandler(func(nc *nats.Conn) {
					fmt.Printf("Connection closed. Reason: %q\n", nc.LastError())
				}),
				nats.ReconnectJitter(500*time.Millisecond, 2*time.Second),
				nats.MaxReconnects(5),
				nats.ReconnectWait(2*time.Second),
			)
			if err != nil {
				return nil, fmt.Errorf("failed to connect to nats: %w", err)
			}
			natsConn = nc
		} else {
			fmt.Println("NATS connection instance already created.")
		}
	} else {
		fmt.Println("NATS connection instance already created.")
	}

	return natsConn, nil
}

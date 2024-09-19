package bootstrap

import (
	"context"
	"errors"
	"fmt"
	"sync"

	"github.com/nats-io/nats.go/jetstream"
)

var kvLock = &sync.Mutex{}
var kvStoreMap = make(map[string]jetstream.KeyValue)

// GetKVStore returns a singleton instance of the requested key value store
// The config is optional; if not provided, a default configuration will be used.
func GetKVStore(configs ...jetstream.KeyValueConfig) (jetstream.KeyValue, context.Context, error) {

	var config jetstream.KeyValueConfig

	if len(configs) > 0 {
		config = configs[0]
	} else {
		config = jetstream.KeyValueConfig{Bucket: "defaultBucket"}
	}

	kvLock.Lock()
	defer kvLock.Unlock()

	if kv, exists := kvStoreMap[config.Bucket]; exists {
		fmt.Println("KV store instance already exists.")
		return kv, context.Background(), nil
	}

	fmt.Println("Creating new KV store instance.")
	ctx := context.Background()
	js, err := GetJetStream()
	if err != nil {
		return nil, nil, fmt.Errorf("error creating JetStream: %w", err)
	}

	kv, err := js.CreateKeyValue(ctx, config)
	if err != nil {
		if !errors.Is(err, jetstream.ErrBucketExists) {
			return nil, nil, fmt.Errorf("error creating key value store: %w", err)
		}
		kv, err = js.KeyValue(ctx, config.Bucket)
		if err != nil {
			return nil, nil, fmt.Errorf("error retrieving existing key value store: %w", err)
		}
	}

	kvStoreMap[config.Bucket] = kv
	return kv, ctx, nil
}

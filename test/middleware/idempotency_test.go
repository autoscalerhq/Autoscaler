package middleware

import (
	"bytes"
	"fmt"
	"github.com/autoscalerhq/autoscaler/internal/nats"
	"github.com/nats-io/nats.go/jetstream"
	"io"
	"log"
	"net/http"
	"testing"
)

// tests the idempotency middleware ensure both nats1 and nats 2 docker containers are running

const (
	IdempotencyKey    = "Idempotency-Key"
	IdempotencyBucket = "idempotent_requests"
)

// TestPostNewIdempotentItem
func TestPostNewIdempotentItem(t *testing.T) {

	// Create the request body
	jsonData := []byte(`{"name": "John Doe"}`)

	// Create a new request using http.NewRequest
	req, err := http.NewRequest("POST", "http://localhost:8090/Post", bytes.NewBuffer(jsonData))
	if err != nil {
		log.Fatal(err)
	}

	idempotentUuid := "038a3382-d1ec-4ffb-b2bc-cd4230ffb208"
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set(IdempotencyKey, idempotentUuid)

	// Perform the request
	client := &http.Client{}
	resp1, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	// Read and print the response body
	headers := resp1.Header
	body, err := io.ReadAll(resp1.Body)
	if err != nil {
		t.Error(err)
	}
	fmt.Println("Response Body:", string(body))
	fmt.Println("Response Headers:", headers)
	kv, ctx, err := natutils.NewKeyValueStore(jetstream.KeyValueConfig{
		Bucket: IdempotencyBucket,
	})
	if err != nil {
		t.Error(err)
	}
	err = kv.Delete(ctx, idempotentUuid)
	if err != nil {
		t.Errorf("Failed to delete key %s, %s", idempotentUuid, err)
	}
}

func TestPostTwoOfTheSame(t *testing.T) {

	// Create the request body
	jsonData := []byte(`{"name": "John Doe"}`)

	// Create a new request using http.NewRequest
	req, err := http.NewRequest("POST", "http://localhost:8090/Post", bytes.NewBuffer(jsonData))
	if err != nil {
		log.Fatal(err)
	}

	idempotentUuid := "038a3382-d1ec-4ffb-b2bc-cd4230ffb208"
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set(IdempotencyKey, idempotentUuid)

	// Perform the request
	client := &http.Client{}
	resp1, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	// Read and print the response body
	headers := resp1.Header
	body, err := io.ReadAll(resp1.Body)
	if err != nil {
		t.Error(err)
	}
	fmt.Println("Response Body:", string(body))
	fmt.Println("Response Headers:", headers)

	// Perform the request
	resp2, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	// Read and print the response body
	headers = resp2.Header
	body, err = io.ReadAll(resp2.Body)
	if err != nil {
		t.Error(err)
	}
	fmt.Println("Response Body:", string(body))
	fmt.Println("Response Headers:", headers)

	kv, ctx, err := natutils.NewKeyValueStore(jetstream.KeyValueConfig{
		Bucket: IdempotencyBucket,
	})
	if err != nil {
		t.Error(err)
	}

	err = kv.Delete(ctx, idempotentUuid)
	if err != nil {
		t.Errorf("Failed to delete key %s: %s", idempotentUuid, err)
	}
}

func TestListKeys(t *testing.T) {

	kv, ctx, err := natutils.NewKeyValueStore(jetstream.KeyValueConfig{
		Bucket: IdempotencyBucket,
	})

	if err != nil {
		t.Error(err)
	}

	keys, err := kv.Keys(ctx)
	if err != nil {
		t.Log(err)
	}

	fmt.Println("Keys:", keys)

	// Retrieve and print the value associated with each key
	for _, key := range keys {
		entry, err := kv.Get(ctx, key)
		if err != nil {
			log.Printf("Error retrieving value for key %s: %v\n", key, err)
			continue
		}

		fmt.Printf("Key: %s, Value: %s\n", key, string(entry.Value()))
	}
}

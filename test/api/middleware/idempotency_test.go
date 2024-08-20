package middleware

import (
	"bytes"
	"context"
	"fmt"
	"github.com/autoscalerhq/autoscaler/internal/nats"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/nats-io/nats.go/jetstream"
	"github.com/stretchr/testify/require"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

// tests the idempotency middleware ensure both nats1 and nats 2 docker containers are running

const (
	IdempotencyKey    = "Idempotency-Key"
	IdempotencyBucket = "idempotent_requests"
)

func teardown(key string) error {
	kv, idempotentCtx, err := natutils.NewKeyValueStore(jetstream.KeyValueConfig{Bucket: "idempotent_requests", TTL: time.Hour * 24})
	if err != nil {
		return fmt.Errorf("error getting new key value store: %w", err)
	}

	err = kv.Delete(idempotentCtx, key)
	if err != nil {
		log.Printf("Error deleting key %s: %v\n", key, err)
		return err
	}
	return nil
}

// TestPostNewIdempotentItem
func TestPostNewIdempotentItem(t *testing.T) {

	// Create the request body
	jsonData := []byte(`{"name": "John Doe"}`)

	// Create a new request using http.NewRequest
	req, err := http.NewRequest("POST", "http://localhost:8090/Post", bytes.NewBuffer(jsonData))
	if err != nil {
		log.Fatal(err)
	}
	idempotentUuid := uuid.New().String()
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

	defer func(key string) {
		err := teardown(key)
		if err != nil {
			t.Error(err)
		}
	}(idempotentUuid)

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

	defer func(key string) {
		err := teardown(key)
		if err != nil {
			t.Error(err)
		}
	}(idempotentUuid)
}

func TestPostToInProgressRequest(t *testing.T) {

	e := echo.New()
	idempotencyUuid := uuid.New()

	e.POST("/hang", func(c echo.Context) error {
		// Simulate a request that hangs for a while completes
		time.Sleep(5 * time.Second)
		return c.String(http.StatusBadRequest, "Hello, World!")
	})
	//e.Use(appmiddleware.IdempotencyMiddleware(kv, idempotentCtx))
	server := httptest.NewServer(e)
	defer server.Close()

	// Create a request to the hanging route
	req, err := http.NewRequest(http.MethodPost, server.URL+"/hang", nil)
	require.NoError(t, err)

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	req = req.WithContext(ctx)
	req.Header.Set(IdempotencyKey, idempotencyUuid.String())
	// Perform the request
	client := &http.Client{}
	resp, err := client.Do(req)
	_, err = client.Do(req)

	// Check if the request timed out
	require.Error(t, err, "Expected a timeout error")
	require.Contains(t, err.Error(), "context deadline exceeded", "Expected context deadline exceeded error")

	if resp != nil {
		defer func(Body io.ReadCloser) {
			err := Body.Close()
			if err != nil {
				t.Error(err)
			}
		}(resp.Body)
	}

	defer func(key string) {
		err := teardown(key)
		if err != nil {
			t.Error(err)
		}
	}(idempotencyUuid.String())
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

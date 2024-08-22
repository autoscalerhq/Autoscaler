package middleware

import (
	"context"
	"fmt"
	"github.com/autoscalerhq/autoscaler/internal/nats"
	appmiddleware "github.com/autoscalerhq/autoscaler/services/api/middleware"
	apphttp "github.com/autoscalerhq/autoscaler/services/api/util"
	"github.com/labstack/echo/v4"
	"github.com/nats-io/nats.go/jetstream"
	"github.com/stretchr/testify/assert"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"runtime"
	"sync"
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

func setup() *httptest.Server {

	kv, idempotentCtx, err := natutils.NewKeyValueStore(jetstream.KeyValueConfig{Bucket: "idempotent_requests", TTL: time.Hour * 24})
	if err != nil {
		return nil
	}
	e := echo.New()
	e.Use(appmiddleware.IdempotencyMiddleware(kv, idempotentCtx))
	e.POST("/hang", func(c echo.Context) error {
		// Simulate a request that hangs
		time.Sleep(1 * time.Second)
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.POST("/Post", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	server := httptest.NewServer(e)
	return server
}

// TestPostNewIdempotentItem
func TestPostNewIdempotentItem(t *testing.T) {

	server := setup()
	defer server.Close()
	// Create the request body
	jsonData := map[string]string{
		"name": "John Doe",
	}
	// Create a new request using http.NewRequest
	idempotentUuid := "038a3382-d1ec-4ffb-b2bc-cd4230ffb208"
	header := http.Header{
		"Idempotency-Key": []string{idempotentUuid},
	}
	resp, headers, err := apphttp.Post(server.URL+"/Post", jsonData, header, nil)
	if err != nil {
		t.Error(err)
	}
	data, err := io.ReadAll(resp.Body)

	// Read and print the response body
	fmt.Println("Response Body:", string(data))
	fmt.Println("Response Status:", resp.StatusCode)
	fmt.Println("Response Headers:", headers)

	defer func(key string) {
		err := teardown(key)
		if err != nil {
			t.Error(err)
		}
	}(idempotentUuid)

}

func TestPostTwoOfTheSame(t *testing.T) {
	server := setup()
	defer server.Close()

	jsonData := map[string]string{
		"name": "John Doe",
	}
	// Create a new request using http.NewRequest
	idempotentUuid := "038a3382-d1ec-4ffb-b2bc-cd4230ffb208"
	header := http.Header{
		"Idempotency-Key": []string{idempotentUuid},
	}

	resp, headers, err := apphttp.Post(server.URL+"/Post", jsonData, header, nil)
	if err != nil {
		return
	}

	// Perform the request
	resp1, headers2, err2 := apphttp.Post(server.URL+"/Post", jsonData, header, nil)
	if err2 != nil {
		return
	}
	body, err := io.ReadAll(resp.Body)
	body1, err := io.ReadAll(resp1.Body)

	fmt.Println("Response Body:", string(body))
	fmt.Println("Response Status", resp.StatusCode)
	fmt.Println("Response Headers:", headers)

	fmt.Println("Response Body:", string(body1))
	fmt.Println("Response Status", resp1.StatusCode)
	fmt.Println("Response Headers:", headers2)

	assert.Equal(t, resp.StatusCode, resp1.StatusCode)

	defer func(key string) {
		err := teardown(key)
		if err != nil {
			t.Error(err)
		}
	}(idempotentUuid)
}

func TestConcurrentRequests(t *testing.T) {

	fmt.Println("Version", runtime.Version())
	fmt.Println("NumCPU", runtime.NumCPU())
	fmt.Println("GOMAXPROCS", runtime.GOMAXPROCS(0))
	server := setup()
	defer server.Close()

	idempotentUuid := "038a3382-d1ec-4ffb-b2bc-cd4230ffb208"
	header := http.Header{
		"Idempotency-Key": []string{idempotentUuid},
	}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var requestWg sync.WaitGroup
	requestWg.Add(2)

	// First goroutine to send the request
	go func() {
		defer requestWg.Done()
		fmt.Println("Goroutine 1 - Sent Request")
		resp, headers, err := apphttp.Post(server.URL+"/hang", nil, header, ctx)
		if err != nil {
			t.Error(err)
			return
		}
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			t.Error(err)
			return
		}
		fmt.Println("Goroutine 1 - Response Body:", string(body))
		fmt.Println("Goroutine 1 - Response Status", resp.StatusCode)
		fmt.Println("Goroutine 1 - Response Headers:", headers)
	}()

	// Second goroutine to send the request with a slight delay
	go func() {
		defer requestWg.Done()

		time.Sleep(50 * time.Millisecond)
		fmt.Println("Goroutine 2 - Sent Request")
		resp, headers, err := apphttp.Post(server.URL+"/hang", nil, header, ctx)
		if err != nil {
			t.Error(err)
			return
		}
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			t.Error(err)
			return
		}
		fmt.Println("Goroutine 2 - Response Body:", string(body))
		fmt.Println("Goroutine 2 - Response Status", resp.StatusCode)
		fmt.Println("Goroutine 2 - Response Headers:", headers)
	}()

	// Wait for all requests to complete
	requestWg.Wait()
	// Cleanup
	defer func(key string) {
		err := teardown(key)
		if err != nil {
			t.Error(err)
		}
	}(idempotentUuid)
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
		//_ = kv.Delete(ctx, key)

		fmt.Printf("Key: %s, Value: %s\n", key, string(entry.Value()))
	}
}

package middleware

import (
	"bytes"
	"fmt"
	"github.com/google/uuid"
	"io"
	"net/http"
	"testing"
)

func TestPostIdempotency(t *testing.T) {
	header := http.Header{}
	idempotentUuid := uuid.New()
	header.Add("Idempotency-key", idempotentUuid.String())

	resp, err := http.Post("http://localhost:8090/Post", "application/json", bytes.NewBuffer([]byte(`{"name": "John Doe"}`)))
	if err != nil {
		t.Error(err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)

	fmt.Println("Response Status:", resp.Status)

	// Read and print the response body
	headers := resp.Header
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Error(err)
	}
	fmt.Println("Response Body:", string(body))
	fmt.Println("Response Headers:", headers)
}

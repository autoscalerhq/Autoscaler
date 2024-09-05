package appmiddleware

import (
	"bytes"
	"context"
	"crypto"
	"encoding/json"
	"errors"
	"github.com/labstack/echo/v4"
	"github.com/nats-io/nats.go/jetstream"
	"io"
	"net/http"
)

const (
	IdempotencyKey = "Idempotency-Key"
)

type CapturingResponseWriter struct {
	http.ResponseWriter
	Body *bytes.Buffer
}

// Override the Write method to capture the response body
func (w *CapturingResponseWriter) Write(b []byte) (int, error) {
	w.Body.Write(b)
	return w.ResponseWriter.Write(b)
}

type IdempotencyStatus int

const (
	Processing IdempotencyStatus = iota
	Completed
)

type Message struct {
	StatusCode        int               `json:"statusCode,omitempty"`
	Body              string            `json:"body,omitempty"`
	IdempotencyStatus IdempotencyStatus `json:"idempotencyStatus,omitempty"`
	Hash              []byte            `json:"hash,omitempty"`
}

// Function that handles the hashing and capturing of the request body
func hashAndCaptureBody(body io.ReadCloser) ([]byte, *bytes.Buffer, error) {
	// Create a new buffer to store the request body
	buf := new(bytes.Buffer)

	// Creating hashed request key
	h := crypto.BLAKE2b_512.New()

	// Create a MultiWriter that writes to both the buffer and the hash
	multiWriter := io.MultiWriter(buf, h)

	// Read the request body from the original request into the MultiWriter
	_, err := io.Copy(multiWriter, body)
	if err != nil {
		return nil, nil, err
	}

	// Compute the hash
	hashRequestBody := h.Sum(nil)

	return hashRequestBody, buf, nil
}

func IdempotencyMiddleware(kv jetstream.KeyValue, ctx context.Context) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			if c.Request().Method == "POST" || c.Request().Method == "PATCH" {

				capturingWriter := &CapturingResponseWriter{
					ResponseWriter: c.Response().Writer,
					Body:           new(bytes.Buffer),
				}
				c.Response().Writer = capturingWriter

				// Check if the request is idempotent
				key := c.Request().Header.Get(IdempotencyKey)
				if key == "" {
					return next(c)
				}

				if len(key) > 255 {
					return c.JSON(http.StatusUnprocessableEntity, Message{Body: "Idempotency key is too long"})
				}

				// Hash the request body and capture it for later use
				hashRequestBody, requestBody, err := hashAndCaptureBody(c.Request().Body)
				if err != nil {
					return err
				}

				// Restore the request body to the original state so it can be read again downstream
				c.Request().Body = io.NopCloser(requestBody)

				// get key if it exists
				existingKey, err := kv.Get(ctx, key)
				if err != nil {
					if !errors.Is(err, jetstream.ErrKeyNotFound) {
						return err
					}
				}

				// Check if the request has already been processed or is being processed
				if existingKey != nil {
					var result Message
					err := json.Unmarshal(existingKey.Value(), &result)
					if err != nil {
						return err
					}
					// if the request body has changed
					if !bytes.Equal(hashRequestBody, result.Hash) {
						return c.JSON(http.StatusUnprocessableEntity, Message{Body: "Request body does not match the original request"})
					}

					// If the request is being processed, return status code 409
					if result.IdempotencyStatus == Processing {
						c.Response().Header().Set("Retry-After", "60") // header describes when client should retry in seconds
						return c.JSON(http.StatusConflict, Message{Body: "Request is being processed"})
					}

					// set the response header
					c.Response().Header().Set("Idempotency-Replay", "true")
					// If the request has already been processed return original response from kv
					return c.JSON(result.StatusCode, result.Body)
				}

				// If the request has not been processed, store the request with the status code accepted,
				// signifying that the request is being processed
				msg := Message{
					Hash:              hashRequestBody,
					IdempotencyStatus: Processing,
				}

				data, err := json.Marshal(msg)
				if err != nil {
					// TODO report these kind of errors to a monitoring system
					return err
				}
				_, err = kv.Put(ctx, key, data)
				if err != nil {
					return err
				}

				// Process the request then update status to the response status
				c.Response().Header().Set("Idempotency-Replay", "false")
				err = next(c)
				if err == nil {
					resp, respErr := json.Marshal(
						Message{
							StatusCode:        c.Response().Status,
							Body:              capturingWriter.Body.String(),
							Hash:              hashRequestBody,
							IdempotencyStatus: Completed,
						})

					if respErr != nil {
						return err
					}

					_, err = kv.Put(ctx, key, resp)
					if err != nil {
						return err
					}
				}
				return err
			}
			return next(c)
		}
	}
}

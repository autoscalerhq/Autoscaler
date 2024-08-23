package appmiddleware

import (
	"bytes"
	"context"
	"crypto"
	"encoding/json"
	"errors"
	"github.com/google/uuid"
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

type IdempotencyStatus string

const (
	Processing IdempotencyStatus = "Processing"
	Completed  IdempotencyStatus = "Completed"
)

type Message struct {
	StatusCode        int               `json:"statusCode,omitempty"`
	Body              string            `json:"body,omitempty"`
	IdempotencyStatus IdempotencyStatus `json:"-"`
	Hash              string            `json:"-"`
}

//TODO 422 Unprocessable Entity missing case If there's an attempt to reuse an idempotency key with a different request payload.

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

				// Validate the idempotency key is a valid UUID
				idempotentUUID, err := uuid.Parse(key)
				if err != nil {
					return c.JSON(http.StatusBadRequest, Message{Body: "Idempotency key must be a valid UUID"})
				}

				requestBody, err := io.ReadAll(c.Request().Body)
				if err != nil {
					return err
				}

				// Creating hashed request key
				h := crypto.BLAKE2b_512.New()
				_, err = h.Write(requestBody)
				if err != nil {
					return err
				}
				hashRequestBody := h.Sum(nil)

				// get key if it exists
				existingKey, err := kv.Get(ctx, idempotentUUID.String())
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

					// if the request body has changed, but uuid has not return status code 422
					if string(hashRequestBody) != result.Hash {
						return c.JSON(http.StatusUnprocessableEntity, Message{Body: "Request body does not match the original request"})
					}

					// If the request is being processed, return status code 409
					if result.IdempotencyStatus == Processing {
						return c.JSON(http.StatusConflict, Message{Body: "Request is being processed"})
					}

					// If the request has already been processed return original response from kv
					return c.JSON(result.StatusCode, result.Body)
				}

				// If the request has not been processed, store the request with the status code accepted,
				// signifying that the request is being processed
				msg := Message{
					Hash:              string(hashRequestBody),
					IdempotencyStatus: Processing,
				}
				data, err := json.Marshal(msg)
				if err != nil {
					// TODO report these kind of errors to a monitoring system
					return err
				}
				_, err = kv.Put(ctx, idempotentUUID.String(), data)
				if err != nil {
					return err
				}

				// Process the request then update status to the response status
				err = next(c)
				if err == nil {
					resp, respErr := json.Marshal(
						Message{
							StatusCode:        c.Response().Status,
							Body:              capturingWriter.Body.String(),
							Hash:              string(hashRequestBody),
							IdempotencyStatus: Completed,
						})
					if respErr != nil {
						return err
					}

					_, err = kv.Put(ctx, idempotentUUID.String(), resp)
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

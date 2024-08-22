package appmiddleware

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/nats-io/nats.go/jetstream"
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

type Message struct {
	StatusCode int    `json:"statusCode,omitempty"`
	Body       string `json:"body"`
}

//TODO 422 Unprocessable Entity missing case If there's an attempt to reuse an idempotency key with a different request payload.

func IdempotencyMiddleware(kv jetstream.KeyValue, ctx context.Context) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			capturingWriter := &CapturingResponseWriter{
				ResponseWriter: c.Response().Writer,
				Body:           new(bytes.Buffer),
			}
			c.Response().Writer = capturingWriter

			if c.Request().Method == "POST" || c.Request().Method == "PATCH" {
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

				// get key if it exists
				existingKey, err := kv.Get(ctx, idempotentUUID.String())
				if err != nil {
					if !errors.Is(err, jetstream.ErrKeyNotFound) {
						// TODO report these kind of errors to a monitoring system
						return c.JSON(http.StatusInternalServerError, Message{Body: "Internal Server Error, please try again later"})
					}
				}

				// Check if the request has already been processed or is being processed
				if existingKey != nil {
					var result Message
					err := json.Unmarshal(existingKey.Value(), &result)
					if err != nil {
						// TODO report these kind of errors to a monitoring system
						return c.JSON(http.StatusInternalServerError, Message{Body: "Internal Server Error, please try again later"})
					}

					// If the request is being processed, return status code 409
					if http.StatusAccepted == result.StatusCode {
						return c.JSON(http.StatusConflict, Message{Body: "Request is being processed"})
					}

					// If the request has already been processed return original response from nats
					return c.JSON(result.StatusCode, result.Body)
				}

				// If the request has not been processed, store the request with the status code accepted,
				// signifying that the request is being processed
				msg := Message{
					StatusCode: http.StatusAccepted,
					Body:       "Request is being processed",
				}
				data, err := json.Marshal(msg)
				if err != nil {
					// TODO report these kind of errors to a monitoring system
					return c.JSON(http.StatusInternalServerError, Message{Body: "Internal Server Error, please try again later"})
				}
				_, err = kv.Put(ctx, idempotentUUID.String(), data)
				if err != nil {
					// TODO report these kind of errors to a monitoring system
					return c.JSON(http.StatusInternalServerError, Message{Body: "Internal Server Error, please try again later"})
				}

				// Process the request then update status to the response status
				err = next(c)
				if err == nil {
					resp, respErr := json.Marshal(
						Message{
							StatusCode: c.Response().Status,
							Body:       capturingWriter.Body.String(),
						})
					if respErr != nil {
						// TODO report these kind of errors to a monitoring system
						return c.JSON(http.StatusInternalServerError, Message{Body: "Internal Server Error, please try again later"})
					}

					_, err = kv.Put(ctx, idempotentUUID.String(), resp)
					if err != nil {
						// TODO report these kind of errors to a monitoring system
						return c.JSON(http.StatusInternalServerError, Message{Body: "Internal Server Error, please try again later"})
					}
				}
				return err
			}
			return next(c)
		}
	}
}

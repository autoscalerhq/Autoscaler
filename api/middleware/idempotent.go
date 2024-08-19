package appmiddleware

import (
	"context"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/nats-io/nats.go/jetstream"
	"net/http"
	"strconv"
)

const (
	IdempotencyKey = "Idempotency-Key"
)

func IdempotencyMiddleware(kv jetstream.KeyValue, ctx context.Context) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			fmt.Print("Idempotency Middleware Has been CALLED")
			if c.Request().Method == "POST" || c.Request().Method == "PATCH" {
				// Check if the request is idempotent
				key := c.Request().Header.Get(IdempotencyKey)
				if key == "" {
					return next(c)
				}

				// validate the idempotency key is a valid UUID
				idempotentUUID, err := uuid.Parse(key)
				if err != nil {
					return c.String(http.StatusBadRequest, "Idempotency key must be a valid UUID")
				}

				// get key if it exists
				value, err := kv.Get(ctx, idempotentUUID.String())
				if err != nil {
					if !errors.Is(err, jetstream.ErrKeyNotFound) {
						return fmt.Errorf("failed to get key %s from KeyValueStore: %w", idempotentUUID.String(), err)
					}
				}

				// check if the request has already been processed or is being processed
				if value != nil {
					status := string(value.Value())
					statusCode, err := strconv.Atoi(status)
					if err != nil {
						return fmt.Errorf("failed converting key -> value from []byte to int :%w", err)
					}
					if statusCode == http.StatusAccepted {
						return c.String(http.StatusConflict, "Request is being processed")
					}

					return c.String(statusCode, "Request has already been processed")
				}

				// If the request has not been processed, store the request

				_, err = kv.Put(ctx, idempotentUUID.String(), []byte(strconv.Itoa(http.StatusAccepted)))
				if err != nil {
					return fmt.Errorf("failed to put new key %s into KeyValueStore: %w", idempotentUUID.String(), err)
				}

				// Process the request update status to the response status
				_, err = kv.Put(ctx, idempotentUUID.String(), []byte(strconv.Itoa(c.Response().Status)))
				if err != nil {
					return fmt.Errorf("failed to update key %s with response status code: %w", idempotentUUID.String(), err)
				}
			}
			return next(c)
		}
	}
}

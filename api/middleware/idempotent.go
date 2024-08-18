package appmiddleware

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/nats-io/nats.go/jetstream"
	"net/http"
	"strconv"
)

func IdempotencyMiddleware(kv jetstream.KeyValue, ctx context.Context) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			if c.Request().Method == "POST" || c.Request().Method == "PATCH" {
				// Check if the request is idempotent
				key := c.Response().Header().Get("Idempotency_key")
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
						return err
					}
				}

				// check if the request has already been processed or is being processed
				if value != nil {
					status, err := strconv.Atoi(string(value.Value()))
					if err != nil {
						return err
					}

					if status == http.StatusAccepted {
						return c.String(http.StatusConflict, "Request is being processed")
					}

					return c.String(status, "Request has already been processed")
				}

				// If the request has not been processed, store the request
				_, err = kv.Put(ctx, idempotentUUID.String(), []byte(string(rune(http.StatusAccepted))))
				if err != nil {
					return err
				}

				// Process the request update status to the response status
				_, err = kv.Put(ctx, idempotentUUID.String(), []byte(string(rune(c.Response().Status))))
				if err != nil {
					return err
				}
			}
			return next(c)
		}
	}
}

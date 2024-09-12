package middleware

import (
	"context"
	loadshedhttp "github.com/kevinconway/loadshed/v2/stdlib/net/http"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/nats-io/nats.go/jetstream"
)

type NatsKeyValue struct {
	KeyValueStore jetstream.KeyValue
	Context       context.Context
}

func ApplyMiddleware(e *echo.Echo, nats NatsKeyValue) {
	e.Use(middleware.Logger())
	// Middleware
	e.Use(RequestCounterMiddleware)
	e.Use(AddRouteToCTXMiddleware)
	// If load is too high, fail before we process anything else. this may need to be moved after logging
	e.Use(echo.WrapMiddleware(loadshedhttp.NewHandlerMiddleware(CreateShedder(), loadshedhttp.HandlerOptionCallback(&RejectionHandler{}))))
	e.Use(middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(100)))
	e.Use(IdempotencyMiddleware(nats.KeyValueStore, nats.Context))
	e.Use(TracingMiddleware())
	e.Use(middleware.Recover())
	ApplyAuthAndCorsMiddleware(e)
}

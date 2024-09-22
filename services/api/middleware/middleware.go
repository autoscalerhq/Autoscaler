package middleware

import (
	"context"
	"github.com/autoscalerhq/autoscaler/services/api/auth"
	loadshedhttp "github.com/kevinconway/loadshed/v2/stdlib/net/http"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/nats-io/nats.go/jetstream"
	"go.opentelemetry.io/contrib/instrumentation/github.com/labstack/echo/otelecho" // nolint:staticcheck  // deprecated.
)

type NatsKeyValue struct {
	KeyValueStore jetstream.KeyValue
	Context       context.Context
}
type MiddlewareParams struct {
	Nats NatsKeyValue
}

func ApplyMiddleware(e *echo.Echo, params MiddlewareParams) {
	e.Use(middleware.Logger())
	// Middleware
	e.Use(RequestCounterMiddleware)
	e.Use(AddRouteToCTXMiddleware)
	// If load is too high, fail before we process anything else. this may need to be moved after logging
	e.Use(echo.WrapMiddleware(loadshedhttp.NewHandlerMiddleware(CreateShedder(), loadshedhttp.HandlerOptionCallback(&RejectionHandler{}))))
	e.Use(middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(100)))
	e.Use(TracingMiddleware())
	e.Use(middleware.Recover())
	e.Use(otelecho.Middleware("echo-conn-server"))
	ApplyAuthAndCorsMiddleware(e)
}

func ApplyProtectedMiddleware(e *echo.Group, params MiddlewareParams) {
	e.Use(echo.WrapMiddleware(auth.VerifySessionMiddleware))
	e.Use(CorsMiddleware())
	e.Use(IdempotencyMiddleware(params.Nats.KeyValueStore, params.Nats.Context))
}

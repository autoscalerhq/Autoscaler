package main

import (
	"context"
	"errors"
	_ "github.com/autoscalerhq/autoscaler/api/docs"
	"github.com/autoscalerhq/autoscaler/api/routes"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
	m "go.opentelemetry.io/otel/metric"
	t "go.opentelemetry.io/otel/trace"
	"log/slog"
	"os"
	"os/signal"
)

// These variables are to be used throughout the application, for logging, tracing, and metrics.
// To understand the differences between these,
// please navigate to this link: https://opentelemetry.io/docs/languages/go/getting-started/
// it's thread safe, @todo in progress.
var (
	name   string
	tracer t.Tracer
	meter  m.Meter
	logger *slog.Logger
)

func main() {

	// Handle SIGINT (CTRL+C) gracefully.
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	// Set up OpenTelemetry.
	otelShutdown, err := setupOTelSDK(ctx)
	if err != nil {
		return
	}

	// Handle shutdown properly so nothing leaks.
	defer func() {
		err = errors.Join(err, otelShutdown(context.Background()))
	}()

	// initializing global variables before the application starts.
	// TODO this needs to be re thought the name should be set based off the file that is using these variables

	//_, file, _, _ := runtime.Caller(1)
	//name = filepath.Base(file)
	//tracer = otel.Tracer(name)
	//meter = otel.Meter(name)
	//logger = otelslog.NewLogger(name) // Replace with actual logger initialization
	e := echo.New()

	// Middleware
	e.Use(tracingMiddleware)
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())
	e.Use(middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(100)))

	// Routes
	e.GET("/Health", routes.HealthCheck)
	e.GET("/Hello", func(c echo.Context) error {
		return c.String(200, "Hello!")
	})

	// swag init -g ./main.go --output ./docs
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	e.Logger.Fatal(e.Start(":8080"))
}

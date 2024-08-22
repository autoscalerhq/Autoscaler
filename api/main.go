package main

import (
	"fmt"
	_ "github.com/autoscalerhq/autoscaler/api/docs"
	"github.com/autoscalerhq/autoscaler/api/middleware"
	"github.com/autoscalerhq/autoscaler/api/routes"
	natutils "github.com/autoscalerhq/autoscaler/internal/nats"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/nats-io/nats.go"
	"github.com/nats-io/nats.go/jetstream"
	echoSwagger "github.com/swaggo/echo-swagger"
	m "go.opentelemetry.io/otel/metric"
	t "go.opentelemetry.io/otel/trace"
	"log/slog"
	"time"
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

	err := godotenv.Load()
	if err != nil {
		// Todo this should be handled better before going to production
		//log.Fatal("Error loading .env file")
	}
	// Handle SIGINT (CTRL+C) gracefully.
	//ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	//defer stop()

	// Set up OpenTelemetry.
	//otelShutdown, err := setupOTelSDK(ctx)
	//if err != nil {
	//	return
	//}

	// Handle shutdown properly so nothing leaks.
	//defer func() {
	//	err = errors.Join(err, otelShutdown(context.Background()))
	//}()

	// initializing global variables before the application starts.
	// TODO this needs to be re thought the name should be set based off the file that is using these variables
	//_, file, _, _ := runtime.Caller(1)
	//name = filepath.Base(file)
	//tracer = otel.Tracer(name)
	//meter = otel.Meter(name)
	//logger = otelslog.NewLogger(name) // Replace with actual logger initialization
	e := echo.New()

	//Middleware
	nc, err := natutils.GetNatsConn()
	defer func(nc *nats.Conn) {
		err := nc.Drain()
		if err != nil {
		}
	}(nc)
	kv, idempotentCtx, err := natutils.NewKeyValueStore(jetstream.KeyValueConfig{Bucket: "idempotent_requests", TTL: time.Hour * 24})
	if err != nil {
		e.Logger.Fatal(fmt.Errorf("error getting new key value store: %w", err))
	}

	// extend the echo.Context to include custom context
	e.Use(appmiddleware.IdempotencyMiddleware(kv, idempotentCtx))
	e.Use(appmiddleware.TracingMiddleware())
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
	e.POST("/Post", func(c echo.Context) error {
		return c.String(200, "POST")
	})
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	e.Logger.Fatal(e.Start(":8090"))
}

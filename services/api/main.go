package main

import (
	"context"
	"errors"
	_ "github.com/autoscalerhq/autoscaler/services/api/docs"
	"github.com/autoscalerhq/autoscaler/services/api/middleware"
	"github.com/autoscalerhq/autoscaler/services/api/routes"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	flagd "github.com/open-feature/go-sdk-contrib/providers/flagd/pkg"
	"github.com/open-feature/go-sdk/openfeature"
	echoSwagger "github.com/swaggo/echo-swagger"
	m "go.opentelemetry.io/otel/metric"
	t "go.opentelemetry.io/otel/trace"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
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

	providerOptions := []flagd.ProviderOption{
		flagd.WithBasicInMemoryCache(),
		flagd.WithRPCResolver(),
		flagd.WithHost("localhost"),
		flagd.WithPort(8013),
	}

	provider := flagd.NewProvider(providerOptions...)

	err = openfeature.SetProvider(provider)
	if err != nil {
		println("Open Feature flag setup err: ", err.Error())
		return
	}

	// Create an empty evaluation context
	evalContext := openfeature.NewEvaluationContext("key", map[string]interface{}{})

	err = provider.Init(evalContext)
	if err != nil {
		println("Unable to init", err.Error())
		return
	}

	// Wait for the provider to be ready
	ready := waitForProvider(provider, 10*time.Second, 500*time.Millisecond)
	if !ready {
		println("Provider not ready after waiting")
		return
	}

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
	e.Use(appmiddleware.TracingMiddleware)
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

	ctx, stop = signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()
	// Start server
	go func() {
		if err := e.Start(":8888"); err != nil && !errors.Is(err, http.ErrServerClosed) {
			e.Logger.Fatal("shutting down the server")
		}
	}()

	// Wait for the interrupt signal to gracefully shut down the server with a timeout of 10 seconds.
	<-ctx.Done()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
}

// waitForProvider waits for the provider to be ready, with a maximum wait time and retry interval.
func waitForProvider(provider *flagd.Provider, maxWait time.Duration, interval time.Duration) bool {
	start := time.Now()
	for {
		println("status", provider.Status())
		if provider.Status() == "READY" {
			return true
		}
		if time.Since(start) > maxWait {
			return false
		}
		time.Sleep(interval)
	}
}

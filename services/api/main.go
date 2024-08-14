package main

import (
	"context"
	"errors"
	"github.com/autoscalerhq/autoscaler/services/api/middleware"
	"github.com/autoscalerhq/autoscaler/services/api/monitoring"
	"github.com/autoscalerhq/autoscaler/services/api/routes"
	"github.com/go-co-op/gocron"
	"github.com/joho/godotenv"
	loadshedhttp "github.com/kevinconway/loadshed/v2/stdlib/net/http"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
	m "go.opentelemetry.io/otel/metric"
	t "go.opentelemetry.io/otel/trace"
	"log/slog"
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
	// Handle SIGINT (CTRL+C) gracefully.
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	otelShutdown, err := monitoring.SetupOTelSDK(ctx)
	if err != nil {
		return
	}

	// Handle shutdown properly so nothing leaks.
	defer func() {
		err = errors.Join(err, otelShutdown(context.Background()))
		if err != nil {
			println(err.Error(), "unable to stop otel")
		}
	}()

	pscope, err := monitoring.SetupProfiling()
	if err != nil {
		return
	}

	defer func() {
		err := pscope.Stop()
		if err != nil {
			println(err.Error(), "unable to stop pyroscope")
		}
	}()

	s := gocron.NewScheduler(time.UTC)
	_, err = s.Every(50).Milliseconds().Do(appmiddleware.GetCPUUsage)
	if err != nil {
		println(err.Error(), "unable to start scheduler for CPU")
		return
	}

	s.StartAsync()

	// initializing global variables before the application starts.
	// TODO this needs to be re thought the name should be set based off the file that is using these variables

	//_, file, _, _ := runtime.Caller(1)
	//name = filepath.Base(file)
	//tracer = otel.Tracer(name)
	//meter = otel.Meter(name)
	//logger = otelslog.NewLogger(name) // Replace with actual logger initialization
	e := echo.New()

	// Middleware
	e.Use(appmiddleware.RequestCounterMiddleware)
	e.Use(appmiddleware.AddRouteToCTXMiddleware)
	// If load is too high, fail before we process anything else. this may need to be moved after logging
	e.Use(echo.WrapMiddleware(loadshedhttp.NewHandlerMiddleware(appmiddleware.CreateShedder(), loadshedhttp.HandlerOptionCallback(&appmiddleware.RejectionHandler{}))))

	e.Use(appmiddleware.TracingMiddleware)
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())
	e.Use(middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(100)))

	// TODO define routes in another method
	// Routes
	e.GET("/health-check", routes.HealthCheck)
	e.GET("/Hello", func(c echo.Context) error {
		return c.String(200, "Hello!")
	})

	// swag init -g ./main.go --output ./docs
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	// TODO allow port costumization
	e.Logger.Fatal(e.Start(":8888"))
}

package main

import (
	"context"
	"errors"
	"github.com/autoscalerhq/autoscaler/internal/bootstrap"
	"github.com/autoscalerhq/autoscaler/services/api/auth"
	"github.com/autoscalerhq/autoscaler/services/api/middleware"
	"github.com/autoscalerhq/autoscaler/services/api/monitoring"
	"github.com/autoscalerhq/autoscaler/services/api/routes"
	"github.com/go-co-op/gocron"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/nats-io/nats.go/jetstream"
	"github.com/supertokens/supertokens-golang/supertokens"
	m "go.opentelemetry.io/otel/metric"
	t "go.opentelemetry.io/otel/trace"
	"log"
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

type Environment struct {
	supertokens   auth.SuperTokensEnv
	listenAddress string
}

func makeDefaultEnv() Environment {
	return Environment{
		supertokens:   auth.MakeDefaultSuperTokensAppInfoEnv(),
		listenAddress: ":8888",
	}
}

// For local development, Nats 1, 2 And flagd must be running
func main() {

	env := makeDefaultEnv()
	err := auth.InitSuperTokens(env.supertokens)
	if err != nil {
		log.Fatal(err)
	}

	err = godotenv.Load()
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
	// Handle shutting down open telemetry when the program stops
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
	// Shutdown pyroscope profiling when the application stops.
	defer func() {
		err := pscope.Stop()
		if err != nil {
			println(err.Error(), "unable to stop pyroscope")
		}
	}()

	s := gocron.NewScheduler(time.UTC)
	_, err = s.Every(50).Milliseconds().Do(middleware.GetCPUUsage)
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

	idempotencykv, idempotentCtx, err := bootstrap.GetKVStore(jetstream.KeyValueConfig{
		Bucket:   "idempotent_requests_api",
		TTL:      time.Hour * 24,
		MaxBytes: 1024 * 1024,
	})

	if err != nil {
		println(err.Error(), "error getting new key value store")
		return
	}
	supertokens.DebugEnabled = true
	e := echo.New()
	middlewareParams := middleware.MiddlewareParams{
		Nats: middleware.NatsKeyValue{KeyValueStore: idempotencykv, Context: idempotentCtx},
	}
	middleware.ApplyMiddleware(e, middlewareParams)
	routes.Route(e, middlewareParams)
	ctx, stop = signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	// Start server
	go func() {
		if err := e.Start(env.listenAddress); err != nil && !errors.Is(err, http.ErrServerClosed) {
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

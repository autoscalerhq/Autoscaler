package main

import (
	"context"
	"errors"
	natutils "github.com/autoscalerhq/autoscaler/internal/nats"
	"github.com/autoscalerhq/autoscaler/services/api/auth"
	"github.com/autoscalerhq/autoscaler/services/api/middleware"
	"github.com/autoscalerhq/autoscaler/services/api/monitoring"
	"github.com/autoscalerhq/autoscaler/services/api/routes"
	"github.com/go-co-op/gocron"
	"github.com/joho/godotenv"
	loadshedhttp "github.com/kevinconway/loadshed/v2/stdlib/net/http"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/nats-io/nats.go"
	"github.com/nats-io/nats.go/jetstream"
	flagd "github.com/open-feature/go-sdk-contrib/providers/flagd/pkg"
	"github.com/open-feature/go-sdk/openfeature"
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
		listenAddress: "localhost:4000",
	}
}

// For local development, Nats 1, 2 And flagd must be running
func main() {

	err := auth.InitSuperTokens(auth.MakeDefaultSuperTokensAppInfoEnv())
	if err != nil {
		log.Fatal(err)
	}
	err = godotenv.Load()
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
	//
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

	//Wait for the provider to be ready
	ready := waitForProvider(provider, 10*time.Second, 500*time.Millisecond)
	if !ready {
		println("Provider not ready after waiting")
		return
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

	nc, err := natutils.GetNatsConn()
	defer func(nc *nats.Conn) {
		err := nc.Drain()
		if err != nil {
		}
	}(nc)
	kv, idempotentCtx, err := natutils.NewKeyValueStore(jetstream.KeyValueConfig{Bucket: "idempotent_requests", TTL: time.Hour * 24})
	if err != nil {
		println(err.Error(), "error getting new key value store")
		return
	}
	supertokens.DebugEnabled = true
	e := echo.New()
	e.Use(middleware.Logger())
	// Middleware
	e.Use(appmiddleware.RequestCounterMiddleware)
	e.Use(appmiddleware.AddRouteToCTXMiddleware)
	// If load is too high, fail before we process anything else. this may need to be moved after logging
	e.Use(echo.WrapMiddleware(loadshedhttp.NewHandlerMiddleware(appmiddleware.CreateShedder(), loadshedhttp.HandlerOptionCallback(&appmiddleware.RejectionHandler{}))))
	e.Use(middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(100)))
	e.Use(appmiddleware.IdempotencyMiddleware(kv, idempotentCtx))
	e.Use(appmiddleware.TracingMiddleware())
	e.Use(middleware.Recover())
	auth.ApplyAuthAndCorsMiddleware(e)
	routes.Route(e)
	ctx, stop = signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()
	// Start server
	go func() {
		if err := e.Start(":4000"); err != nil && !errors.Is(err, http.ErrServerClosed) {
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

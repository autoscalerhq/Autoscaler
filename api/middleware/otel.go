package appmiddleware

import (
	"github.com/labstack/echo/v4"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/codes"
)

// TracingMiddleware This middleware is used to start and end a tracing span for each HTTP request
func TracingMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		tracer := otel.Tracer("echo-server")
		ctx, span := tracer.Start(c.Request().Context(), c.Path())
		defer span.End()
		req := c.Request().WithContext(ctx)
		c.SetRequest(req)
		err := next(c)

		// Set the span status based on the outcome
		if err != nil {
			span.SetStatus(codes.Error, err.Error())
			span.RecordError(err)
		} else {
			span.SetStatus(codes.Ok, "OK")
		}

		return err
	}
}

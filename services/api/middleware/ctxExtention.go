package middleware

import (
	"context"
	"github.com/labstack/echo/v4"
)

type ctxKeyType string

const ctxKey ctxKeyType = "urlPath"

type URLClassifier struct{}

func AddRouteToCTXMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// Extract URL path and add it to context
		ctx := context.WithValue(c.Request().Context(), ctxKey, c.Request().URL.Path)
		// Replace the request context
		c.SetRequest(c.Request().WithContext(ctx))
		// Proceed with the next handler
		return next(c)
	}
}

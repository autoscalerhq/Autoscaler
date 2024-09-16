package middleware

import (
	"github.com/labstack/echo/v4"
	"sync/atomic"
)

var activeRequests int64

func RequestCounterMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		atomic.AddInt64(&activeRequests, 1)
		defer atomic.AddInt64(&activeRequests, -1)
		return next(c)
	}
}

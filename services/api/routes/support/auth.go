package support

import (
	"github.com/autoscalerhq/autoscaler/services/api/middleware"
	"github.com/labstack/echo/v4"
)

func MakeProtectedRouteGroup(e *echo.Echo, middlewareParams middleware.MiddlewareParams) *echo.Group {
	authRoutes := e.Group("")
	middleware.ApplyProtectedMiddleware(authRoutes, middlewareParams)
	return authRoutes
}

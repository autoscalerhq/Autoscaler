package support

import (
	"github.com/autoscalerhq/autoscaler/services/api/auth"
	"github.com/autoscalerhq/autoscaler/services/api/middleware"
	"github.com/labstack/echo/v4"
)

func MakeProtectedRouteGroup(e *echo.Echo) *echo.Group {
	authRoutes := e.Group("")
	authRoutes.Use(echo.WrapMiddleware(auth.VerifySessionMiddleware))
	authRoutes.Use(middleware.CorsMiddleware())
	return authRoutes
}

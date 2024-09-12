package routes

import (
	"github.com/autoscalerhq/autoscaler/services/api/middleware"
	"github.com/autoscalerhq/autoscaler/services/api/routes/support"
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func Route(e *echo.Echo, middlewareParams middleware.MiddlewareParams) {
	// public routes
	e.GET("/health-check", HealthCheck)
	// swag init -g ./main.go --output ./docs
	e.GET("/swagger/*", echoSwagger.WrapHandler)
	// protected routes
	protectedRoutes := support.MakeProtectedRouteGroup(e, middlewareParams)
	protectedRoutes.GET("/comment", CommentRoute)
}

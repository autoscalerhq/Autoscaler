package routes

import (
	"github.com/autoscalerhq/autoscaler/services/api/routes/support"
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func Route(e *echo.Echo) {
	// public routes
	e.GET("/health-check", HealthCheck)
	// swag init -g ./main.go --output ./docs
	e.GET("/swagger/*", echoSwagger.WrapHandler)
	// protected routes
	protectedRoutes := support.MakeProtectedRouteGroup(e)
	protectedRoutes.GET("/comment", CommentRoute)
}

package routes

import (
	"github.com/autoscalerhq/autoscaler/services/api/auth"
	"github.com/labstack/echo/v4"
)

func CommentRoute(c echo.Context) error {
	userID := auth.GetUserId(c)
	return c.String(200, "Hello "+userID)
}

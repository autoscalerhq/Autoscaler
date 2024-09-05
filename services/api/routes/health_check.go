package routes

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

// HealthCheck godoc
// @Summary Health check endpoint
// @Description Check if the api is up and running
// @Tags health
// @Produce json
// @Success 200 {object} map[string]interface{}+
// @Router / [get]
func HealthCheck(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": "Server is up and running",
	})
}

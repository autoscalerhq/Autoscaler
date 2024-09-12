package middleware

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/supertokens/supertokens-golang/supertokens"
)

func CorsMiddleware() echo.MiddlewareFunc {
	supertokensAllowedHeaders := supertokens.GetAllCORSHeaders()
	return middleware.CORSWithConfig(middleware.CORSConfig{
		Skipper:      middleware.DefaultCORSConfig.Skipper,
		AllowOrigins: []string{"http://localhost:3000"},
		AllowMethods: middleware.DefaultCORSConfig.AllowMethods,
		AllowHeaders: append(
			[]string{
				"Content-Type",
			},
			supertokensAllowedHeaders...,
		),
		AllowCredentials: true,
	})
}

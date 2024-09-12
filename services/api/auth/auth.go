package auth

import (
	appmiddleware "github.com/autoscalerhq/autoscaler/services/api/middleware"
	"github.com/labstack/echo/v4"
	"github.com/supertokens/supertokens-golang/recipe/session"
	"github.com/supertokens/supertokens-golang/supertokens"
)

func GetUserId(c echo.Context) string {
	sessionContainer := session.GetSessionFromRequestContext(c.Request().Context())
	if sessionContainer == nil {
		return ""
	}
	userID := sessionContainer.GetUserID()
	return userID
}

func ApplyAuthAndCorsMiddleware(e *echo.Echo) {
	// cors must be applied before the supertokens middleware, otherwise the supertokens middleware will not be able to set the appropriate headers
	e.Use(appmiddleware.CorsMiddleware())
	e.Use(echo.WrapMiddleware(supertokens.Middleware))
}

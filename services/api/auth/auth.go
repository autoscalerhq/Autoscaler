package auth

import (
	"github.com/labstack/echo/v4"
	"github.com/supertokens/supertokens-golang/recipe/session"
)

func GetUserId(c echo.Context) string {
	sessionContainer := session.GetSessionFromRequestContext(c.Request().Context())
	if sessionContainer == nil {
		return ""
	}
	userID := sessionContainer.GetUserID()
	return userID
}

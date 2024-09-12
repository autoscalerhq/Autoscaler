package appmiddleware

import (
	"github.com/labstack/echo/v4"
	"github.com/supertokens/supertokens-golang/supertokens"
)

func ApplyAuthAndCorsMiddleware(e *echo.Echo) {
	// cors must be applied before the supertokens middleware, otherwise the supertokens middleware will not be able to set the appropriate headers
	e.Use(CorsMiddleware())
	e.Use(echo.WrapMiddleware(supertokens.Middleware))
}

package routes

import (
	"mrsydar/apiserver/configs/auth0"
	"mrsydar/apiserver/controllers"

	"github.com/labstack/echo/v4"
)

func ApplyCallback(e *echo.Echo) {
	e.GET(auth0.CallbackEndpoint, controllers.FetchJWTToken)
}

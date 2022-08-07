package routes

import (
	"mrsydar/apiserver/configs/auth0"
	"mrsydar/apiserver/controllers"
	d2middleware "mrsydar/apiserver/middleware"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func ApplyAccount(e *echo.Echo) {
	e.GET("/account", controllers.GetAccount,
		middleware.JWTWithConfig(middleware.JWTConfig{SigningKey: []byte(auth0.Secret)}),
		d2middleware.AssociateAccountWithRequest,
	)
}

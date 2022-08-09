package routes

import (
	"mrsydar/apiserver/configs/auth0"
	"mrsydar/apiserver/controllers"
	d2middlewares "mrsydar/apiserver/middlewares"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func ApplyNotes(e *echo.Echo) {
	e.GET("/notes", controllers.GetNotes,
		middleware.JWTWithConfig(middleware.JWTConfig{SigningKey: []byte(auth0.Secret)}),
		d2middlewares.AssociateAccountWithRequest,
	)

	e.POST("/notes", controllers.PostNote,
		middleware.JWTWithConfig(middleware.JWTConfig{SigningKey: []byte(auth0.Secret)}),
		d2middlewares.AssociateAccountWithRequest,
	)
}

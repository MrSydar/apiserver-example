package main

import (
	_ "mrsydar/apiserver/configs"
	"mrsydar/apiserver/routes"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	e.Use(middleware.Recover())
	e.Use(middleware.Logger())

	routes.ApplyCallback(e)
	routes.ApplyNotes(e)

	e.Logger.Fatal(e.Start(":9000"))
}

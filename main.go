package main

import (
	"spells/cmd/handlers"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.GET("/", handlers.Home)
	e.Logger.Fatal(e.Start(":8080"))
}

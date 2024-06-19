package main

import (
	"spells/cmd/handlers"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", handlers.Home)
	e.GET("/spells", handlers.GetAllSpells)
	e.GET("/spells/:spellId", handlers.GetSpell)
	e.POST("/spells", handlers.CreateSpell)
	e.PUT("/spells/:spellId", handlers.UpdateSpell)
	e.DELETE("/spells/:spellID", handlers.DeleteSpell)
	e.Logger.Fatal(e.Start(":8080"))
}

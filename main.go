package main

import (
	"github.com/MYK12397/handlers"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", handlers.HomeHandler)
	e.POST("/data", handlers.PostDataHandler)
	e.GET("/data", handlers.GetAllDataHandler)
	e.Logger.Fatal(e.Start(":8081"))

}

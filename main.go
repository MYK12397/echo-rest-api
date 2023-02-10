package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Data struct {
	UserId int
	Id     int
	Title  string
	Body   string
}

func main() {

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", homeHandler)
	e.POST("/data", postHandler)
	e.Logger.Fatal(e.Start(":8081"))

}

func homeHandler(c echo.Context) error {

	return c.String(http.StatusOK, "This is a Home Page!")
}

func postHandler(c echo.Context) error {

	data := Data{UserId: 1, Id: 12, Title: "User-1", Body: "Blockchain"}
	return c.JSON(http.StatusOK, data)
}

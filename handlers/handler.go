package handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type Users struct {
	Data []Data
}
type Data struct {
	UserId int
	Id     int
	Title  string
	Body   string
}

func HomeHandler(c echo.Context) error {

	return c.String(http.StatusOK, "This is a Home Page!")
}

func PostDataHandler(c echo.Context) error {

	data := Data{UserId: 1, Id: 12, Title: "User-1", Body: "Blockchain"}
	return c.JSON(http.StatusOK, data)
}

func GetAllDataHandler(c echo.Context) error {

	r, err := ioutil.ReadFile("data.json")

	if err != nil {
		c.String(http.StatusBadGateway, "unable to process data")
	}
	var user Users

	err = json.Unmarshal(r, &user.Data)

	if err != nil {
		c.String(http.StatusBadGateway, "unable to process data")
	}
	res := make(map[string]interface{})
	res["status"] = "ok"
	res["data"] = user

	return c.JSON(http.StatusOK, res)
}

func GetDataByIdHandler(c echo.Context) error {

	r, err := ioutil.ReadFile("data.json")

	id := c.Param("id")
	ID, _ := strconv.Atoi(id)
	if err != nil {
		c.String(http.StatusBadGateway, "unable to process data")
	}
	var user Users

	err = json.Unmarshal(r, &user.Data)

	if err != nil {
		c.String(http.StatusBadGateway, "unable to process data")
	}
	var ans Users

	for _, doc := range user.Data {
		if doc.Id == ID {
			ans.Data = append(ans.Data, doc)
		}
	}

	res := make(map[string]interface{})
	res["status"] = "ok"
	res["data"] = ans
	return c.JSON(http.StatusOK, res)
}

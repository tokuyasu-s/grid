package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func Index(c echo.Context) error {
	return c.String(http.StatusOK, "hello world")
}

func GetUser(c echo.Context) error {
	id := c.Param("id")
	return c.String(http.StatusOK, id)
}

func Show(c echo.Context) error {
	team := c.QueryParam("team")
	member := c.QueryParam("member")
	return c.String(http.StatusOK, "team"+team+" + member"+member)
}

func Save(c echo.Context) error {
	// Get name and email
	name := c.FormValue("name")
	email := c.FormValue("email")
	return c.String(http.StatusOK, "name:"+name+", email:"+email)
}

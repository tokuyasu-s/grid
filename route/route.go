package route

import (
	"grid/controller"

	"github.com/labstack/echo/v4"
)

func Init() *echo.Echo {
	e := echo.New()
	e.GET("/", controller.Index)
	e.GET("/user/:id", controller.GetUser)
	e.GET("/show", controller.Show)
	e.POST("/save", controller.Save)

	return e
}

package routes

import (
	"restaurant/controllers"

	"github.com/labstack/echo/v4"
)

func InitRoute() *echo.Echo {
	e := echo.New()
	e.GET("/restaurants", controllers.GetRestaurantController)
	e.POST("/restaurants", controllers.InsertRestaurant)
	return e
}

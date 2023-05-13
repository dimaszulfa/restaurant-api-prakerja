package routes

import (
	"net/http"
	"restaurant/controllers"
	"restaurant/models"

	"github.com/labstack/echo/v4"
)

func InitRoute() *echo.Echo {
	e := echo.New()

	//Restaurants
	e.GET("/restaurants", controllers.GetRestaurantController)
	e.GET("/restaurants/:id", func(c echo.Context) error {
		id := c.Param("id")

		restaurant, err := controllers.GetRestaurantById(id)
		if err != nil {
			return c.JSON(http.StatusNotFound, models.BaseResponse{
				Status: false, Message: "Restaurant not found!", Data: nil,
			})
		}

		return c.JSON(http.StatusOK, restaurant)
	})
	e.POST("/restaurants", controllers.InsertRestaurant)
	e.POST("/menus/:id", controllers.InsertNewMenuByRestaurantId)

	//Menu

	return e
}

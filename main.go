package main

import (
	"os"
	"restaurant/configs"
	"restaurant/routes"

	"github.com/labstack/echo/v4"
)

type BaseResponse struct {
	Status  bool        `json:"true"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func main() {
	configs.ConnectDatabase()
	e := routes.InitRoute()
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	e.Start(":" + port)

}

func GetRestaurantsController(c echo.Context) error {
	return c.JSON(200, BaseResponse{
		true, "success", "hello",
	})
}

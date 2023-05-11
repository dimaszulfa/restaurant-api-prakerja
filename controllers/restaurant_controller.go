package controllers

import (
	"fmt"
	"net/http"

	"restaurant/configs"
	"restaurant/models"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func GetRestaurantController(c echo.Context) error {
	var restaurant []models.Restaurant
	// var restaurants []models.Restaurant

	// result := configs.DB.Find(&restaurant)
	result := configs.DB.Model(&models.Restaurant{}).Preload("MenuItem").Find(&restaurant)

	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, models.BaseResponse{
			Status: false, Message: "Gagal get data restaurant dari database", Data: nil,
		})
	}

	return c.JSON(200, models.BaseResponse{
		Status: true, Message: "success", Data: restaurant,
	})
}

func InsertRestaurant(c echo.Context) error {
	var insertUser models.Restaurant
	c.Bind(&insertUser)
	fmt.Println(insertUser)

	// logic bisnis
	// di cek database ada ?
	// 409
	result := configs.DB.First(&models.Restaurant{}, "name = ?", insertUser.Name)
	if result.Error != gorm.ErrRecordNotFound {
		return c.JSON(http.StatusConflict, models.BaseResponse{
			Status: false, Message: "Restaurant telah ada", Data: nil,
		})
	}

	// masukkan ke database
	result = configs.DB.Create(&insertUser)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, models.BaseResponse{
			Status: false, Message: "Gagal insert ke database", Data: nil,
		})
	}
	return c.JSON(http.StatusOK, models.BaseResponse{
		Status: true, Message: "berhasil insert database restaurant", Data: insertUser,
	})
}

package controllers

import (
	"net/http"

	"restaurant/configs"
	"restaurant/models"
	"strconv"

	"github.com/labstack/echo/v4"
)

func InsertNewMenuByRestaurantId(c echo.Context) error {
	// Get the user ID from the request URL parameter

	idStr := c.Param("id")

	// Get the request body
	var menu models.MenuItem
	if err := c.Bind(&menu); err != nil {
		// Handle the error (bad request)
		return c.String(http.StatusBadRequest, "Invalid request")
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.String(http.StatusBadRequest, "Invalid ID")
	}
	menu.RestaurantID = uint(id)

	err = insertMenu(&menu)
	if err != nil {
		return c.String(http.StatusInternalServerError, "Failed to insert menu")
	}

	return c.String(http.StatusOK, "menu inserted successfully")
}

func insertMenu(menuItem *models.MenuItem) error {
	result := configs.DB.Create(menuItem)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

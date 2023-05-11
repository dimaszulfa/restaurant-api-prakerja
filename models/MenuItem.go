package models

type MenuItem struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	Price        int    `json:"price"`
	Category     string `json:"category"`
	RestaurantID uint   `json:"restaurant_id"`
}

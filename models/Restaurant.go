package models

type Restaurant struct {
	ID          int        `json:"id"`
	Name        string     `json:"name"`
	Description string     `json:"description"`
	Address     string     `json:"address"`
	Cuisine     string     `json:"cuisine"`
	Rating      float32    `json:"rating"`
	MenuItem    []MenuItem `gorm:"foreignKey:RestaurantID" json:"menus"`
}

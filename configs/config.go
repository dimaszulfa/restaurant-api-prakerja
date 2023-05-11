package configs

import (
	"fmt"
	"os"
	"restaurant/models"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	loadenv()

	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbUser, dbPassword, dbHost, dbPort, dbName)

	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Init database failed")
	}
	migratDatabase()

	// Connection established
	fmt.Println("Connected to MySQL database!")
}

func loadenv() {
	err := godotenv.Load()
	if err != nil {
		panic("Failed load env file")
	}
}

func migratDatabase() {
	DB.AutoMigrate(&models.Restaurant{}, &models.MenuItem{})

}

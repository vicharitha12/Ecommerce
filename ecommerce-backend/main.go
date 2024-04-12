package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"ecommerce-backend/models"
	"ecommerce-backend/routes"
)

func main() {
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("Failed to connect to database")
	}
	defer db.Close()

	db.AutoMigrate(&models.User{}, &models.Item{}, &models.Cart{}, &models.Order{})

	router := gin.Default()
	routes.SetupRoutes(router, db)

	router.Run(":8080")
}
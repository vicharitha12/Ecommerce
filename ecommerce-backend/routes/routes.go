package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"ecommerce-backend/controllers"
)

func SetupRoutes(router *gin.Engine, db *gorm.DB) {
	// User routes
	router.POST("/users", controllers.CreateUser)
	router.GET("/users", controllers.ListUsers)
	router.POST("/users/login", controllers.LoginUser)

	// Cart routes
	router.POST("/carts", controllers.CreateCart)
	router.GET("/carts/:id", controllers.GetCart)
	router.PUT("/carts/:id", controllers.UpdateCart)
	router.DELETE("/carts/:id", controllers.DeleteCart)

	// Order routes
	router.POST("/orders", controllers.CreateOrder)
	router.GET("/orders/:id", controllers.GetOrder)
	router.PUT("/orders/:id", controllers.UpdateOrder)
	router.DELETE("/orders/:id", controllers.DeleteOrder)

	// Additional routes for cart items can be added here
}
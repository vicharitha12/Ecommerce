package controllers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	_"github.com/jinzhu/gorm"
	"ecommerce-backend/models"
)

func CreateUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create user logic (e.g., save user to database)
	c.JSON(http.StatusOK, gin.H{"message": "User created successfully"})
}

func ListUsers(c *gin.Context) {
	// List all users logic (e.g., fetch users from database)
	users := []models.User{} // Fetch users from database
	c.JSON(http.StatusOK, users)
}

func LoginUser(c *gin.Context) {
	// User login logic (e.g., validate credentials)
	c.JSON(http.StatusOK, gin.H{"message": "User logged in successfully"})
}
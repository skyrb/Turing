package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// User represents a user entity
type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// createUser creates a new user
func createUser(c *gin.Context) {
	var user User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Simulate database insertion
	if user.ID == 0 {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Invalid user ID",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "User created successfully",
		"user":    user,
	})
}

func main() {
	router := gin.Default()
	router.POST("/users", createUser)

	router.Run(":8080")
}

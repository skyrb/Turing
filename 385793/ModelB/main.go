package main

import (
	"github.com/gin-gonic/gin"
)

// Define your custom error struct
type APIError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func someHandler(c *gin.Context) {
	// Simulate some error scenario
	// In a real application, this could be database errors, validation errors, or any other relevant error.
	if someCondition {
		c.JSON(400, APIError{
			Code:    400,
			Message: "Invalid request parameters",
		})
		return
	}

	// If no error, proceed with the handler logic
	result := "Some result from your handler"
	c.JSON(200, gin.H{"result": result})
}

func main() {
	r := gin.Default()
	r.GET("/some-endpoint", someHandler)

	err := r.Run(":8080")
	if err != nil {
		panic(err)
	}
}

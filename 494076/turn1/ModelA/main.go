package main

import (
	"fmt"
)

// CustomError is a user-defined error type with custom formatting.
type CustomError struct {
	message string
	args    []interface{}
}

// Error implements the error interface to provide a formatted error message.
func (e *CustomError) Error() string {
	return fmt.Sprintf(e.message, e.args...)
}

// NewCustomError creates a new CustomError with a formatted message.
func NewCustomError(format string, args ...interface{}) error {
	return &CustomError{message: format, args: args}
}

func main() {
	// Create a custom error with formatted message
	err := NewCustomError("File '%s' not found with error: %v", "example.txt", fmt.Errorf("path not valid"))

	// Print the formatted error message
	fmt.Println(err)
}
package main

import (
	"fmt"
	"runtime/debug"
)

// MyError is a custom error type that implements the error interface.
type MyError struct {
	Err     error
	Message string
	Stack   []byte
	Meta    map[string]string
}

func (e *MyError) Error() string {
	// Use fmt.Sprintf to format the error message with the included fields
	return fmt.Sprintf("error: %s, message: %s, stack trace: %s, metadata: %v",
		e.Err.Error(), e.Message, string(e.Stack), e.Meta)
}

// NewMyError creates a new MyError instance with the given parameters.
func NewMyError(err error, message string, meta map[string]string) error {
	return &MyError{
		Err:     err,
		Message: message,
		Stack:   debug.Stack(), // Capture the current stack trace
		Meta:    meta,
	}
}

func main() {
	// Example usage:
	err := someFunction()
	if err != nil {
		fmt.Println(err) // Output will include the stack trace and metadata
	}
}

func someFunction() error {
	// Simulate an error condition
	innerErr := fmt.Errorf("inner error occurred")

	// Create metadata for the error
	meta := map[string]string{
		"user_id": "123",
		"request_id": "456",
	}

	// Return the custom MyError with detailed information
	return NewMyError(innerErr, "something went wrong", meta)
}
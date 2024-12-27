package main

import (
	"fmt"
	"runtime"
)

// MyError is a custom error type that includes a reason, metadata, and a stack trace.
type MyError struct {
	Reason   string
	Metadata map[string]interface{}
	Stack    string
}

func (e *MyError) Error() string {
	// Basic error message
	msg := fmt.Sprintf("error: %s", e.Reason)

	// Add metadata to the message
	if len(e.Metadata) > 0 {
		metadataStr := fmt.Sprintf(" [metadata: %v]", e.Metadata)
		msg += metadataStr
	}

	// Add stack trace to the message
	if e.Stack != "" {
		msg += fmt.Sprintf("\nstack:\n%s", e.Stack)
	}

	return msg
}

// NewMyError creates a new MyError with the given reason and metadata.
func NewMyError(reason string, metadata map[string]interface{}) *MyError {
	// Get the stack trace
	buf := make([]byte, 1024)
	runtime.Stack(buf, false)
	stack := string(buf)

	return &MyError{
		Reason:   reason,
		Metadata: metadata,
		Stack:    stack,
	}
}

func main() {
	// Create a custom error with metadata and a stack trace
	metadata := map[string]interface{}{
		"userId": "123",
		"apiKey": "***",
	}
	err := NewMyError("data validation failed", metadata)

	// Print the detailed error message
	fmt.Println(err)
}
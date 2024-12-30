package main

import (
	"fmt"
	"time"
)

// CustomError represents a structured error with additional information
type CustomError struct {
	Message      string
	Details      string
	Timestamp    time.Time
	Severity      string // e.g., "error", "warning", "info"
	CorrelationID string
}

// Error implements the error interface
func (e *CustomError) Error() string {
	details := fmt.Sprintf("Details: %s, Timestamp: %s, Severity: %s, CorrelationID: %s", e.Details, e.Timestamp, e.Severity, e.CorrelationID)
	return fmt.Sprintf("%s\n%s", e.Message, details)
}

// NewCustomError creates a new CustomError with specified fields
func NewCustomError(message, details string, timestamp time.Time, severity, correlationID string) *CustomError {
	return &CustomError{Message: message, Details: details, Timestamp: timestamp, Severity: severity, CorrelationID: correlationID}
}

// Example function that demonstrates the use of CustomError
func processData(data []byte, correlationID string) error {
	// Simulate processing
	if len(data) == 0 {
		timestamp := time.Now()
		return NewCustomError("Data processing failed", "Data is empty", timestamp, "error", correlationID)
	}
	
	// Simulate a detailed error
	if string(data) == "invalid" {
		timestamp := time.Now()
		return NewCustomError("Data processing failed", "Input data is invalid", timestamp, "error", correlationID)
	}

	return nil
}

func main() {
	// Generate a correlation ID
	correlationID := "1234567890abcdef"

	// Case 1: Simple error without details
	data1 := []byte{}
	if err := processData(data1, correlationID); err != nil {
		fmt.Println("Error:", err)
	}

	// Case 2: Error with details
	data2 := []byte("invalid")
	if err := processData(data2, correlationID); err != nil {
		fmt.Println("Error:", err)
	}
}
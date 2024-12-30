package main

import (
	"fmt"
	"time"
)

type ErrorSeverity string

const (
	ErrorSeverityCritical ErrorSeverity = "CRITICAL"
	ErrorSeverityError   ErrorSeverity = "ERROR"
	ErrorSeverityWarning ErrorSeverity = "WARNING"
)

type CustomError struct {
	Message       string
	Severity      ErrorSeverity
	Details       string
	CorrelationID string
	Timestamp     time.Time
}

func (e *CustomError) Error() string {
	details := fmt.Sprintf("Details: %s, CorrelationID: %s, Timestamp: %s", e.Details, e.CorrelationID, e.Timestamp)
	return fmt.Sprintf("[%s] %s\n%s", e.Severity, e.Message, details)
}

func NewCustomError(message string, severity ErrorSeverity, details, correlationID string, timestamp time.Time) *CustomError {
	return &CustomError{
		Message:       message,
		Severity:      severity,
		Details:       details,
		CorrelationID: correlationID,
		Timestamp:     timestamp,
	}
}

func processData(data []byte, correlationID string) error {
	// Simulate processing
	if len(data) == 0 {
		timestamp := time.Now()
		return NewCustomError("Data processing failed", ErrorSeverityError, "Data is empty", correlationID, timestamp)
	}

	if string(data) == "invalid" {
		timestamp := time.Now()
		return NewCustomError("Data processing failed", ErrorSeverityWarning, "Input data is invalid", correlationID, timestamp)
	}

	return nil
}

func main() {
	correlationID := "ABCD1234"

	// Case 1: Error with details and correlation ID
	data1 := []byte{}
	if err := processData(data1, correlationID); err != nil {
		fmt.Println("Error:", err)
	}

	// Case 2: Warning with details and correlation ID
	data2 := []byte("invalid")
	if err := processData(data2, correlationID); err != nil {
		fmt.Println("Error:", err)
	}

	// Case 3: Error without details, but still with correlation ID
	data3 := []byte("some-data")
	err := processData(data3, "")
	if err != nil {
		// Use type assertion to get the underlying CustomError
		customErr, ok := err.(*CustomError)
		if ok {
			fmt.Println("Error:", customErr.Message)
			fmt.Println("Severity:", customErr.Severity)
			fmt.Println("Details:", customErr.Details)
			fmt.Println("CorrelationID:", customErr.CorrelationID)
			fmt.Println("Timestamp:", customErr.Timestamp)
		} else {
			fmt.Println("Error:", err)
		}
	}
}

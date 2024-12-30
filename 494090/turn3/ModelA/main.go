package main

import (
	"fmt"
	"sync"
	"time"
)

// CustomError represents a structured error with additional information
type CustomError struct {
	Message       string
	Severity      string
	Details       string
	CorrelationID string
	Timestamp     time.Time
	GoroutineID   int
}

func (e *CustomError) Error() string {
	details := fmt.Sprintf("Details: %s, CorrelationID: %s, GoroutineID: %d, Timestamp: %s", e.Details, e.CorrelationID, e.GoroutineID, e.Timestamp)
	return fmt.Sprintf("[%s] %s\n%s", e.Severity, e.Message, details)
}

func NewCustomError(message, severity string, details, correlationID string, timestamp time.Time, goroutineID int) *CustomError {
	return &CustomError{
		Message:       message,
		Severity:      severity,
		Details:       details,
		CorrelationID: correlationID,
		Timestamp:     timestamp,
		GoroutineID:   goroutineID,
	}
}

var errorBuffer sync.Buffer

func generateErrorMessage(err error, severity string, correlationID string) *CustomError {
	timestamp := time.Now()
	goroutineID := fmt.Sprintf("%d", getGoroutineID())

	errorBuffer.Reset()
	fmt.Fprintf(&errorBuffer, "Error: %v", err)
	customErr := NewCustomError(errorBuffer.String(), severity, "", correlationID, timestamp, goroutineID)

	return customErr
}

func getGoroutineID() int {
	// Note: This is a non-portable way to get a Goroutine ID, only works on some platforms
	var id uintptr
	runtime.GetGoroutineID(&id)
	return int(id)
}

func processData(data []byte, correlationID string) error {
	// Simulate processing
	if len(data) == 0 {
		return fmt.Errorf("Data processing failed: Data is empty")
	}

	if string(data) == "invalid" {
		return fmt.Errorf("Data processing failed: Invalid data")
	}

	return nil
}

func main() {
	correlationID := "ABCD1234"
	var wg sync.WaitGroup

	// Start multiple Goroutines simulating data processing
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			data := []byte("some-data")
			if err := processData(data, correlationID); err != nil {
				customErr := generateErrorMessage(err, "ERROR", correlationID)
				fmt.Println("Error from Goroutine:", customErr)
			}
		}()
	}

	wg.Wait()
}
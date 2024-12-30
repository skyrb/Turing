package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

var (
	errorMutex sync.Mutex
)

// ErrorMessage formats a thread-safe error message with runtime information.
func ErrorMessage(format string, a ...interface{}) string {
	errorMutex.Lock()
	defer errorMutex.Unlock()

	goroutineID := runtime.NumGoroutine()
	timestamp := time.Now().Format(time.RFC3339)

	return fmt.Sprintf("[Goroutine %d - %s] %s", goroutineID, timestamp, fmt.Sprintf(format, a...))
}

func processData(data []byte, ch chan<- error) {
	// Simulate processing
	if len(data) == 0 {
		ch <- fmt.Errorf(ErrorMessage("Data processing failed: Data is empty"))
		return
	}

	if string(data) == "invalid" {
		ch <- fmt.Errorf(ErrorMessage("Data processing failed: Input data is invalid"))
		return
	}

	ch <- nil
}

func main() {
	errorsCh := make(chan error, 10)

	// Start multiple Goroutines to process data
	for i := 0; i < 5; i++ {
		go processData([]byte(fmt.Sprintf("data-%d", i)), errorsCh)
	}

	// Collect and print errors from the Goroutines
	for i := 0; i < 5; i++ {
		if err := <-errorsCh; err != nil {
			fmt.Println("Error:", err)
		}
	}

	close(errorsCh)
}
 
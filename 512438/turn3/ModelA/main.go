package main

import (
	"fmt"
	"io/ioutil"
	"sync"
	"time"
)

const (
	maxRetries     = 3  // Number of retry attempts
	errorBufferSize = 1000 // Adjust the buffer size as needed
)

// RetryableError defines an error that can be retried
type RetryableError struct {
	err   error
	delay time.Duration
}

func (r *RetryableError) Error() string {
	return fmt.Sprintf("retryable error: %v (retrying in %v)", r.err, r.delay)
}

func processFile(filePath string, wg *sync.WaitGroup, errors chan<- error) {
	defer wg.Done()

	retryCount := 0
	for retryCount < maxRetries {
		err := processFileAttempt(filePath, retryCount, errors)
		if err == nil {
			return // Success, no error to process
		}

		// Check if the error is retryable
		if retryErr, ok := err.(*RetryableError); ok {
			time.Sleep(retryErr.delay) // Retry with a delay
			retryCount++
			continue
		}

		break // Unrecoverable error, no further retries
	}

	errors <- err // If no success after retries, log the error
}

func processFileAttempt(filePath string, retryCount int, errors chan<- error) error {
	// Simulate various error conditions
	switch filePath {
	case "file2.txt":
		if retryCount < maxRetries {
			return &RetryableError{err: fmt.Errorf("file is locked for %s, retry in 1s", filePath), delay: 1 * time.Second}
		} else {
			return fmt.Errorf("file is locked for %s, too many retries", filePath)
		}
	case "file4.txt":
		return fmt.Errorf("file not found: %s", filePath) // Non-retryable error
	case "file5.txt":
		if retryCount < maxRetries {
			return &RetryableError{err: fmt.Errorf("read timeout for %s, retry in 2s", filePath), delay: 2 * time.Second}
		} else {
			return fmt.Errorf("read timeout for %s, too many retries", filePath)
		}
	}

	// Simulate file read with a timeout (adjust the duration as needed)
	data, err := readFileWithTimeout(filePath, 2*time.Second)
	if err != nil {
		if retryCount < maxRetries {
			return &RetryableError{err: err, delay: time.Second} // Assume transient error on retry
		} else {
			return err
		}
	}

	// Process the file content (replace this with your actual processing)
	fmt.Printf("Processed file: %s - Content: %s\n", filePath, string(data))
	return nil
}

func readFileWithTimeout(filePath string, timeout time.Duration) ([]byte, error) {
	c := make(chan []byte, 1)
	go func() {
		data, err := ioutil.ReadFile(filePath)
		if err != nil {
			c <- nil
			return
		}
		c <- data
	}()
	select {
	case data := <-c:
		if data == nil {
			return nil, fmt.Errorf("failed to read %s", filePath)
		}
		return data, nil
	case <-time.After(timeout):
		return nil, fmt.Errorf("read timeout for %s", filePath)
	}
}

func logErrors(errors <-chan error, logWg *sync.WaitGroup) {
	defer logWg.Done()
	for err := range errors {
		fmt.Printf("Error: %v\n", err)
	}
}

func main() {
	filePaths := []string{"file1.txt", "file2.txt", "file3.txt", "file4.txt", "file5.txt", "file6.txt", "file7.txt", "file8.txt", "file9.txt", "file10.txt"} // Add more file paths as needed

	var fileWg sync.WaitGroup // WaitGroup for file processing goroutines
	var logWg sync.WaitGroup  // WaitGroup for logging goroutine
	errors := make(chan error, errorBufferSize)

	// Start the logErrors goroutine
	logWg.Add(1)
	go logErrors(errors, &logWg)

	// Start the processFile goroutines
	fileWg.Add(len(filePaths))
	for _, filePath := range filePaths {
		go processFile(filePath, &fileWg, errors)
	}

	// Wait for all file processing to complete
	fileWg.Wait()

	// Close the error channel after all processFile goroutines have finished
	close(errors)

	// Wait for the logErrors goroutine to finish
	logWg.Wait()

	fmt.Println("All files processed.")
}

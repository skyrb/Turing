package main

import (
	"fmt"
	"math/rand"
	"net"
	"net/http"
	"sync"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// Define a custom error interface
type MyError interface {
	error
	IsCritical() bool
}

// Implement the error interface
type NetworkError struct {
	message string
}

func (ne *NetworkError) Error() string {
	return ne.message
}

func (ne *NetworkError) IsCritical() bool {
	return true
}

type DataFormattingError struct {
	message string
}

func (dfe *DataFormattingError) Error() string {
	return dfe.message
}

func (dfe *DataFormattingError) IsCritical() bool {
	return false
}

// fetchData simulates fetching data from a URL. It returns a string response and an error.
func fetchData(url string) (string, error) {
	time.Sleep(time.Duration(rand.Intn(200)) * time.Millisecond)

	// Simulate network request
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// Simulate occasional error
	if rand.Intn(10) == 0 {
		return "", fmt.Errorf("simulated error fetching data from %s", url)
	}

	return fmt.Sprintf("Fetched from %s", url), nil
}

func main() {
	// URLs to fetch data from
	urls := []string{
		"http://example.com",
		"http://google.com",
		"http://github.com",
		"http://stackoverflow.com",
		"http://non-existent-url.com",
	}

	// Make a channel to send error messages
	errorChan := make(chan MyError, len(urls))

	// Synchronization to ensure we wait for all Goroutines to finish
	var wg sync.WaitGroup

	for _, url := range urls {
		wg.Add(1)
		go func(url string) {
			defer wg.Done()
			_, err := fetchData(url)
			if err != nil {
				if _, ok := err.(net.Error); ok {
					// Treat network errors as critical
					errorChan <- &NetworkError{message: err.Error()}
				} else {
					// Treat other errors as minor
					errorChan <- &DataFormattingError{message: err.Error()}
				}
			}
		}(url)
	}

	// Wait for all Goroutines to complete
	wg.Wait()

	// Close the error channel
	close(errorChan)

	// Collect and report errors
	errorsOccurred := false
	for err := range errorChan {
		if err.IsCritical() {
			fmt.Printf("Critical Error: %s\n", err.Error())
			// Handle critical errors immediately
			// You might want to stop processing or retry
		} else {
			fmt.Printf("Minor Error: %s\n", err.Error())
			// Log minor errors or accumulate them for later handling
		}
		errorsOccurred = true
	}

	// Check if there were any errors
	if errorsOccurred {
		fmt.Println("Some errors occurred during data fetch.")
	} else {
		fmt.Println("All data fetched successfully.")
	}
}
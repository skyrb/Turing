package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"sync"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
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
	errorChan := make(chan error, len(urls))

	// Synchronization to ensure we wait for all Goroutines to finish
	var wg sync.WaitGroup

	for _, url := range urls {
		wg.Add(1)
		go func(url string) {
			defer wg.Done()
			_, err := fetchData(url)
			if err != nil {
				errorChan <- err
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
		fmt.Println("Error:", err)
		errorsOccurred = true
	}

	// Check if there were any errors
	if errorsOccurred {
		fmt.Println("Some errors occurred during data fetch.")
	} else {
		fmt.Println("All data fetched successfully.")
	}
}
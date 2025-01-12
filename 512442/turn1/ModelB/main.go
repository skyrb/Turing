package main

import (
	"fmt"
	"net/http"
	"sync"
)

func fetchData(url string) (string, error) {
	// Simulate some network activity
	// time.Sleep(2 * time.Second)

	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// Read the response body
	// body, err := ioutil.ReadAll(resp.Body)
	// if err != nil {
	// 	return "", err
	// }

	return "Success", nil
}

func main() {
	urls := []string{
		"https://www.example.com",
		"https://www.example.com/not-found",
		"https://www.example.com/other",
	}

	// Create a channel to receive errors from the worker goroutines
	errorChan := make(chan error)

	// WaitGroup to ensure the main goroutine waits for all worker goroutines to finish
	var wg sync.WaitGroup

	// Launch a worker goroutine for each URL
	for _, url := range urls {
		wg.Add(1)
		go func(url string) {
			defer wg.Done()
			_, err := fetchData(url)
			if err != nil {
				// Send the error to the error channel
				errorChan <- err
			}
		}(url)
	}

	// Go routine to collect errors from the error channel
	go func() {
		wg.Wait()
		close(errorChan)
	}()

	// Collect and display errors from the error channel
	var errors []error
	for err := range errorChan {
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		fmt.Println("Errors occurred:")
		for _, err := range errors {
			fmt.Printf("- %s\n", err)
		}
	} else {
		fmt.Println("All fetches completed successfully.")
	}
}  
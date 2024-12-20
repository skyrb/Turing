package main

import (
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"sync"
	"time"
)

const (
	timeoutDuration = 2 * time.Second // Default timeout duration for API calls
)

type apiCallInfo struct {
	url       string
	timeout   time.Duration
	response string
	error     error
}

func simulateApiCall(info *apiCallInfo) {
	// Simulate API call with a random delay
	delay := time.Second + time.Duration(rand.Intn(2))*time.Second
	time.Sleep(delay)

	// Make the HTTP GET request
	resp, err := http.Get(info.url)
	if err != nil {
		info.error = err
		return
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		info.error = err
		return
	}
	info.response = string(body)
}

func main() {
	rand.Seed(time.Now().UnixNano()) // Seed the random number generator

	var wg sync.WaitGroup
	var apiCalls []apiCallInfo

	// Define the list of API call information including URLs and individual timeouts
	apiCalls = append(apiCalls, apiCallInfo{
		url:       "https://jsonplaceholder.typicode.com/todos/1",
		timeout:   5 * time.Second,
	}, apiCallInfo{
		url:       "https://jsonplaceholder.typicode.com/todos/2",
		timeout:   3 * time.Second,
	}, apiCallInfo{
		url:       "https://jsonplaceholder.typicode.com/todos/3",
		timeout:   2 * time.Second,
	})

	// Start each API call concurrently
	for _, info := range apiCalls {
		wg.Add(1)
		go func(info apiCallInfo) {
			defer wg.Done() // Ensure goroutine is properly closed
			simulateApiCall(&info)
		}(info)
	}

	// Wait for all API calls to finish
	wg.Wait()

	// Display results of each API call
	for _, info := range apiCalls {
		if info.error != nil {
			fmt.Printf("Error for URL %q: %v\n", info.url, info.error)
		} else if info.response != "" {
			fmt.Printf("Response for URL %q: %s\n", info.url, info.response)
		} else {
			fmt.Printf("No response for URL %q\n", info.url)
		}
	}
}
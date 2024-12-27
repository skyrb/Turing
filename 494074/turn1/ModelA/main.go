package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"sync"
	"time"
)

// RetryFunction represents a function that can be retried.
type RetryFunction func(context.Context) error

// MaxRetries is the maximum number of retry attempts.
const MaxRetries = 5

// RetryWithBackoff performs a function with backoff retries.
func RetryWithBackoff(ctx context.Context, f RetryFunction, backoff time.Duration) error {
	for retryCount := 0; retryCount <= MaxRetries; retryCount++ {
		err := f(ctx)
		if err == nil {
			return nil
		}

		if retryCount < MaxRetries {
			select {
			case <-ctx.Done():
				log.Printf("Retry cancelled: %v", err)
				return ctx.Err()
			case <-time.After(backoff):
				log.Printf("Retry attempt %d: %v", retryCount+1, err)
			}
		} else {
			log.Printf("Maximum retries reached: %v", err)
		}
	}
	return fmt.Errorf("failed after %d retries", MaxRetries)
}

// SimulateNetworkCall simulates a network call that may fail.
func SimulateNetworkCall(ctx context.Context, url string, retryBackoff time.Duration) error {
	select {
	case <-time.After(time.Duration(rand.Intn(500))*time.Millisecond):
		// Simulate a delay
	case <-ctx.Done():
		return ctx.Err()
	}

	if rand.Intn(2) == 0 {
		// Simulate a failure
		return fmt.Errorf("network call to %s failed", url)
	}

	// Simulate a successful response
	log.Printf("Network call to %s succeeded", url)
	return nil
}

// Main function to handle multiple network calls with retries.
func main() {
	var wg sync.WaitGroup
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	urls := []string{
		"http://example.com/api/1",
		"http://example.com/api/2",
		"http://example.com/api/3",
	}

	for _, url := range urls {
		wg.Add(1)
		go func(u string) {
			defer wg.Done()
			if err := RetryWithBackoff(ctx, func(ctx context.Context) error {
				return SimulateNetworkCall(ctx, u, 1*time.Second)
			}, 2*time.Second); err != nil {
				log.Printf("Failed to process URL %s: %v", u, err)
			}
		}(url)
	}

	wg.Wait()
	log.Println("All network calls completed")
}
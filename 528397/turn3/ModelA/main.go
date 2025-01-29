package main

import (
	"errors"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const maxRetries = 3

type CDN struct {
	name string
}

func (cdn *CDN) distributeContent(content string) error {
	// Simulate a failure for the "London" CDN to demonstrate retry and fallback
	if cdn.name == "London" {
		fmt.Printf("Failed to deliver content to %s CDN\n", cdn.name)
		return errors.New("failed delivery")
	}
	fmt.Printf("Successfully delivered content to %s CDN\n", cdn.name)
	return nil
}

func retryWithFallback(content string, primary *CDN, secondary *CDN, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 0; i < maxRetries; i++ {
		err := primary.distributeContent(content)
		if err == nil {
			return
		}
		fmt.Printf("Retry %d for %s...\n", i+1, primary.name)
		time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond) // Random backoff
	}

	fmt.Printf("Switching to secondary CDN: %s.\n", secondary.name)
	if err := secondary.distributeContent(content); err != nil {
		fmt.Printf("Failed to deliver content to secondary CDN %s after retries.\n", secondary.name)
	}
}

func main() {
	var wg sync.WaitGroup

	content := "product images"

	primaryCDN := &CDN{name: "London"}
	secondaryCDN := &CDN{name: "Frankfurt"}

	wg.Add(1)
	go retryWithFallback(content, primaryCDN, secondaryCDN, &wg)

	wg.Wait()
	fmt.Println("Content distribution process completed.")
}
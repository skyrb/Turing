package main

import (
	"fmt"
	"sync"
	"time"
)

// Function to simulate distribution of content to a regional CDN server
func distributeContent(region string, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Printf("Starting content distribution to %s region\n", region)

	// Simulating product image distribution
	distributeProductImages(region)

	// Simulating promotional banners distribution
	distributePromotionalBanners(region)

	fmt.Printf("Completed content distribution to %s region\n", region)
}

// Simulating the distribution of product images
func distributeProductImages(region string) {
	fmt.Printf("Distributing product images to %s region...\n", region)
	// Simulate network delay
	time.Sleep(2 * time.Second)
	fmt.Printf("Product images distributed to %s region\n", region)
}

// Simulating the distribution of promotional banners
func distributePromotionalBanners(region string) {
	fmt.Printf("Distributing promotional banners to %s region...\n", region)
	// Simulate network delay
	time.Sleep(2 * time.Second)
	fmt.Printf("Promotional banners distributed to %s region\n", region)
}

func main() {
	var wg sync.WaitGroup

	regions := []string{"North America", "Europe", "Asia"}

	// Launch a goroutine for each region
	for _, region := range regions {
		wg.Add(1)
		go distributeContent(region, &wg)
	}

	// Wait for all goroutines to complete
	wg.Wait()

	fmt.Println("All content distribution completed.")
}
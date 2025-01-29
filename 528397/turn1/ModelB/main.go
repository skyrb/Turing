package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Product represents a product with its image and promotional banner URLs.
type Product struct {
	ID    int
	Image string
	Banner string
}

// Region represents a region with its CDN server URL.
type Region struct {
	Name    string
	CDNURL string
}

func distributeContent(wg *sync.WaitGroup, product Product, region Region) {
	defer wg.Done()
	// Simulate content distribution time (random delay between 1-5 seconds)
	time.Sleep(time.Duration(rand.Intn(5)+1) * time.Second)

	fmt.Printf("Distributing product %d image to %s CDN\n", product.ID, region.Name)
	fmt.Printf("Distributing product %d banner to %s CDN\n", product.ID, region.Name)
}

func main() {
	rand.Seed(time.Now().UnixNano())
	// List of products to distribute
	products := []Product{
		{ID: 1, Image: "image1.jpg", Banner: "banner1.jpg"},
		{ID: 2, Image: "image2.jpg", Banner: "banner2.jpg"},
		// Add more products as needed
	}

	// List of regions to distribute to
	regions := []Region{
		{Name: "North America", CDNURL: "https://cdn.na.example.com"},
		{Name: "Europe", CDNURL: "https://cdn.eu.example.com"},
		{Name: "Asia", CDNURL: "https://cdn.asia.example.com"},
	}

	var wg sync.WaitGroup

	// Distribute content to each region concurrently
	for _, product := range products {
		for _, region := range regions {
			wg.Add(2)
			go distributeContent(&wg, product, region)
			go distributeContent(&wg, product, region) // Distribute both image and banner concurrently
		}
	}

	// Wait for all content distribution tasks to complete
	wg.Wait()

	fmt.Println("Content distribution completed successfully!")
}
package main

import (
	"fmt"
	"sort"
)

// Define the Product struct
type Product struct {
	Name     string
	Category string
	Price    float64
	Stock    int
	Rating   float64
}

// Filter and sort products
func GetTopRatedAndSortedProducts(products []Product, minRating float64) []Product {
	// Filter products based on rating
	var filteredProducts []Product
	for _, product := range products {
		if product.Rating >= minRating {
			filteredProducts = append(filteredProducts, product)
		}
	}

	// Sort the filtered products by price
	sort.Slice(filteredProducts, func(i, j int) bool {
		return filteredProducts[i].Price < filteredProducts[j].Price
	})

	return filteredProducts
}

func main() {
	// Example list of products
	products := []Product{
		{"Laptop", "Electronics", 999.99, 50, 4.7},
		{"Smartphone", "Electronics", 599.99, 230, 4.4},
		{"Headphones", "Audio", 199.99, 0, 4.8},
		{"Coffee Maker", "Kitchen", 89.99, 100, 3.9},
		{"Monitor", "Electronics", 299.99, 70, 4.6},
	}

	// Get top-rated and price-sorted products
	topRatedProducts := GetTopRatedAndSortedProducts(products, 4.5)

	// Display the result
	fmt.Println("Top Rated Products:")
	for _, product := range topRatedProducts {
		fmt.Printf("Name: %s, Category: %s, Price: %.2f, Rating: %.1f\n",
			product.Name, product.Category, product.Price, product.Rating)
	}
}
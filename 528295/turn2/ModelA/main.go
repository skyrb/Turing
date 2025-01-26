package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
)

// Define the Product and Update structs
type Product struct {
	ID       string  // Assume an ID field for unique identification
	Name     string
	Category string
	Price    float64
	Stock    int
	Rating   float64
}

type Update struct {
	ID    string  `json:"id"`
	Price *float64 `json:"price,omitempty"` // Use pointers to handle zero values explicitly
	Stock *int     `json:"stock,omitempty"`
}

// Parse and apply updates
func ApplyUpdates(products map[string]*Product, updates []Update) error {
	// Create a backup for rollback purposes
	backup := make(map[string]*Product)

	// Save current state
	for id, product := range products {
		copyProduct := *product // Make a copy
		backup[id] = &copyProduct
	}

	// Track errors during update
	var errMessages []string

	// Process each update
	for _, update := range updates {
		// Fetch the product
		product, exists := products[update.ID]

		if !exists {
			errMessages = append(errMessages, fmt.Sprintf("Product ID %s not found", update.ID))
			continue
		}

		// Validate update values
		if update.Price != nil && *update.Price < 0 {
			errMessages = append(errMessages, fmt.Sprintf("Invalid Price for Product ID %s", update.ID))
			continue
		}

		// Apply updates if they are valid
		if update.Price != nil {
			product.Price = *update.Price
		}
		if update.Stock != nil {
			product.Stock = *update.Stock
		}
	}

	// If there are errors, rollback and report
	if len(errMessages) > 0 {
		// Roll back to original state
		for id, oldProduct := range backup {
			products[id] = oldProduct
		}

		// Log errors and return composite error
		for _, msg := range errMessages {
			log.Println("Error:", msg)
		}
		return errors.New("update process failed with errors")
	}

	return nil
}

func main() {
	// Example product catalog
	productCatalog := map[string]*Product{
		"1": {"1", "Laptop", "Electronics", 999.99, 50, 4.7},
		"2": {"2", "Smartphone", "Electronics", 599.99, 230, 4.4},
		"3": {"3", "Monitor", "Electronics", 299.99, 70, 4.6},
	}

	// JSON payload representing a batch update
	jsonPayload := `[{"id": "1", "price": 899.99}, {"id": "2", "stock": 210}, {"id": "4", "price": -100}]`

	// Unmarshal JSON data into updating objects
	var updates []Update
	err := json.Unmarshal([]byte(jsonPayload), &updates)
	if err != nil {
		log.Fatal("Error parsing JSON payload: ", err)
	}

	// Apply updates
	err = ApplyUpdates(productCatalog, updates)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println("Updates applied successfully.")
		for id, product := range productCatalog {
			fmt.Printf("ID: %s, Name: %s, Price: %.2f, Stock: %d\n", id, product.Name, product.Price, product.Stock)
		}
	}
}
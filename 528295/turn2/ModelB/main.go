package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
)

// Define the Product struct
type Product struct {
	Name     string
	Category string
	Price    float64
	Stock    int
	Rating   float64
}

// Define the Update struct
type Update struct {
	Name  string
	Price float64
	Stock int
}

// Update product catalog
func UpdateProductCatalog(products []Product, updates []Update) error {
	updatedProducts := make(map[string]Product)

	// Validate and update products
	for _, update := range updates {
		if err := ValidateUpdate(update); err != nil {
			return err // Return error on invalid update
		}

		// Find the product to update
		for i, product := range products {
			if product.Name == update.Name {
				// Apply the update
				products[i].Price = update.Price
				products[i].Stock = update.Stock

				// Add to updatedProducts for rollback
				updatedProducts[product.Name] = product
				break
			}
		}
	}

	return nil
}

// Validate the update data
func ValidateUpdate(update Update) error {
	if update.Price < 0 {
		return errors.New("Price must be non-negative")
	}
	// Add more validation rules as needed
	return nil
}

// Rollback function
func RollbackUpdates(products []Product, updatedProducts map[string]Product) {
	for name, updatedProduct := range updatedProducts {
		for i, product := range products {
			if product.Name == name {
				// Undo the update
				products[i].Price = updatedProduct.Price
				products[i].Stock = updatedProduct.Stock
				break
			}
		}
	}
}

func main() {
	// Example list of products
	products := []Product{
		{"Laptop", "Electronics", 999.99, 50, 4.7},
		{"Smartphone", "Electronics", 599.99, 230, 4.4},
		{"Headphones", "Audio", 199.99, 0, 4.8},
	}

	// Example JSON payload with updates
	updatePayload := `[
		{"Name": "Laptop", "Price": 1099.99, "Stock": 30},
		{"Name": "Smartphone", "Price": 549.99, "Stock": 200},
		{"Name": "InvalidProduct", "Price": 100, "Stock": -10}
	]`

	// Parse the JSON payload
	var updates []Update
	if err := json.Unmarshal([]byte(updatePayload), &updates); err != nil {
		log.Fatal("Error parsing JSON payload:", err)
	}

	// Perform the updates with rollback
	if err := UpdateProductCatalog(products, updates); err != nil {
		log.Println("Error updating product catalog:", err)
		// Rollback the changes
		RollbackUpdates(products, map[string]Product{
			"Laptop":       products[0],
			"Smartphone":   products[1],
			"InvalidProduct": products[2], // Assume we know the initial state before update
		})
		log.Println("Rollback successful. Product catalog restored to previous state.")
	} else {
		log.Println("Product catalog updated successfully.")
	}

	// Display the updated products
	fmt.Println("Updated Products:")
	for _, product := range products {
		fmt.Printf("Name: %s, Price: %.2f, Stock: %d\n",
			product.Name, product.Price, product.Stock)
	}
}
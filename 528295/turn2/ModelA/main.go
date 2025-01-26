package main

import (
    "encoding/json"
    "errors"
    "fmt"
    "log"
)

type Product struct {
    ID       int
    Name     string
    Category string
    Price    float64
    Stock    int
    Rating   float64
}

type ProductUpdate struct {
    ID    int
    Price *float64 `json:"price,omitempty"`
    Stock *int     `json:"stock,omitempty"`
}

func main() {
    // Simulated product catalog
    productCatalog := map[int]Product{
        1: {ID: 1, Name: "Product 1", Category: "Cat1", Price: 10.99, Stock: 100, Rating: 4.7},
        2: {ID: 2, Name: "Product 2", Category: "Cat2", Price: 5.99, Stock: 50, Rating: 3.8},
    }

    // Example JSON payload for bulk updates
    jsonPayload := `[
        {"id": 1, "price": 12.99, "stock": 110},
        {"id": 2, "price": -4.99, "stock": 45}
    ]`

    updates := []ProductUpdate{}
    err := json.Unmarshal([]byte(jsonPayload), &updates)
    if err != nil {
        log.Fatalf("Error unmarshalling JSON: %v", err)
    }

    if err := applyUpdates(productCatalog, updates); err != nil {
        log.Printf("Failed to apply updates: %v", err)
    }

    fmt.Println("Updated Product Catalog:")
    for _, product := range productCatalog {
        fmt.Printf("%+v\n", product)
    }
}

func applyUpdates(productCatalog map[int]Product, updates []ProductUpdate) error {
    // Backup original state for rollback
    originalState := make(map[int]Product, len(productCatalog))
    for id, p := range productCatalog {
        originalState[id] = p
    }

    for _, update := range updates {
        // Validate and apply each update
        if err := applyUpdate(productCatalog, update); err != nil {
            log.Printf("Error applying update for product ID %d: %v", update.ID, err)
            // Rollback to original state due to error
            rollback(productCatalog, originalState)
            return errors.New("rollback: changes reverted due to an error")
        }
    }
    return nil
}

func applyUpdate(productCatalog map[int]Product, update ProductUpdate) error {
    product, exists := productCatalog[update.ID]
    if !exists {
        return fmt.Errorf("product with ID %d not found", update.ID)
    }

    if update.Price != nil {
        if *update.Price < 0 {
            return fmt.Errorf("invalid price for product ID %d: must be non-negative", update.ID)
        }
        product.Price = *update.Price
    }

    if update.Stock != nil {
        product.Stock = *update.Stock
    }

    productCatalog[update.ID] = product
    return nil
}

func rollback(productCatalog, originalState map[int]Product) {
    for id, p := range originalState {
        productCatalog[id] = p
    }
}
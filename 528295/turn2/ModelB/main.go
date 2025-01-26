package main  
import (  
    "encoding/json"
    "fmt"
    "log"
    "strings"
)

type Product struct {  
    Name     string
    Category string
    Price    float64
    Stock    int
    Rating   float64
}

type ProductUpdate struct {  
    ID       int `json:"id"`
    Price    float64 `json:"price"`
    Stock    int `json:"stock"`
}

func main() {  
    // Sample data with tens of thousands of entries
    products := []Product{
        {Name: "Product 1", Category: "Cat1", Price: 10.99, Stock: 100, Rating: 4.7},
        {Name: "Product 2", Category: "Cat2", Price: 5.99, Stock: 50, Rating: 3.8},
        // Add more product entries here
    }

    // Sample JSON payload for bulk updates
    updatePayload := `[
        {"id": 1, "price": 12.99, "stock": 150},
        {"id": 2, "price": -3.00, "stock": 30},
        {"id": 3, "price": 7.99, "stock": -20}
    ]`

    var updates []ProductUpdate
    if err := json.Unmarshal([]byte(updatePayload), &updates); err != nil {
        log.Fatalf("Error decoding JSON payload: %v", err)
    }

    // Process updates in bulk and validate data
    if err := processBulkUpdates(products, updates); err != nil {
        log.Printf("Error processing updates: %v", err)
        // Implement rollback mechanism here (e.g., use a transaction or maintain a backup of the original data)
    }

    // Display updated products
    for _, product := range products {  
        fmt.Printf("%s - $%.2f - Stock: %d\n", product.Name, product.Price, product.Stock)
    }
}

func processBulkUpdates(products []Product, updates []ProductUpdate) error {  
    var errors []string

    // Create a map for efficient product lookup by ID
    productMap := make(map[int]int)
    for i, product := range products {  
        productMap[product.ID] = i
    }

    for _, update := range updates {  
        // Validate update data
        if update.Price < 0 {
            errors = append(errors, fmt.Sprintf("Invalid price for Product %d: $%.2f", update.ID, update.Price))
            continue
        }
        if update.Stock < 0 {
            errors = append(errors, fmt.Sprintf("Invalid stock for Product %d: %d", update.ID, update.Stock))
            continue
        }

        // Find the corresponding product in the original array
        if productIndex, ok := productMap[update.ID]; ok {
            products[productIndex].Price = update.Price
            products[productIndex].Stock = update.Stock
        } else {
            errors = append(errors, fmt.Sprintf("Product ID %d not found", update.ID))
        }
    }

    if len(errors) > 0 {
        return fmt.Errorf("Found %d errors: %s", len(errors), strings.Join(errors, ", "))
    }
    return nil
}
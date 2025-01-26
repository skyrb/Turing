
package main

import (
    "encoding/json"
    "errors"
    "fmt"
    "log"
    "sort"
    "sync"
    "time"
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

type Filter struct {
    Category     string
    MinPrice     *float64
    MaxPrice     *float64
    InStockOnly bool
}

type Page struct {
    PageNumber int
    PageSize   int
}

func main() {
    productCatalog := new(sync.Map)

    // Simulate initial loading of data
    loadData(productCatalog)

    // Simulate real-time updates in a goroutine
    go startUpdater(productCatalog)

    // Filter and pagination examples
    filter := Filter{Category: "Cat1", MinPrice: newFloat64(10), MaxPrice: newFloat64(15), InStockOnly: true}
    page := Page{PageNumber: 1, PageSize: 10}
    products := filterAndPaginateProducts(productCatalog, filter, page)

    for _, p := range products {
        fmt.Printf("%+v\n", p)
    }
}

// Helper function to create a pointer to float64
func newFloat64(f float64) *float64 {
    return &f
}

// Simulate real-time updates in a goroutine
func startUpdater(productCatalog *sync.Map) {
    for {
        // Simulate updates arriving at random intervals
        time.Sleep(time.Second * 2)

        update := ProductUpdate{ID: 3, Price: newFloat64(13.99), Stock: newInt(120)}
        if err := applyUpdate(productCatalog, update); err != nil {
            log.Printf("Error applying update: %v", err)
        }
    }
}

func applyUpdate(productCatalog *sync.Map, update ProductUpdate) error {
    var product Product
    if exists := productCatalog.Load(update.ID, &product); !exists {
        return fmt.Errorf("product with ID %d not found", update.ID)
    }

    // Apply the update
    if update.Price != nil {
        product.Price = *update.Price
    }

    if update.Stock != nil {
        product.Stock = *update.Stock
    }

    // Store the updated product back into the sync.Map
    productCatalog.Store(update.ID, product)
    return nil
}

// Helper function to create a pointer to int
func newInt(i int) *int {
    return &i
}

func filterAndPaginateProducts(productCatalog *sync.Map, filter Filter, page Page) []Product {
    // Create a slice to store filtered products
    filteredProducts := []Product{}

    productCatalog.Range(func(key, value interface{}) bool {
        product := value.(Product)

        // Apply filters
        if filter.Category != "" && product.Category != filter.Category {
            return true // continue with next product
        }

        if filter.MinPrice != nil && product.Price < *filter.MinPrice {
            return true
        }

        if filter.MaxPrice != nil && product.Price > *filter.MaxPrice {
            return true
        }

        if filter.InStockOnly && product.Stock == 0 {
            return true
        }

        filteredProducts = append(filteredProducts, product)
        return true
    })

    // Sort filtered products by price
    sort.Slice(filteredProducts, func(i, j int) bool {
        return filteredProducts[i].Price < filteredProducts[j].Price
    })

    // Apply pagination
    startIndex := (page.PageNumber - 1) * page.PageSize
    endIndex := startIndex + page.PageSize
    if endIndex > len(filteredProducts) {
        endIndex = len(filteredProducts)
    }

    return filteredProducts[startIndex:endIndex]
}

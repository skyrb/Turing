package main

import (
    "fmt"
)

type Product struct {
    ID       int
    Name     string
    Category string
    Price    float64
    Stock    int
}

func main() {
    // Example simulated product catalog
    productCatalog := generateLargeProductCatalog()

    // Example filter and pagination criteria
    category := "Electronics"
    priceRange := [2]float64{50.0, 200.0}
    availableStock := true
    page := 1
    pageSize := 10

    filteredProducts := filterProducts(productCatalog, category, priceRange, availableStock)
    paginatedProducts := paginateProducts(filteredProducts, page, pageSize)

    fmt.Println("Filtered and Paginated Products:")
    for _, product := range paginatedProducts {
        fmt.Printf("%+v\n", product)
    }
}

func generateLargeProductCatalog() []Product {
    // Generate a large number of sample products (over hundreds of thousands)
    // In practice, these would be loaded from a database or other data source
    // Here, we'll simulate with a simple example
    var products []Product
    categories := []string{"Electronics", "Books", "Fashion"}
    for i := 0; i < 100000; i++ {
        category := categories[i%len(categories)]
        products = append(products, Product{
            ID:       i,
            Name:     fmt.Sprintf("Product %d", i),
            Category: category,
            Price:    float64(i%100 + 10),
            Stock:    i % 50,
        })
    }
    return products
}

func filterProducts(products []Product, category string, priceRange [2]float64, availableStock bool) []Product {
    var filtered []Product
    for _, product := range products {
        if (category == "" || product.Category == category) &&
            product.Price >= priceRange[0] && product.Price <= priceRange[1] &&
            (!availableStock || product.Stock > 0) {
            filtered = append(filtered, product)
        }
    }
    return filtered
}

func paginateProducts(products []Product, page, pageSize int) []Product {
    start := (page - 1) * pageSize
    if start > len(products) {
        return []Product{}
    }
    end := start + pageSize
    if end > len(products) {
        end = len(products)
    }
    return products[start:end]
}
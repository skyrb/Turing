package main  
import (  
    "fmt"
    "sort"
)

type Product struct {  
    Name     string
    Category string
    Price    float64
    Stock    int
    Rating   float64
}

func main() {  
    // Sample data with tens of thousands of entries
    products := []Product{
        {Name: "Product 1", Category: "Cat1", Price: 10.99, Stock: 100, Rating: 4.7},
        {Name: "Product 2", Category: "Cat2", Price: 5.99, Stock: 50, Rating: 3.8},
        // Add more product entries here
    }

    // Filter and sort products by rating >= 4.5 and price
    filteredProducts := filterAndSortProducts(products)

    // Display the filtered and sorted products
    for _, product := range filteredProducts {  
        fmt.Printf("%s - $%.2f - Rating: %.1f\n", product.Name, product.Price, product.Rating)
    }
}

func filterAndSortProducts(products []Product) []Product {
    // Create a new slice to store filtered products
    filteredProducts := make([]Product, 0, len(products))

    // Filter products with rating >= 4.5
    for _, product := range products {  
        if product.Rating >= 4.5 {
            filteredProducts = append(filteredProducts, product)
        }
    }

    // Sort filtered products by price using the sort package
    sort.Slice(filteredProducts, func(i, j int) bool {
        return filteredProducts[i].Price < filteredProducts[j].Price
    })

    return filteredProducts
}
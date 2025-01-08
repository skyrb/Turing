

package main  
import (  
    "fmt"
    "math/rand"
    "sync"
    "time"
)
type MarketplaceData struct {  
    Name       string
    SalesData []float64
}  
func processSalesData(marketplace MarketplaceData, workerChan chan int, wg *sync.WaitGroup) {  
    defer wg.Done()
    fmt.Printf("Processing data from %s...\n", marketplace.Name)  
    <-workerChan  
    // Simulate some work with random sleep time
    sleepDuration := time.Duration(rand.Intn(3000)) * time.Millisecond
    time.Sleep(sleepDuration)  
    totalRevenue := calculateTotalRevenue(marketplace.SalesData)  
    fmt.Printf("Total revenue from %s: $%.2f\n", marketplace.Name, totalRevenue)  
    workerChan <- 1  
}  
func calculateTotalRevenue(salesData []float64) float64 {  
    var total float64
    for _, sale := range salesData {  
        total += sale
    }  
    return total
}  
func main() {  
    rand.Seed(time.Now().UnixNano())
    // Number of worker goroutines
    numWorkers := 3  
    workerChan := make(chan int, numWorkers)  
    var wg sync.WaitGroup  
    // Sample sales data for each marketplace
    marketplaces := []MarketplaceData{  
        {  
            Name:       "Amazon",
            SalesData: []float64{123.50, 99.99, 49.99, 249.99, 79.99},
        },  
        {  
            Name:       "eBay",
            SalesData: []float64{69.99, 19.99, 89.99, 129.99, 39.99},
        },  
        {  
            Name:       "Shopify",
            SalesData: []float64{149.99, 59.99, 29.99, 79.99, 49.99},
        },  
    }  
    for i := 0; i < numWorkers; i++ {  
        workerChan <- 1  
    }  
    for _, marketplace := range marketplaces {  
        wg.Add(1)  
        go processSalesData(marketplace, workerChan, &wg)  
    }  
    // Close the worker channel to signal that all work is done
    close(workerChan)  
    wg.Wait()  
    fmt.Println("\nProcessing completed. All marketplaces finished.")
    // Calculate and display the total revenue for all marketplaces
    totalRevenue := 0.0
    for _, marketplace := range marketplaces {
        totalRevenue += calculateTotalRevenue(marketplace.SalesData)
    }
    fmt.Printf("Total revenue for the day: $%.2f\n", totalRevenue)
}
  
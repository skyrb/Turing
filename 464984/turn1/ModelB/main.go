package main  
import (  
    "fmt"
    "sync"
)  
  
func validateData(data []int, validValues map[int]bool, wg *sync.WaitGroup) {  
    defer wg.Done()  
    invalidCount := 0  
    for _, value := range data {  
        if _, ok := validValues[value]; !ok {  
            invalidCount++  
        }  
    }  
    fmt.Println("Invalid count:", invalidCount)  
}  
func main() {  
    // Create a map of valid values for efficient lookups
    validValues := make(map[int]bool, 1000)  
    for i := 0; i < 1000; i++ {  
        validValues[i] = true  
    }  
    // Generate large dataset
    data := make([]int, 10000000)  
    for i := 0; i < len(data); i++ {  
        data[i] = i % 1000  
    }  
    const workerCount = 4  
    wg := new(sync.WaitGroup)  
    wg.Add(workerCount)  
    // Partition the data into chunks and process each chunk concurrently
    chunkSize := len(data) / workerCount + 1  
    for i := 0; i < len(data); i += chunkSize {  
        end := i + chunkSize  
        if end > len(data) {  
            end = len(data)  
        }  
        go validateData(data[i:end], validValues, wg)  
    }  
    wg.Wait()  
}  
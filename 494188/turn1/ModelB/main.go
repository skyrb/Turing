package main  
import (  
    "fmt"
    "math/rand"
    "time"
)

// Function to perform a non-blocking I/O operation (simulating random sleep)
func doIO(id int, duration time.Duration, ch chan int) {
    time.Sleep(duration)
    ch <- id // Send the ID back through the channel
}

func main() {
    rand.Seed(time.Now().UnixNano())
    numIOs := 5 // Number of I/O operations to perform
    resultCh := make(chan int, numIOs) // Buffered channel to hold results

    // Start non-blocking I/O operations
    for i := 0; i < numIOs; i++ {
        go doIO(i, time.Duration(rand.Intn(5000)) * time.Millisecond, resultCh)
    }

    // Wait for all I/O operations to complete and collect results
    var results []int
    for i := 0; i < numIOs; i++ {
        result := <-resultCh // Receive result from the channel
        results = append(results, result)
    }

    // Display the results in the order they were completed
    fmt.Println("Results:", results)
}
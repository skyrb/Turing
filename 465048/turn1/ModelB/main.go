package main  
import (  
    "fmt"
    "math/rand"
    "sync"
    "time"
)

// Define a structure to represent the data stream
type DataStream struct {
    Data float64
    Timestamp time.Time
}

// Function to process the data stream
func processDataStream(dataStream <-chan DataStream, wg *sync.WaitGroup) {
    defer wg.Done()
    for data := range dataStream {
        // Simulate processing time
        time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
        fmt.Printf("Processed data: %f at %s\n", data.Data, data.Timestamp.Format("15:04:05.000"))
    }
}

func main() {
    // Number of Goroutines to process data streams
    numGoroutines := 4
    wg := sync.WaitGroup{}
    dataStream := make(chan DataStream)

    // Start the Goroutines
    wg.Add(numGoroutines)
    for i := 0; i < numGoroutines; i++ {
        go processDataStream(dataStream, &wg)
    }

    // Simulate incoming data streams
    for i := 0; i < 100; i++ {
        data := DataStream{Data: rand.Float64(), Timestamp: time.Now()}
        dataStream <- data
    }

    // Close the channel to signal the end of data processing
    close(dataStream)

    // Wait for all Goroutines to finish
    wg.Wait()
    fmt.Println("All data processed.")
}
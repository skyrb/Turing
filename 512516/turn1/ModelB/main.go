package main  
import (  
    "fmt"
    "sync"
)

// UserRecord represents a single user record
type UserRecord struct {  
    UserID  int
    Age     int
    Salary  int
}

// complexCalculation simulates a complex calculation on a user record
func complexCalculation(userRecord UserRecord) UserRecord {
    // Simulate a complex calculation by doing some arithmetic operations
    result := UserRecord{
        UserID:  userRecord.UserID,
        Age:     userRecord.Age * 2,
        Salary:  userRecord.Salary * 3,
    }
    return result
}

// worker function that processes a batch of records
func worker(wg *sync.WaitGroup, workerID int, records <-chan UserRecord, results chan<- UserRecord) {
    defer wg.Done()
    fmt.Printf("Worker %d started\n", workerID)

    for record := range records {
        // Perform the complex calculation on each record
        processedRecord := complexCalculation(record)
        results <- processedRecord
    }

    fmt.Printf("Worker %d finished\n", workerID)
}

func main() {
    // Sample dataset of user records
    userRecords := []UserRecord{
        {UserID: 1, Age: 25, Salary: 5000},
        {UserID: 2, Age: 30, Salary: 7000},
        // Add more user records as needed
    }

    numWorkers := 4 // Number of workers to run concurrently
    batchSize := 10 // Size of each batch of records processed by a worker

    // Create channels for communication between workers and main goroutine
    recordsChan := make(chan UserRecord, batchSize)
    resultsChan := make(chan UserRecord, batchSize)

    var wg sync.WaitGroup

    // Start the specified number of workers
    wg.Add(numWorkers)
    for i := 1; i <= numWorkers; i++ {
        go worker(&wg, i, recordsChan, resultsChan)
    }

    // Distribute the records to the workers
    go func() {
        for _, record := range userRecords {
            recordsChan <- record
        }
        close(recordsChan)
    }()

    // Collect the processed records from the workers
    processedRecords := make([]UserRecord, len(userRecords))
    i := 0
    for record := range resultsChan {
        processedRecords[i] = record
        i++
    }

    // Wait for all workers to finish
    wg.Wait()

    // Display the processed records
    fmt.Println("Processed User Records:")
    for _, record := range processedRecords {
        fmt.Printf("UserID: %d, Age: %d, Salary: %d\n", record.UserID, record.Age, record.Salary)
    }
}
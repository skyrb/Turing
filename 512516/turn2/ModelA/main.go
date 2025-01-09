package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

// UserRecord represents a single user record
type UserRecord struct {
	UserID    int
	Age       int
	Salary    int
	Posts     []string
	Comments  []string
}

// sentimentAnalysis simulates sentiment analysis on posts and comments
func sentimentAnalysis(posts []string, comments []string) (float64, float64) {
	// Simulate sentiment analysis by calculating random scores
	postScore := rand.Float64()
	commentScore := rand.Float64()
	return postScore, commentScore
}

// complexCalculation simulates a complex calculation on a user record
func complexCalculation(ctx context.Context, userRecord UserRecord) (UserRecord, error) {
	// Simulate a complex calculation by doing some arithmetic operations and sentiment analysis
	// Randomly introduce a delay to simulate variable processing time
	delay := time.Duration(rand.Intn(1000)) * time.Millisecond
	time.Sleep(delay)

	// Perform sentiment analysis
	postScore, commentScore, err := sentimentAnalysis(userRecord.Posts, userRecord.Comments)
	if err != nil {
		return UserRecord{}, err
	}

	result := UserRecord{
		UserID:    userRecord.UserID,
		Age:       userRecord.Age * 2,
		Salary:    userRecord.Salary * 3,
		PostScore: postScore,
		CommentScore: commentScore,
	}
	return result, nil
}

// worker function that processes a batch of records
func worker(ctx context.Context, wg *sync.WaitGroup, workerID int, records <-chan UserRecord, results chan<- UserRecord) {
	defer wg.Done()
	fmt.Printf("Worker %d started\n", workerID)

	for record := range records {
		select {
		case <-ctx.Done():
			fmt.Printf("Worker %d shutting down\n", workerID)
			return
		default:
			// Perform the complex calculation on each record
			processedRecord, err := complexCalculation(ctx, record)
			if err != nil {
				fmt.Printf("Worker %d: Error processing user %d: %v\n", workerID, record.UserID, err)
				continue
			}
			results <- processedRecord
		}
	}

	fmt.Printf("Worker %d finished\n", workerID)
}

func main() {
	// Seed the random number generator
	rand.Seed(time.Now().UnixNano())

	// Sample dataset of user records
	userRecords := []UserRecord{
		{UserID: 1, Age: 25, Salary: 5000, Posts: []string{"Great!", "Love it!"}, Comments: []string{"Awesome!", "Best ever!"}},
		{UserID: 2, Age: 30, Salary: 7000, Posts: []string{"I hate this", "This is bad"}, Comments: []string{"Not good", "Terrible!"}},
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
		go worker(context.Background(), &wg, i, recordsChan, resultsChan)
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
		fmt.Printf("UserID: %d, Age: %d, Salary: %d, PostScore: %.2f, CommentScore: %.2f\n", record.UserID, record.Age, record.Salary, record.PostScore, record.CommentScore)
	}
}
package main

import (
	"fmt"
	"strings"
	"sync"
)

func main() {
	// Initialize a large slice of text documents (each document is a string)
	documents := []string{
		"This is the first document.",
		"The second document is about go programming.",
		"This is the third document. Go is a powerful language.",
		"In the fourth document, we discuss text processing.",
		"The fifth document is related to psycholinguistics.",
	}

	// Define the number of Goroutines to use for processing
	numGoroutines := 4

	// Create a channel to receive processed documents
	processedDocsChan := make(chan string)

	// Create a wait group to synchronize the main thread with Goroutines
	var wg sync.WaitGroup
	wg.Add(numGoroutines)

	// Launch Goroutines to process chunks of documents concurrently
	for i := 0; i < numGoroutines; i++ {
		go processDocumentsConcurrently(documents, i, numGoroutines, processedDocsChan, &wg)
	}

	// Close the channel after all Goroutines have finished
	go func() {
		wg.Wait()
		close(processedDocsChan)
	}()

	// Process the processed documents from the channel
	var processedDocuments []string
	for doc := range processedDocsChan {
		processedDocuments = append(processedDocuments, doc)
	}

	// Perform some example text processing and categorization on the processed documents
	categoryCounts := categorizeDocuments(processedDocuments)

	// Display the category counts
	fmt.Println("Category Counts:")
	for cat, count := range categoryCounts {
		fmt.Printf("%s: %d\n", cat, count)
	}
}

// Function to process each chunk of documents concurrently
func processDocumentsConcurrently(documents []string, workerID int, numGoroutines int, processedDocsChan chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()

	// Calculate the start and end indices for this worker's chunk
	startIndex := workerID * len(documents) / numGoroutines
	endIndex := (workerID + 1) * len(documents) / numGoroutines

	// Process each document in this worker's chunk
	for _, doc := range documents[startIndex:endIndex] {
		// Example processing: lowercase, remove punctuation, and split into words
		doc = strings.ToLower(doc)
		doc = strings.ReplaceAll(doc, ".", " ")
		words := strings.Fields(doc)
		processedDoc := strings.Join(words, " ")

		// Send the processed document to the channel
		processedDocsChan <- processedDoc
	}
}

// Function to categorize documents into predefined categories
func categorizeDocuments(documents []string) map[string]int {
	categories := map[string]int{
		"Programming": 0,
		"Psycholinguistics": 0,
		"Other": 0,
	}

	for _, doc := range documents {
		if strings.Contains(doc, "go") {
			categories["Programming"]++
		} else if strings.Contains(doc, "psycholinguistics") {
			categories["Psycholinguistics"]++
		} else {
			categories["Other"]++
		}
	}
	return categories
} 
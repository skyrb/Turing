package main

import (
	"fmt"
	"strings"
	"sync"
)

func main() {
	// Initialize a slice of text documents (each document is a string)
	documents := []string{
		"This is the first document.",
		"The second document is about go programming.",
		"This is the third document. Go is a powerful language.",
		"In the fourth document, we discuss text processing.",
		"The fifth document is related to psycholinguistics.",
		// Add more documents as needed
	}

	// Define the number of worker Goroutines
	numWorkers := 4
	// Initialize the sync group
	var wg sync.WaitGroup
	// Create a channel to receive processed documents
	processedDocumentsChan := make(chan string)

	// Start worker Goroutines
	for w := 0; w < numWorkers; w++ {
		wg.Add(1)
		go worker(w, documents, processedDocumentsChan)
	}

	// Send documents to worker Goroutines
	for _, doc := range documents {
		processedDocumentsChan <- doc
	}

	// Close the channel to signal workers that no more documents are coming
	close(processedDocumentsChan)

	// Wait for all workers to finish
	wg.Wait()

	// Gather and display processed documents
	var processedDocuments []string
	for doc := range processedDocumentsChan {
		processedDocuments = append(processedDocuments, doc)
	}

	fmt.Println("Processed Documents:")
	for _, doc := range processedDocuments {
		fmt.Println(doc)
	}

	// Categorize processed documents
	categoryCounts := categorizeDocuments(processedDocuments)

	// Display category counts
	fmt.Println("\nCategory Counts:")
	for cat, count := range categoryCounts {
		fmt.Printf("%s: %d\n", cat, count)
	}
}

// Worker function: processes each document and sends the result back to the main goroutine
func worker(id int, documents []string, processedDocumentsChan chan string) {
	defer wg.Done()
	for doc := range processedDocumentsChan {
		// Example processing: lowercase, remove punctuation, and split into words
		doc = strings.ToLower(doc)
		doc = strings.ReplaceAll(doc, ".", " ")
		words := strings.Fields(doc)
		processedDocumentsChan <- strings.Join(words, " ") // Send back the processed document
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
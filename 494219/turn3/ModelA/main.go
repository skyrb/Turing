package main

import (
	"fmt"
	"math/rand"
	"strings"
	"sync"
	"time"
)

func main() {
	// Initialize a large slice of text documents
	documents := generateLargeDataset()

	// Set the number of Goroutines
	numGoroutines := 4

	// Create a channel to send chunks of documents
	docChunkChan := make(chan []string, numGoroutines)

	// Create a channel to receive dynamically determined categories
	categoryChan := make(chan string)

	// Create a wait group for synchronization
	var wg sync.WaitGroup
	wg.Add(numGoroutines)

	// Launch processing Goroutines
	for i := 0; i < numGoroutines; i++ {
		go processDocumentsConcurrently(docChunkChan, categoryChan, &wg)
	}

	// Send chunks of documents to Goroutines
	chunkSize := len(documents) / numGoroutines
	for i := 0; i < numGoroutines; i++ {
		startIndex := i * chunkSize
		endIndex := (i + 1) * chunkSize
		if i == numGoroutines-1 {
			endIndex = len(documents)
		}
		docChunkChan <- documents[startIndex:endIndex]
	}

	// Close the document chunk channel
	close(docChunkChan)

	// Receive dynamically determined categories
	categories := make(map[string]int)
	for category := range categoryChan {
		categories[category]++
	}

	// Close the category channel
	close(categoryChan)

	// Wait for all Goroutines to finish
	wg.Wait()

	// Display the dynamic category counts
	fmt.Println("Dynamic Category Counts:")
	for cat, count := range categories {
		fmt.Printf("%s: %d\n", cat, count)
	}
}

// Function to generate a large dataset of documents
func generateLargeDataset() []string {
	// Generate a large dataset of random documents for demonstration
	n := 10000 // Number of documents
	dataset := make([]string, n)
	for i := 0; i < n; i++ {
		// Randomly generate a document
		doc := strings.Join(generateWords(20), " ")
		dataset[i] = doc
	}
	return dataset
}

// Function to generate random words for demonstration
func generateWords(numWords int) []string {
	words := []string{"go", "lang", "programming", "learn", "code", "psycholinguistics", "text", "processing", "concurrency", "channels"}
	rand.Seed(time.Now().UnixNano())
	return rand.Perm(numWords)[:]
}

// Function to process each chunk of documents concurrently
func processDocumentsConcurrently(docChunkChan <-chan []string, categoryChan chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()

	for docChunk := range docChunkChan {
		// Process each document in the chunk
		for _, doc := range docChunk {
			// Example processing: lowercase and replace punctuation
			doc = strings.ToLower(doc)
			doc = strings.ReplaceAll(doc, ".", " ")
			doc = strings.ReplaceAll(doc, ",", " ")

			// Dynamically determine categories based on keywords
			keywords := []string{"go", "lang", "psycholinguistics", "processing"}
			categories := make(map[string]int)
			for _, keyword := range keywords {
				if strings.Contains(doc, keyword) {
					categories[keyword]++
				}
			}

			// Determine the dominant category for this document
			if len(categories) > 0 {
				maxCount := categories[keywords[0]]
				dominantCategory := keywords[0]
				for keyword, count := range categories {
					if count > maxCount {
						maxCount = count
						dominantCategory = keyword
					}
				}
				categoryChan <- dominantCategory
			} else {
				categoryChan <- "Other"
			}
		}
	}
}
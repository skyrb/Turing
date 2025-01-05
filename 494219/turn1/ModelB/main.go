package main

import (
	"fmt"
	"strings"
)

func main() {
	// Initialize a slice of text documents (each document is a string)
	documents := []string{
		"This is the first document.",
		"The second document is about go programming.",
		"This is the third document. Go is a powerful language.",
		"In the fourth document, we discuss text processing.",
		"The fifth document is related to psycholinguistics.",
	}

	// Perform some example text processing and categorization using slices
	processedDocuments := processDocuments(documents)
	categoryCounts := categorizeDocuments(processedDocuments)

	// Display the processed documents and category counts
	fmt.Println("Processed Documents:")
	for _, doc := range processedDocuments {
		fmt.Println(doc)
	}

	fmt.Println("\nCategory Counts:")
	for cat, count := range categoryCounts {
		fmt.Printf("%s: %d\n", cat, count)
	}
}

// Function to process each document in the slice
func processDocuments(documents []string) []string {
	var processedDocuments []string
	for _, doc := range documents {
		// Example processing: lowercase, remove punctuation, and split into words
		doc = strings.ToLower(doc)
		doc = strings.ReplaceAll(doc, ".", " ")
		words := strings.Fields(doc)
		processedDocuments = append(processedDocuments, strings.Join(words, " "))
	}
	return processedDocuments
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
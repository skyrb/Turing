package main

import (
	"fmt"
	"strings"
)

// CategorizeText categorizes a text into one of three psycholinguistic categories.
func categorizeText(text string) string {
	// Simple heuristics for categorization (replace these with a more sophisticated model)
	switch strings.ToLower(text) {
	case "happy", "good", "great":
		return "positive"
	case "sad", "bad", "terrible":
		return "negative"
	default:
		return "neutral"
	}
}

func main() {
	// Initialize an empty slice to hold the text data
	textData := []string{}

	// Example: Read text data from a file or another source
	// For demonstration, we'll add some example texts manually
	textData = append(textData, "I am feeling happy today!")
	textData = append(textData, "This is a terrible day.")
	textData = append(textData, "The weather is okay.")

	// Print the original text data
	fmt.Println("Original text data:")
	for _, text := range textData {
		fmt.Println(text)
	}

	// Create a slice to store the categorized results
	categorizedTexts := make([]string, len(textData))

	// Categorize each text in the dataset
	for i, text := range textData {
		categorizedTexts[i] = categorizeText(text)
	}

	// Print the categorized text data
	fmt.Println("\nCategorized text data:")
	for _, category := range categorizedTexts {
		fmt.Println(category)
	}
}
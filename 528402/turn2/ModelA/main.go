package main

import (
	"fmt"
	"strings"
)

func main() {
	// Store sentences in a Go slice
	sentences := []string{
		"The sky is blue.",
		"Natural language processing is fascinating.",
		"Golang is efficient for such tasks.",
	}

	// Create a map to group words by their length
	wordsByLength := make(map[int][]string)

	// Process the sentences
	for _, sentence := range sentences {
		// Tokenize the sentence into words
		words := strings.Fields(sentence)

		// Group words by their length
		for _, word := range words {
			// Remove punctuations, for cleaner grouping (if necessary)
			cleanedWord := strings.Trim(word, ".,")

			// Get the length of the word
			wordLength := len(cleanedWord)

			// Add the word to the corresponding length group
			wordsByLength[wordLength] = append(wordsByLength[wordLength], cleanedWord)
		}
	}

	// Display the grouped words by their lengths
	fmt.Println("Words grouped by their lengths:")
	for length, words := range wordsByLength {
		fmt.Printf("Length %d: %v\n", length, words)
	}
}
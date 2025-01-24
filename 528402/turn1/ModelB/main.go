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

	// Process the sentences
	for _, sentence := range sentences {
		words := strings.Fields(sentence)
		fmt.Println("Sentence:", sentence)
		fmt.Println("Words:", words)
		fmt.Println("Number of words:", len(words))
		fmt.Println("--------------------")
	}

	// Append a new sentence to the slice
	newSentence := "Go is a powerful programming language."
	sentences = append(sentences, newSentence)
	fmt.Println("\nAppended new sentence:", newSentence)
	fmt.Println("Total sentences:", len(sentences))
}
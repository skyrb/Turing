package main

import (
	"fmt"
	"strings"
)

func contains(s, substr string) bool {
	return strings.Contains(s, substr)
}

func main() {
	txtCorpus := []string{
		"The sky is blue.",
		"Natural language processing is fascinating.",
		"Golang is efficient for such tasks.",
	}
	newSentences := []string{
		"Data science is growing rapidly.",
		"Concurrency in Go is powerful.",
	}
	
	// Append new sentences to the existing slice
	txtCorpus = append(txtCorpus, newSentences...)
	
	// Print the slice to verify contents
	fmt.Println(txtCorpus)

	// Access specific sentence by index
	firstSentence := txtCorpus[0]
	fmt.Println("First Sentence:", firstSentence)

	// Iterate over all sentences
	fmt.Println("All Sentences:")
	for _, sentence := range txtCorpus {
		fmt.Println(sentence)
	}

	keyword := "Golang"
	for _, sentence := range txtCorpus {
		if contains(sentence, keyword) {
			fmt.Println("Contains keyword:", sentence)
		}
	}

	fmt.Println("Sentences in Uppercase:")
	for _, sentence := range txtCorpus {
		fmt.Println(strings.ToUpper(sentence))
}
}
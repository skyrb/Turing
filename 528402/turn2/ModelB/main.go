package main  
import (  
    "fmt"
    "strings"
)  
func main() {  
    sentences := []string{  
        "The sky is blue.",  
        "Natural language processing is fascinating.",  
        "Golang is efficient for such tasks.",  
    }  
    wordGroups := make(map[int][]string)  
    for _, sentence := range sentences {  
        words := strings.Fields(sentence)  
        fmt.Println("Sentence:", sentence)  
        fmt.Println("Words:", words)  
        fmt.Println("Number of words:", len(words))  
        fmt.Println("--------------------")
        // Group words by their lengths
        for _, word := range words {
            wordLength := len(word)
            wordGroups[wordLength] = append(wordGroups[wordLength], word)
        }
    }
    // Display the word groups
    fmt.Println("\nWord Groups:")
    for length, words := range wordGroups {
        fmt.Printf("Length %d: %v\n", length, words)
    }
}
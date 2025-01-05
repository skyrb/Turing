package main
import (
	"fmt"
	"strings"
	"sync"
)

type Categorizer struct {
	mutex sync.Mutex
	categories map[string]int
}

func NewCategorizer() *Categorizer {
	return &Categorizer{
		categories: make(map[string]int),
	}
}

func (c *Categorizer) LearnCategory(category string) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	if _, exists := c.categories[category]; !exists {
		c.categories[category] = 0
	}
}

func (c *Categorizer) IncrementCategory(category string) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.categories[category]++
}

func (c *Categorizer) GetCategories() map[string]int {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	return c.categories
}

func main() {
	documents := []string{
		"This is the first document about programming.",
		"Go is a fantastic language for web development.",
		"Psycholinguistics research is exciting.",
		"This document is about AI and machine learning.",
		"The fifth document is related to mental health and psychology.",
		"Apple is a tech company. Let's discuss their latest products.",
		"Shopping is great. Need to buy some groceries.",
		"Weather forecasting is important.",
	}

	categorizer := NewCategorizer()

	numGoroutines := 4
	processedDocsChan := make(chan string)
	var wg sync.WaitGroup
	wg.Add(numGoroutines)

	for i := 0; i < numGoroutines; i++ {
		go processDocumentsConcurrently(documents, i, numGoroutines, processedDocsChan, &wg, categorizer)
	}

	go func() {
		wg.Wait()
		close(processedDocsChan)
	}()

	var processedDocuments []string
	for doc := range processedDocsChan {
		processedDocuments = append(processedDocuments, doc)
	}

	categoryCounts := categorizer.GetCategories()

	fmt.Println("Category Counts:")
	for cat, count := range categoryCounts {
		fmt.Printf("%s: %d\n", cat, count)
	}
}

func processDocumentsConcurrently(documents []string, workerID int, numGoroutines int, processedDocsChan chan<- string, wg *sync.WaitGroup, categorizer *Categorizer) {
	defer wg.Done()

	startIndex := workerID * len(documents) / numGoroutines
	endIndex := (workerID + 1) * len(documents) / numGoroutines

	for _, doc := range documents[startIndex:endIndex] {
		doc = strings.ToLower(doc)
		doc = strings.ReplaceAll(doc, ".", " ")
		words := strings.Fields(doc)

		var categories []string
		// Now, use NLP techniques like string matching, tokenization, or topic modeling to determine categories
		for _, word := range words {
			if strings.Contains(doc, word) {
				categories = append(categories, word)
			}
		}

		//Learn new categories from the input
		for _, cat := range categories {
			categorizer.LearnCategory(cat)
			categorizer.IncrementCategory(cat)
		}
		processedDocsChan <- strings.Join(categories, " ")
	}
} 
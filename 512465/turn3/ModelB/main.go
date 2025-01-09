package main

import (
	"fmt"
	"strings"
)

type Message struct {
	User    string
	Content string
}

type Trie struct {
	root *trieNode
}

type trieNode struct {
	children map[rune]*trieNode
	isWord   bool
	message  Message
}

func (t *Trie) Insert(message Message) {
	current := t.root
	for _, char := range strings.ToLower(message.Content + " " + message.User) {
		if current.children[char] == nil {
			current.children[char] = &trieNode{children: make(map[rune]*trieNode)}
		}
		current = current.children[char]
	}
	current.isWord = true
	current.message = message
}

func (t *Trie) Search(keywords []string) []Message {
	results := []Message{}
	keywordSet := make(map[string]struct{}, len(keywords))
	for _, keyword := range keywords {
		keywordSet[strings.ToLower(keyword)] = struct{}{}
	}

	t.searchRecursive(t.root, "", keywordSet, &results)
	return results
}

func (t *Trie) searchRecursive(current *trieNode, word string, keywordSet map[string]struct{}, results *[]Message) {
	if current == nil {
		return
	}

	if current.isWord && len(keywordSet) == 0 {
		*results = append(*results, current.message)
		return
	}

	if _, ok := keywordSet[word]; ok {
		delete(keywordSet, word)
	}

	for char, next := range current.children {
		t.searchRecursive(next, word+string(char), keywordSet, results)
	}
}

func main() {
	app := MessagingApp{}

	// Add some sample messages
	app.AddMessage("Alice", "Hi Bob! Let's meet at 3pm for lunch.")
	app.AddMessage("Bob", "Okay, Alice! See you then.")
	app.AddMessage("Charlie", "Just checking in. How are you?")
	app.AddMessage("David", "Hi Alice, can we have a meeting at 2pm?")
	app.AddMessage("Eve", "I'm free for lunch at 3pm today.")
	app.AddMessage("Frank", "Hi everyone! Want to join for coffee at 4pm?")
    
    // Build the Trie data structure
    trie := Trie{root: &trieNode{children: make(map[rune]*trieNode)}}
    for _, msg := range app.Messages {
        trie.Insert(msg)
    }

    // Search for messages containing "Alice", "meeting", and "lunch"
    keywords := []string{"Alice", "meeting", "lunch"}
    results := trie.Search(keywords)

    fmt.Printf("Messages containing '%s': \n", strings.Join(keywords, " and "))
    for _, result := range results {
        fmt.Printf("%s: %s\n", result.User, result.Content)
    }
} 
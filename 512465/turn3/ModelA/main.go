package main

import (
	"fmt"
	"strings"
)

// TrieNode represents each node in the Trie
type TrieNode struct {
	children map[rune]*TrieNode
	isEnd    bool
}

// Trie represents the entire Trie structure
type Trie struct {
	root *TrieNode
}

// NewTrie initializes a new Trie
func NewTrie() *Trie {
	return &Trie{root: &TrieNode{children: make(map[rune]*TrieNode)}}
}

// Insert inserts a word into the Trie
func (t *Trie) Insert(word string) {
	node := t.root
	for _, char := range word {
		if child, ok := node.children[char]; !ok {
			node.children[char] = &TrieNode{children: make(map[rune]*TrieNode)}
		}
		node = node.children[char]
	}
	node.isEnd = true
}

// Search searches for a word in the Trie
func (t *Trie) Search(word string) bool {
	node := t.root
	for _, char := range word {
		if child, ok := node.children[char]; !ok {
			return false
		}
		node = child
	}
	return node.isEnd
}

type Message struct {
	User    string
	Content string
}

type MessagingApp struct {
	Messages []Message
	trie     *Trie
}

func (app *MessagingApp) AddMessage(user, content string) {
	app.Messages = append(app.Messages, Message{User: user, Content: content})
	// Split content into words and insert each into the Trie
	words := strings.Fields(content)
	for _, word := range words {
		app.trie.Insert(word)
	}
}

func (app *MessagingApp) SearchMessages(keywords ...string) []Message {
	var results []Message
	trie := NewTrie()

	// Insert all keywords into a separate Trie
	for _, keyword := range keywords {
		trie.Insert(keyword)
	}

	for _, message := range app.Messages {
		found := true
		// Check if all keywords are present in the message
		for _, word := range strings.Fields(message.Content) {
			if !trie.Search(word) {
				found = false
				break
			}
		}
		if found {
			results = append(results, message)
		}
	}
	return results
}

func main() {
	app := MessagingApp{trie: NewTrie()}

	// Add some sample messages
	app.AddMessage("Alice", "Hi Bob! Let's meet at 3pm for lunch.")
	app.AddMessage("Bob", "Okay, Alice! See you then.")
	app.AddMessage("Charlie", "Just checking in. How are you?")
	app.AddMessage("David", "Hi Alice, can we have a meeting at 2pm?")
	app.AddMessage("Alice", "Sure, David! That sounds good.")

	// Search for messages containing both "Alice" and "meeting"
	keywords := []string{"Alice", "meeting"}
	results := app.SearchMessages(keywords...)

	fmt.Printf("Messages containing '%s': \n", strings.Join(keywords, " and "))
	for _, result := range results {
		fmt.Printf("%s: %s\n", result.User, result.Content)
	}
}
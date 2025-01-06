package main

import (
	"fmt"
	"log"
	"net/http"
	"encoding/json"
)

// Post represents a single content post
type Post struct {
	ID    string  `json:"id"`
	Title string  `json:"title"`
	Body  string  `json:"body"`
}

// ContentManager manages a collection of posts using a map
var ContentManager map[string]Post = make(map[string]Post)

// AddPost adds a new post to the content manager
func AddPost(post Post) {
	ContentManager[post.ID] = post
}

// UpdatePost updates an existing post in the content manager
func UpdatePost(post Post) {
	ContentManager[post.ID] = post
}

// GetPost retrieves a post by its ID from the content manager
func GetPost(id string) (*Post, error) {
	post, exists := ContentManager[id]
	if !exists {
		return nil, fmt.Errorf("post not found: %s", id)
	}
	return &post, nil
}

// Main server handler
func main() {
	http.HandleFunc("/add-post", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
			return
		}

		var post Post
		if err := json.NewDecoder(r.Body).Decode(&post); err != nil {
			http.Error(w, "Error decoding JSON", http.StatusBadRequest)
			log.Printf("Error decoding JSON: %v", err)
			return
		}

		AddPost(post)
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(post)
	})

	http.HandleFunc("/update-post", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
			return
		}

		var post Post
		if err := json.NewDecoder(r.Body).Decode(&post); err != nil {
			http.Error(w, "Error decoding JSON", http.StatusBadRequest)
			log.Printf("Error decoding JSON: %v", err)
			return
		}

		UpdatePost(post)
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(post)
	})

	http.HandleFunc("/get-post/{id}", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
			return
		}

		id := r.URL.Query().Get("id")
		if id == "" {
			http.Error(w, "ID is required", http.StatusBadRequest)
			return
		}

		post, err := GetPost(id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			log.Printf("Error retrieving post: %v", err)
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(post)
	})

	log.Println("Starting CMS server on port :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
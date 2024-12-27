package main

import (
	"fmt"
	"time"
)

// Define a Post struct
type Post struct {
	ID        int
	Title     string
	Content   string
	Timestamp time.Time
	Comments  []Comment
}

// Define a Comment struct
type Comment struct {
	ID        int
	Author    string
	Content   string
	Timestamp time.Time
}

// Create a PostManager struct to hold a list of posts
type PostManager struct {
	Posts []Post
}

// Add a new post to the PostManager
func (pm *PostManager) AddPost(post Post) {
	post.ID = len(pm.Posts) + 1 // Assign a unique ID based on the number of posts
	pm.Posts = append(pm.Posts, post)
}

// Update an existing post by its ID
func (pm *PostManager) UpdatePost(id int, updatedPost Post) {
	for i, post := range pm.Posts {
		if post.ID == id {
			pm.Posts[i] = updatedPost
			return
		}
	}
	fmt.Printf("Post with ID %d not found\n", id)
}

// Retrieve a post by its ID
func (pm *PostManager) GetPost(id int) Post {
	for _, post := range pm.Posts {
		if post.ID == id {
			return post
		}
	}
	return Post{} // Return an empty Post if not found
}

// Add a new comment to a specific post by its ID
func (pm *PostManager) AddComment(postID int, comment Comment) {
	for i, post := range pm.Posts {
		if post.ID == postID {
			comment.ID = len(post.Comments) + 1 // Assign a unique ID based on the number of comments
			post.Comments = append(post.Comments, comment)
			pm.Posts[i] = post
			return
		}
	}
	fmt.Printf("Post with ID %d not found\n", postID)
}

// List all posts
func (pm *PostManager) ListPosts() {
	for _, post := range pm.Posts {
		fmt.Printf("Post ID: %d, Title: %s, Timestamp: %s\n", post.ID, post.Title, post.Timestamp.Format(time.RFC822))
	}
}

// Main function to demonstrate the functionality
func main() {
	// Initialize a PostManager
	pm := PostManager{}

	// Add a new post
	pm.AddPost(Post{
		Title:     "Welcome to My Blog",
		Content:   "This is the first post on my blog.",
		Timestamp: time.Now(),
	})

	// Add a new comment to the first post
	pm.AddComment(1, Comment{
		Author:    "Alice",
		Content:   "Great first post!",
		Timestamp: time.Now(),
	})

	// List all posts
	pm.ListPosts()

	// Update an existing post
	pm.UpdatePost(1, Post{
		Title:     "Welcome to My Blog (Updated)",
		Content:   "This is the first post on my blog, now updated.",
		Timestamp: time.Now(),
	})

	// Retrieve and print a specific post
	post := pm.GetPost(1)
	fmt.Printf("\nRetrieved Post:\nTitle: %s\nContent: %s\nTimestamp: %s\n", post.Title, post.Content, post.Timestamp.Format(time.RFC822))
	fmt.Println("Comments:")
	for _, comment := range post.Comments {
		fmt.Printf("Comment ID: %d, Author: %s, Content: %s, Timestamp: %s\n", comment.ID, comment.Author, comment.Content, comment.Timestamp.Format(time.RFC822))
	}
}
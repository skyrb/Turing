package main

import (
	"fmt"
	"log"
	"net/http"
	"runtime/pprof"
)

// Define a Post struct
type Post struct {
	ID       int
	Title    string
	Content  string
	Comments []Comment
}

// Define a Comment struct
type Comment struct {
	ID     int
	PostID int
	Text   string
}

func main() {
	// Start CPU profiling
	f, err := os.Create("cpu.prof")
	if err != nil {
		log.Fatal("could not create CPU profile:", err)
	}
	pprof.StartCPUProfile(f)
	defer pprof.StopCPUProfile()

	// Start memory profiling
	memoryProfile := make(chan *pprof.Profile, 1)
	go func() {
		memoryProfile <- pprof.Lookup("memory")
	}()
	defer memoryProfile <- nil

	// Initialize slices to store posts and comments
	posts := []Post{}
	comments := []Comment{}

	// Add a new post
	newPost := Post{ID: 1, Title: "Hello, World!", Content: "This is my first post."}
	posts = append(posts, newPost)

	// Add a new comment for the first post
	newComment := Comment{ID: 1, PostID: 1, Text: "Great post!"}
	comments = append(comments, newComment)

	// Update the content of the first post
	posts[0].Content = "This is my updated first post."

	// Retrieve and display a post by its ID
	postIDToRetrieve := 1
	for _, post := range posts {
		if post.ID == postIDToRetrieve {
			fmt.Println("Post:")
			fmt.Println("ID:", post.ID)
			fmt.Println("Title:", post.Title)
			fmt.Println("Content:", post.Content)
			fmt.Println("Comments:")
			for _, comment := range post.Comments {
				fmt.Println("\tID:", comment.ID)
				fmt.Println("\tText:", comment.Text)
			}
			break
		}
	}

	// Retrieve and display comments for a post by its ID
	postIDToRetrieveComments := 1
	fmt.Println("Comments for Post ID", postIDToRetrieveComments, ":")
	for _, comment := range comments {
		if comment.PostID == postIDToRetrieveComments {
			fmt.Println("\tID:", comment.ID)
			fmt.Println("\tText:", comment.Text)
		}
	}
}
package main

import (
	"fmt"
)

type Post struct {
	ID     int
	Title  string
	Content string
}

var posts []Post

func main() {
	menu()
}

func menu() {
	fmt.Println("Blogging Platform")
	fmt.Println("1. Add Post")
	fmt.Println("2. Edit Post")
	fmt.Println("3. Delete Post")
	fmt.Println("4. List Posts")
	fmt.Println("5. Exit")

	var choice int
	fmt.Scanf("%d", &choice)

	switch choice {
	case 1:
		addPost()
	case 2:
		editPost()
	case 3:
		deletePost()
	case 4:
		listPosts()
	case 5:
		fmt.Println("Exiting...")
	default:
		fmt.Println("Invalid choice. Please try again.")
		menu()
	}
}

func addPost() {
	var title, content string
	fmt.Println("Enter the title:")
	fmt.Scanln(&title)
	fmt.Println("Enter the content:")
	fmt.Scanln(&content)
	newPost := Post{ID: len(posts) + 1, Title: title, Content: content}
	posts = append(posts, newPost)
	fmt.Println("Post added successfully!")
	menu()
}

func editPost() {
	var id int
	fmt.Println("Enter the ID of the post to edit:")
	fmt.Scanf("%d", &id)
	id-- // Convert from 1-based to 0-based index
	if id >= 0 && id < len(posts) {
		var title, content string
		fmt.Println("Enter the new title:")
		fmt.Scanln(&title)
		fmt.Println("Enter the new content:")
		fmt.Scanln(&content)
		posts[id].Title = title
		posts[id].Content = content
		fmt.Println("Post edited successfully!")
	} else {
		fmt.Println("Post not found.")
	}
	menu()
}

func deletePost() {
	var id int
	fmt.Println("Enter the ID of the post to delete:")
	fmt.Scanf("%d", &id)
	id-- // Convert from 1-based to 0-based index
	if id >= 0 && id < len(posts) {
		fmt.Println("Are you sure you want to delete this post? (y/n)")
		var confirm string
		fmt.Scanln(&confirm)
		if confirm == "y" || confirm == "Y" {
			posts = append(posts[:id], posts[id+1:]...)
			fmt.Println("Post deleted successfully!")
		} else {
			fmt.Println("Deletion cancelled.")
		}
	} else {
		fmt.Println("Post not found.")
	}
	menu()
}

func listPosts() {
	if len(posts) == 0 {
		fmt.Println("No posts found.")
		return
	}

	fmt.Println("List of Posts:")
	for _, post := range posts {
		fmt.Printf("ID: %d, Title: %s, Content: %s\n", post.ID, post.Title, post.Content)
	}
	menu()
}
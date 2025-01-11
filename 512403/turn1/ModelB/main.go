
package main  
import (  
    "fmt"
    "log"
    "net/http"
    "strconv"
)

type Post struct {  
    ID      int
    Title   string
    Content string
}

var posts []Post

// Handler functions
func indexHandler(w http.ResponseWriter, r *http.Request) {  
    fmt.Fprintf(w, "Blogging Platform\n")
    for _, post := range posts {  
        fmt.Fprintf(w, "%d. %s\n", post.ID, post.Title)
    }
}

func newPostHandler(w http.ResponseWriter, r *http.Request) {  
    if r.Method != "POST" {  
        fmt.Fprintf(w, "<form method=\"post\">\n")
        fmt.Fprintf(w, "<label for=\"title\">Title:</label>\n")
        fmt.Fprintf(w, "<input type=\"text\" id=\"title\" name=\"title\"><br>\n")
        fmt.Fprintf(w, "<label for=\"content\">Content:</label><br>\n")
        fmt.Fprintf(w, "<textarea id=\"content\" name=\"content\" rows=\"4\" cols=\"50\"></textarea><br>\n")
        fmt.Fprintf(w, "<input type=\"submit\" value=\"Submit\">\n")
        fmt.Fprintf(w, "</form>")
        return
    }
    // Handle POST request
    title := r.FormValue("title")
    content := r.FormValue("content")
    // Validate input
    if title == "" || content == "" {
        fmt.Fprintf(w, "Error: Title and content are required.")
        return
    }
    // Create new post
    newPost := Post{
        ID:      len(posts) + 1,
        Title:   title,
        Content: content,
    }
    posts = append(posts, newPost)
    fmt.Fprintf(w, "Post added successfully! <a href=\"/index\">Back to index</a>")
}

func editPostHandler(w http.ResponseWriter, r *http.Request) {  
    postID, err := strconv.Atoi(r.FormValue("id"))
    if err != nil {
        fmt.Fprintf(w, "Error: Invalid post ID.")
        return
    }
    for index, post := range posts {  
        if post.ID == postID {  
            if r.Method != "POST" {  
                // Show edit form
                fmt.Fprintf(w, "<form method=\"post\">\n")
                fmt.Fprintf(w, "<label for=\"title\">Title:</label>\n")
                fmt.Fprintf(w, "<input type=\"text\" id=\"title\" name=\"title\" value=\"%s\"><br>\n", post.Title)
                fmt.Fprintf(w, "<label for=\"content\">Content:</label><br>\n")
                fmt.Fprintf(w, "<textarea id=\"content\" name=\"content\" rows=\"4\" cols=\"50\">%s</textarea><br>\n", post.Content)
                fmt.Fprintf(w, "<input type=\"submit\" value=\"Submit\">\n")
                fmt.Fprintf(w, "</form>")
                return
            }
            // Handle POST request
            title := r.FormValue("title")
            content := r.FormValue("content")
            // Validate input
            if title == "" || content == "" {
                fmt.Fprintf(w, "Error: Title and content are required.")
                return
            }
            // Update post
            posts[index].Title = title
            posts[index].Content = content
            fmt.Fprintf(w, "Post updated successfully! <a href=\"/index\">Back to index</a>")
            return
        }
    }
    fmt.Fprintf(w, "Error: Post not found.")
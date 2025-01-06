package main  
import (  
    "fmt"
    "sync"
)
// User represents a user in the database
type User struct {  
    ID   int    `json:"id"`
    Name string `json:"name"`
    Email string `json:"email"`
}
// In-memory database using maps
var users map[int]User
var m sync.Mutex

func init() {
    users = make(map[int]User)
}

func createUser(user User) {
    m.Lock()
    defer m.Unlock()
    users[user.ID] = user
}

func getUser(id int) (User, bool) {
    m.Lock()
    defer m.Unlock()
    user, ok := users[id]
    return user, ok
}

func updateUser(user User) {
    m.Lock()
    defer m.Unlock()
    users[user.ID] = user
}

func deleteUser(id int) {  
    m.Lock()
    defer m.Unlock()
    delete(users, id)
}

func main() {
    // Create some users
    createUser(User{ID: 1, Name: "Alice", Email: "alice@example.com"})
    createUser(User{ID: 2, Name: "Bob", Email: "bob@example.com"})
    fmt.Println("Users:")
    for _, user := range users {
        fmt.Printf("ID: %d, Name: %s, Email: %s\n", user.ID, user.Name, user.Email)
    }
    // Get a user
    user, found := getUser(1)
    if found {
        fmt.Printf("Found User: ID: %d, Name: %s, Email: %s\n", user.ID, user.Name, user.Email)
    } else {
        fmt.Println("User not found.")
    }
    // Update a user
    updatedUser := User{ID: 2, Name: "Bobby", Email: "bobby@example.com"}
    updateUser(updatedUser)
    fmt.Println("Updated Users:")
    for _, user := range users {
        fmt.Printf("ID: %d, Name: %s, Email: %s\n", user.ID, user.Name, user.Email)
    }
    // Delete a user
    deleteUser(1)
    fmt.Println("Users after deletion:")
    for _, user := range users {  
        fmt.Printf("ID: %d, Name: %s, Email: %s\n", user.ID, user.Name, user.Email)
    }
}
package main  
import (  
    "encoding/json"
    "fmt"
    "log"
    "net/http"
    "strconv"
    "sync"
)

type User struct {  
    ID   int    `json:"id"`
    Name string `json:"name"`
    Email string `json:"email"`
}

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

func handleCreateUser(w http.ResponseWriter, r *http.Request) {  
    if r.Method != http.MethodPost {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }
    var user User
    if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
    createUser(user)
    w.WriteHeader(http.StatusCreated)
}

func handleGetUser(w http.ResponseWriter, r *http.Request) {  
    if r.Method != http.MethodGet {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }
    idStr := r.URL.Query().Get("id")
    id, err := strconv.Atoi(idStr)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
    user, found := getUser(id)
    if !found {
        http.Error(w, "User not found", http.StatusNotFound)
        return
    }
    json.NewEncoder(w).Encode(user)
}

func handleUpdateUser(w http.ResponseWriter, r *http.Request) {  
    if r.Method != http.MethodPut {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }
    var user User
    if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
    updateUser(user)
}

func handleDeleteUser(w http.ResponseWriter, r *http.Request) {  
    if r.Method != http.MethodDelete {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }
    idStr := r.URL.Query().Get("id")
    id, err := strconv.Atoi(idStr)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
    deleteUser(id)
}

func main() {  
    http.HandleFunc("/users", handleCreateUser)
    http.HandleFunc("/users/", handleGetUser)
    http.HandleFunc("/users/update", handleUpdateUser)
    http.HandleFunc("/users/delete", handleDeleteUser)
    fmt.Println("Server is running on port 8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
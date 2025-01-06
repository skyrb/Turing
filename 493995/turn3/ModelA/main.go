package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"sync"
	"sync/atomic"

	"github.com/gorilla/mux"
)

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

var users map[int]User
var userCount uint64
var mutex sync.Mutex

func init() {
	users = make(map[int]User)
}

// Create a new user
func handleCreateUser(w http.ResponseWriter, r *http.Request) {
	var user User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Error decoding JSON request body", http.StatusBadRequest)
		return
	}

	mutex.Lock()
	defer mutex.Unlock()

	users[user.ID] = user
	atomic.AddUint64(&userCount, 1)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}

// Get an existing user
func handleGetUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Error parsing user ID", http.StatusBadRequest)
		return
	}

	user, found := users[id]

	if !found {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(user)
}

// Update an existing user
func handleUpdateUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Error parsing user ID", http.StatusBadRequest)
		return
	}

	var user User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Error decoding JSON request body", http.StatusBadRequest)
		return
	}

	mutex.Lock()
	defer mutex.Unlock()

	users[id] = user

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}

// Delete an existing user
func handleDeleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Error parsing user ID", http.StatusBadRequest)
		return
	}

	mutex.Lock()
	defer mutex.Unlock()

	delete(users, id)
	atomic.AddUint64(&userCount, -1)

	w.WriteHeader(http.StatusNoContent)
}

// Get total number of users
func handleGetUserCount(w http.ResponseWriter, r *http.Request) {
	count := atomic.LoadUint64(&userCount)
	fmt.Fprintf(w, "%d", count)
}

func main() {
	router := mux.NewRouter()

	// Define routes
	router.HandleFunc("/users", handleCreateUser).Methods("POST")
	router.HandleFunc("/users/{id}", handleGetUser).Methods("GET")
	router.HandleFunc("/users/{id}", handleUpdateUser).Methods("PUT")
	router.HandleFunc("/users/{id}", handleDeleteUser).Methods("DELETE")
	router.HandleFunc("/user-count", handleGetUserCount).Methods("GET")

	log.Println("Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
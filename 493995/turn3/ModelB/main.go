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
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

// We will use a sync.Map instead of a map[int]User to handle concurrent updates safely.
var users sync.Map

func main() {
	// Routing configuration remains the same...
}

// Use a RWMutex to manage concurrent access to the sync.Map
var rwMutex sync.RWMutex

// Create a new user
func handleCreateUser(w http.ResponseWriter, r *http.Request) {
	var user User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Error decoding JSON request body", http.StatusBadRequest)
		return
	}

	// Aquire the write lock before updating the map
	rwMutex.Lock()
	defer rwMutex.Unlock()

	users.Store(user.ID, user)

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

	//Aquire the read lock before accessing the map
	rwMutex.RLock()
	defer rwMutex.RUnlock()

	user, ok := users.Load(id)

	if !ok {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(user.(User))
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
	//Aquire the write lock before updating the map
	rwMutex.Lock()
	defer rwMutex.Unlock()
	
	if _, ok := users.Load(id); !ok {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}
    users.Store(id, user)

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
    //Aquire the write lock before deleting the map
    rwMutex.Lock()
	defer rwMutex.Unlock()
	
	if _, ok := users.Load(id); !ok {
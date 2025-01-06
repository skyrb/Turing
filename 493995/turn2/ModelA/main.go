package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
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

func handleCreateUser(w http.ResponseWriter, r *http.Request) {
	var user User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Error decoding JSON request body", http.StatusBadRequest)
		return
	}

	m.Lock()
	defer m.Unlock()
	users[user.ID] = user

	w.WriteHeader(http.StatusCreated)
	fmt.Fprint(w, user)
}

func handleGetUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Error parsing user ID", http.StatusBadRequest)
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

	m.Lock()
	defer m.Unlock()
	users[id] = user

	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, user)
}

func handleDeleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Error parsing user ID", http.StatusBadRequest)
		return
	}

	m.Lock()
	defer m.Unlock()
	delete(users, id)

	w.WriteHeader(http.StatusNoContent)
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/users", handleCreateUser).Methods("POST")
	mux.HandleFunc("/users/{id}", handleGetUser).Methods("GET")
	mux.HandleFunc("/users/{id}", handleUpdateUser).Methods("PUT")
	mux.HandleFunc("/users/{id}", handleDeleteUser).Methods("DELETE")

	log.Println("Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
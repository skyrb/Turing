package main

import (
	"encoding/json"
	"net/http"
)

// User struct represents a user's details
type User struct {
	ID        int64   `json:"id"`
	Username  string  `json:"username"`
	Email     string  `json:"email"`
	Profile   Profile `json:"profile"`
	Address   Address `json:"address"`
	Orders    []Order `json:"orders"`
}

// Profile struct represents the user's profile
type Profile struct {
	Bio      string `json:"bio"`
	Website  string `json:"website"`
	Location string `json:"location"`
}

// Address struct represents the user's address
type Address struct {
	Street   string  `json:"street"`
	City     string  `json:"city"`
	Zip      string  `json:"zip"`
	Country  string  `json:"country"`
}

// Order struct represents a user's order
type Order struct {
	OrderID int64   `json:"order_id"`
	Item    string  `json:"item"`
	Quantity int    `json:"quantity"`
	Date    string  `json:"date"`
}

// exampleUser creates a sample user with nested data
func exampleUser() *User {
	user := &User{
		ID:        1,
		Username:  "john_doe",
		Email:     "john@example.com",
		Profile: Profile{
			Bio:      "I'm a developer.",
			Website:  "http://example.com",
			Location: "New York, USA",
		},
		Address: Address{
			Street:   "123 Main St",
			City:     "New York",
			Zip:      "10001",
			Country:  "USA",
		},
		Orders: []Order{
			{OrderID: 101, Item: "Laptop", Quantity: 1, Date: "2023-09-15"},
			{OrderID: 102, Item: "Mouse", Quantity: 2, Date: "2023-09-20"},
		},
	}
	return user
}

// handler function to serve the user data in JSON format
func handler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		user := exampleUser()

		jsonData, err := json.Marshal(user)
		if err != nil {
			http.Error(w, "error marshalling JSON", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonData)
	} else {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}
}

func main() {
	http.HandleFunc("/user", handler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
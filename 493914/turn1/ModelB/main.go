package main

import (
	"fmt"
	"math/rand"
	"time"
)

type User struct {
	ID   int
	Name string
}

// UserBuilder is a struct that helps build a User object
type UserBuilder struct {
	user User
}

// NewUserBuilder returns a new UserBuilder
func NewUserBuilder() *UserBuilder {
	return &UserBuilder{}
}

// WithID sets the ID of the User
func (b *UserBuilder) WithID(id int) *UserBuilder {
	b.user.ID = id
	return b
}

// WithName sets the Name of the User
func (b *UserBuilder) WithName(name string) *UserBuilder {
	b.user.Name = name
	return b
}

// Build returns the constructed User object
func (b *UserBuilder) Build() User {
	return b.user
}

func getUserFromDB(id int, callback func(User, error)) {
	go func() {
		time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)

		// Simulate a database query
		user := NewUserBuilder().
			WithID(id).
			WithName(fmt.Sprintf("User %d", id)).
			Build()

		callback(user, nil)
	}()
}

func main() {
	getUserFromDB(1, func(user User, err error) {
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		fmt.Println("User:", user)
	})
}
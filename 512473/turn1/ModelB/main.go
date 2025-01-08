package main

import (
	"fmt"
)

type User struct {
	Name string
	Age  int
}

func main() {
	users := []User{
		{Name: "Alice", Age: 25},
		{Name: "Bob", Age: 30},
		{Name: "Charlie", Age: 22},
	}

	fmt.Println("Users:")
	for _, user := range users {
		fmt.Printf("%10s: %d\n", user.Name, user.Age)
	}
}
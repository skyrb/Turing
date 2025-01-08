package main

import (
	"fmt"
	"strings"
)

type Address struct {
	City    string
	ZipCode string
}

type User struct {
	Name    string
	Age     int
	Address Address
}

func printUsersTable(users []User) {
	// Determine the maximum length of each column
	nameLength := len("Name")
	ageLength := len("Age")
	cityLength := len("City")
	zipLength := len("Zip")

	for _, user := range users {
		if len(user.Name) > nameLength {
			nameLength = len(user.Name)
		}
		if len(user.Address.City) > cityLength {
			cityLength = len(user.Address.City)
		}
		if len(user.Address.ZipCode) > zipLength {
			zipLength = len(user.Address.ZipCode)
		}
	}

	// Print header
	fmt.Printf("%*s | %*s | %*s | %*s\n", nameLength, "Name", ageLength, "Age", cityLength, "City", zipLength, "Zip")
	fmt.Printf("%*s-%-s-%-s-%-s\n", nameLength, "-", ageLength, "-", cityLength, "-", zipLength, "-")

	// Print users
	for _, user := range users {
		fmt.Printf("%*s | %d | %*s | %*s\n", nameLength, user.Name, user.Age, cityLength, strings.ReplaceAll(user.Address.City, "", "N/A"), zipLength, strings.ReplaceAll(user.Address.ZipCode, "", "N/A"))
	}
}

func main() {
	users := []User{
		{
			Name: "Alice",
			Age:  25,
			Address: Address{
				City:    "New York",
				ZipCode: "10001",
			},
		},
		{
			Name: "Bob",
			Age:  30,
			Address: Address{
				City:    "Los Angeles",
				ZipCode: "90001",
			},
		},
		{
			Name: "Charlie",
			Age:  22,
			Address: Address{
				City:    "Chicago",
				ZipCode: "60616",
			},
		},
		{
			Name: "Sudhir",
			Age:  40,
			Address: Address{
				City:    "",
				ZipCode: "",
			},
		},
	}

	printUsersTable(users)
}
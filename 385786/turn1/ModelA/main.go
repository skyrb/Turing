package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

// MyData is the data structure used to populate the template.
type MyData struct {
	Name     string
	Age      int
	Greeting string
}

// capitalize is a custom function to capitalize the first letter of a string.
func capitalize(s string) string {
	if len(s) == 0 {
		return s
	}
	return string(s[0:1]) + s[1:]
}

// greet is a custom function to generate a greeting message based on age.
func greet(age int) string {
	if age < 18 {
		return "Hello, Youngster!"
	} else {
		return "Hello!"
	}
}

func main() {
	t, err := template.ParseFiles("index.html")
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data := &MyData{
			Name:     "alice",
			Age:      17,
			Greeting: greet(17),
		}

		if err := t.Execute(w, data); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	fmt.Println("Server is running on http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
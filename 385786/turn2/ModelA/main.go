package main

import (
	"fmt"
	"net/http"
	"text/template"
)

type Person struct {
	Name string
	Age  int
}

func main() {
	// Define a custom function for template transformation
	funcs := template.FuncMap{
		"ageToString": func(age int) string {
			if age < 0 {
				return "Unknown"
			}
			return fmt.Sprintf("%d years old", age)
		},
	}

	// Parse the template from the index.html file
	tpl, err := template.ParseFiles("index.html")
	if err != nil {
		panic(err)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		person := Person{
			Name: "Alice",
			Age:  25,
		}
		if err := tpl.Execute(w, person); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})

	fmt.Println("Server running on http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
package main

import (
	"fmt"
	"net/http"
	"path/filepath"
	"text/template"
)

type Person struct {
	Name string
	Age  int
}

func main() {
	// Define the templates directory
	tplDir := "templates"

	// Define a custom function for template transformation
	funcs := template.FuncMap{
		"ageToString": func(age int) string {
			if age < 0 {
				return "Unknown"
			}
			return fmt.Sprintf("%d years old", age)
		},
	}

	// Parse the template from the file system
	tpl := template.Must(template.New("").Funcs(funcs).ParseGlob(filepath.Join(tplDir, "*.html")))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		person := Person{
			Name: "Bob",
			Age:  -1,
		}

		if err := tpl.ExecuteTemplate(w, "index", person); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})

	fmt.Println("Server running on http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
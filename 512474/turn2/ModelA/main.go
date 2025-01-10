package main

import (
	"fmt"
	"sort"
)

// Define a struct to represent a student with a name and score
type Student struct {
	Name string
	Score int
}

func main() {
	students := []Student{
		{"Alice", 85},
		{"Charlie", 90},
		{"Bob", 75},
	}

	// Sort the students by score in descending order
	sort.Slice(students, func(i, j int) bool {
		return students[i].Score > students[j].Score
	})

	// Format and print the output as readable name-score pairs
	for _, student := range students {
		fmt.Printf("%s: %d\n", student.Name, student.Score)
	}
}
package main

import (
	"fmt"
	"sort"
)

// Define a struct to represent a student with a name, class, and score
type Student struct {
	Name  string
	Class string
	Score int
}

func main() {
	students := []Student{
		{"Alice", "Class A", 85},
		{"Charlie", "Class B", 90},
		{"Bob", "Class A", 75},
	}

	// Sort the students first by class, then by score within each class in descending order
	sort.Slice(students, func(i, j int) bool {
		if students[i].Class != students[j].Class {
			return students[i].Class < students[j].Class
		}
		return students[i].Score > students[j].Score
	})

	// Create a map to collect students by class
	classMap := make(map[string][]Student)
	for _, student := range students {
		classMap[student.Class] = append(classMap[student.Class], student)
	}

	// Format and print the output
	for class, studentsInClass := range classMap {
		// Sort students within the class by score in descending order
		sort.Slice(studentsInClass, func(i, j int) bool {
			return studentsInClass[i].Score > studentsInClass[j].Score
		})

		// Format and print the class and students with scores
		studentStrings := make([]string, len(studentsInClass))
		for i, student := range studentsInClass {
			studentStrings[i] = fmt.Sprintf("%s - %d", student.Name, student.Score)
		}
		fmt.Printf("%s: %s\n", class, fmt.Sprintf("[%s]", fmt.Sprint(studentStrings)))
	}
}
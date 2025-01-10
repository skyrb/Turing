package main

import (
	"fmt"
	"sort"
)

func main() {
	students := []string{"Alice", "Charlie", "Bob"}

	// Sort the list of student names alphabetically
	sort.Strings(students)

	// Format the sorted list into a single string with each name separated by a comma and a space
	sortedStudentsString := fmt.Sprintf("%s", students)

	fmt.Println("Sorted student names:", sortedStudentsString)
}
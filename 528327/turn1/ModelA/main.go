package main

import (
	"fmt"
)

func main() {
	project := "Smart Home System"
	owner := "Alice Johnson"
	designStatus := "Completed"
	developmentStatus := "In Progress"
	testingStatus := "Pending"

	result := fmt.Sprintf(
		"Project: %s\n"+
			"Owner: %s\n"+
			"Tasks:\n"+
			"    - Design: %s\n"+
			"    - Development: %s\n"+
			"    - Testing: %s\n",
		project, owner, designStatus, developmentStatus, testingStatus)

	fmt.Print(result)
}
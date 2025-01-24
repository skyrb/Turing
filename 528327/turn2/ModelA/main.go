package main

import (
	"fmt"
)

type Project struct {
	Name         string
	Owner        string
	Status       string
	Completion   int
}

func main() {
	projects := []Project{
		{"Smart Home System", "Alice Johnson", "In Progress", 65},
		{"AI Chatbot", "Bob Lee", "Completed", 100},
		{"Autonomous Vehicle", "Clara Smith", "Pending", 0},
	}

	fmt.Println("Projects Summary:")
	fmt.Println("-----------------------------------------------------------------------")
	fmt.Printf("%-30s %-25s %-20s %-15s\n", "Project Name", "Owner", "Status", "Completion (%)")

	for _, project := range projects {
		fmt.Printf("%-30s %-25s %-20s %-15d\n", project.Name, project.Owner, project.Status, project.Completion)
	}
	fmt.Println("-----------------------------------------------------------------------")
}
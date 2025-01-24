package main

import (
	"fmt"
	"strings"
	"time"
)

type Project struct {
	Name       string
	Owner      string
	Status     string
	Completion int
}

func main() {
	// Sample data
	projects := []Project{
		{"Smart Home System", "Alice Johnson", "In Progress", 65},
		{"AI Chatbot", "Bob Lee", "Completed", 100},
		{"Autonomous Vehicle", "Clara Smith", "Pending", 0},
	}

	// Calculate statistics
	totalProjects := len(projects)
	completedProjects := 0
	totalCompletion := 0

	for _, project := range projects {
		totalCompletion += project.Completion
		if project.Status == "Completed" {
			completedProjects++
		}
	}

	averageCompletion := 0
	if totalProjects > 0 {
		averageCompletion = totalCompletion / totalProjects
	}

	// Define column widths for consistent formatting
	projectNameWidth := 30
	ownerWidth := 20
	statusWidth := 20
	completionWidth := 15
	totalWidth := projectNameWidth + ownerWidth + statusWidth + completionWidth + 10

	// Header
	currentDate := time.Now().Format("2006-01-02")
	fmt.Printf("Project Report - Generated on: %s\n\n", currentDate)

	// Table Header
	fmt.Println("Projects Summary:")
	fmt.Println(strings.Repeat("-", totalWidth))
	fmt.Printf(
		"%-*s %-*s %-*s %-*s\n",
		projectNameWidth, "Project Name",
		ownerWidth, "Owner",
		statusWidth, "Status",
		completionWidth, "Completion (%)",
	)
	fmt.Println(strings.Repeat("-", totalWidth))

	// Table Rows
	for _, project := range projects {
		fmt.Printf(
			"%-*s %-*s %-*s %-*d\n",
			projectNameWidth, project.Name,
			ownerWidth, project.Owner,
			statusWidth, project.Status,
			completionWidth, project.Completion,
		)
	}

	// Table Footer
	fmt.Println(strings.Repeat("-", totalWidth))

	// Detailed Statistics
	fmt.Printf("Total Projects: %d\n", totalProjects)
	fmt.Printf("Completed Projects: %d\n", completedProjects)
	fmt.Printf("Average Completion: %d%%\n", averageCompletion)
}
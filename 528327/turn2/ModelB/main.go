package main

import (
	"fmt"
)

func main() {
	projects := []struct {
		Name   string
		Owner  string
		Status string
		CompletionPercent float64
	}{
		{
			Name:   "Smart Home System",
			Owner:  "Alice Johnson",
			Status: "In Progress",
			CompletionPercent: 65.0,
		},
		{
			Name:   "AI Chatbot",
			Owner:  "Bob Lee",
			Status: "Completed",
			CompletionPercent: 100.0,
		},
		{
			Name:   "Autonomous Vehicle",
			Owner:  "Clara Smith",
			Status: "Pending",
			CompletionPercent: 0.0,
		},
	}

	const projectSummaryTemplate = `Projects Summary:
-----------------------------------------------------------------------
Project Name                    Owner                       Status                 Completion (%)
%s
-----------------------------------------------------------------------`

	var result string
	for _, project := range projects {
		projectRow := fmt.Sprintf("%-30s %-25s %-15s %10.1f%%\n", project.Name, project.Owner, project.Status, project.CompletionPercent)
		result += projectRow
	}

	finalResult := fmt.Sprintf(projectSummaryTemplate, result)
	fmt.Print(finalResult)
}
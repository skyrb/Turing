package main  
import (  
    "fmt"
    "strings"
    "time"
)

type Project struct {  
    Name        string
    Owner       string
    Status      string
    Completion int
}

func main() {  
    projects := []Project{  
        {"Smart Home System", "Alice Johnson", "In Progress", 65},  
        {"AI Chatbot", "Bob Lee", "Completed", 100},  
        {"Autonomous Vehicle", "Clara Smith", "Pending", 0},  
    }  
    // Report heading with current date
    t := time.Now()
    header := fmt.Sprintf("Project Report - Generated on: %d-%02d-%02d", t.Year(), t.Month(), t.Day())
    footer := `
Total Projects: %d
Completed Projects: %d
Average Completion: %.1f%%`
    
    // Define column widths 
    projectNameWidth := 30
    ownerWidth := 20
    statusWidth := 15 
    completionWidth := 15 
    
    // Calculate total table width 
    totalWidth := projectNameWidth + ownerWidth + statusWidth + completionWidth + 10
    
    fmt.Println(header)
    fmt.Println(strings.Repeat("-", len(header))) 
    // Print table header 
    fmt.Printf("%-*s %-*s %-*s %-*s\n", projectNameWidth, "Project Name", ownerWidth, "Owner", statusWidth, "Status", completionWidth, "Completion (%)") 
    fmt.Println(strings.Repeat("-", totalWidth)) 
    
    // Print table rows
    totalCompleted := 0
    totalProjects := len(projects)
    var totalCompletion int

    for _, project := range projects {
        fmt.Printf("%-*s %-*s %-*s %-*d%%\n", projectNameWidth, project.Name, ownerWidth, project.Owner, statusWidth, project.Status, completionWidth, project.Completion)
        if project.Status == "Completed" {
            totalCompleted++
        }
        totalCompletion += project.Completion
    } 

    // Print footer
    fmt.Println(strings.Repeat("-", totalWidth))
    averageCompletion := float64(totalCompletion) / float64(totalProjects)
    fmt.Printf(footer, totalProjects, totalCompleted, averageCompletion)
    fmt.Println()
}
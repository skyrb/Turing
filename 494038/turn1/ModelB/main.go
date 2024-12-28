package main  
import (  
    "fmt"
)  
func main() {  
    emailTemplate := "Dear %s,\n\nThank you so much for your recent purchase. Your support means the world to us!\n\nBest regards,\n[Your Name]\n[Your Company Name]\n[Contact Information]"
    recipientName := "Alice"
    
    personalizedEmail := fmt.Sprintf(emailTemplate, recipientName)
    
    fmt.Println(personalizedEmail)
}  
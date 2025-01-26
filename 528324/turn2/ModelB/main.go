package main  
import (  
    "fmt"
    "regexp"
    "strings"
)

// MaskEmail partially masks an email address.
func MaskEmail(email string) string {
    // Use a regex to identify the local part of the email.
    re := regexp.MustCompile(`(^[^@]+)@`)
    matches := re.FindStringSubmatch(email)
    if len(matches) > 1 {
        localPart := matches[1]
        if len(localPart) > 2 {
            return localPart[:1] + strings.Repeat("*", len(localPart)-2) + localPart[len(localPart)-1:] + "@" + email[len(localPart)+1:]
        }
    }
    // Return the original email if it's incomplete or invalid
    return email
}

// MaskPhone partially masks a phone number.
func MaskPhone(phone string) string {
    if len(phone) > 4 {
        return strings.Repeat("*", len(phone)-4) + phone[len(phone)-4:]
    }
    // Return the original phone number if it's incomplete or invalid
    return phone
}

// RedactSensitiveInfo redacts sensitive information in a given string (email or log message).
func RedactSensitiveInfo(input string) string {
    maskedEmail := MaskEmail(input)
    _ := MaskPhone(input)
    // Return the string with the most sensitive information redacted
    return strings.Replace(input, maskedEmail, "***", -1)
}

// GenerateEmail creates an email template with masked information.
func GenerateEmail(template, userName, email, phone string) string {
    maskedEmail := MaskEmail(email)
    maskedPhone := MaskPhone(phone)
    return fmt.Sprintf(template, userName, maskedEmail, maskedPhone)
}

// LogEmail logs the generated email with redacted sensitive information.
func LogEmail(template, userName, email, phone string) {
    redactedEmail := RedactSensitiveInfo(email)
    logMessage := fmt.Sprintf(template, userName, redactedEmail, MaskPhone(phone))
    fmt.Println("Log:", logMessage)
}

func main() {  
    template := "Hello %[1]s,\n\nThank you for contacting support. We have received your request and will get back to you shortly.\n\nYour registered email is: %[2]s\nYour contact number ends with: %[3]s\n\nBest regards,\nCustomer Support"
    
    userName := "John Doe"
    email := "john.doe@example.com"
    phone := "1234567890"
    
    _ := GenerateEmail(template, userName, email, phone)
    LogEmail(template, userName, email, phone)
    
    // Test cases with incomplete or invalid data
    testEmail1 := "johndoe" // Incomplete email
    testEmail2 := "@example.com" // Incomplete email
    testPhone1 := "123" // Incomplete phone number
    testPhone2 := "InvalidPhone" // Non-numeric phone number
    
    fmt.Println("Generated Email for testEmail1:", GenerateEmail(template, userName, testEmail1, phone))
    fmt.Println("Generated Email for testEmail2:", GenerateEmail(template, userName, testEmail2, phone))
    fmt.Println("Generated Email for testPhone1:", GenerateEmail(template, userName, email, testPhone1))
    fmt.Println("Generated Email for testPhone2:", GenerateEmail(template, userName, email, testPhone2))
    
    LogEmail(template, userName, testEmail1, phone)
    LogEmail(template, userName, testEmail2, phone)
    LogEmail(template, userName, email, testPhone1)
    LogEmail(template, userName, email, testPhone2)
}
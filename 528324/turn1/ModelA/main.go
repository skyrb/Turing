package main

import (
	"fmt"
	"regexp"
)

// MaskEmail partially masks an email address.
func MaskEmail(email string) string {
	// Use a regex to identify the local part of the email.
	re := regexp.MustCompile(`(^[^@]+)@`)
	matches := re.FindStringSubmatch(email)
	if len(matches) > 1 {
		localPart := matches[1]
		if len(localPart) > 2 {
			return localPart[:1] + "****" + localPart[len(localPart)-1:] + "@" + email[len(localPart)+1:]
		}
	}
	return email
}

// MaskPhone partially masks a phone number.
func MaskPhone(phone string) string {
	if len(phone) > 4 {
		return "*****" + phone[len(phone)-4:]
	}
	return phone
}

// GenerateEmail creates an email template with masked information.
func GenerateEmail(template, userName, email, phone string) string {
	maskedEmail := MaskEmail(email)
	maskedPhone := MaskPhone(phone)
	return fmt.Sprintf(template, userName, maskedEmail, maskedPhone)
}

func main() {
	template := "Hello %[1]s,\n\nThank you for contacting support. We have received your request and will get back to you shortly.\n\nYour registered email is: %[2]s\nYour contact number ends with: %[3]s\n\nBest regards,\nCustomer Support"

	userName := "John Doe"
	email := "john.doe@example.com"
	phone := "1234567890"

	emailContent := GenerateEmail(template, userName, email, phone)
	fmt.Println(emailContent)
}
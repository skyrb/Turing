package main

import (
	"fmt"
	"log"
	"regexp"
)

// EmailTemplate holds templates for different languages
var EmailTemplates = map[string]string{
	"en": "Hello %[1]s,\n\nThank you for contacting support. We have received your request and will get back to you shortly.\n\nYour registered email is: %[2]s\nYour contact number ends with: %[3]s\n\nBest regards,\nCustomer Support",
	"fr": "Bonjour %[1]s,\n\nMerci de contacter le support. Nous avons reçu votre demande et vous répondrons sous peu.\n\nVotre email enregistré est : %[2]s\nVotre numéro de contact se termine par : %[3]s\n\nCordialement,\nSupport Client",
	// Add more languages as needed
}

// MaskEmail partially masks an email address, handling invalid formats.
func MaskEmail(email string) string {
	re := regexp.MustCompile(`(^[^@]+)@([^@]+)$`)
	matches := re.FindStringSubmatch(email)
	if len(matches) == 3 {
		localPart := matches[1]
		domain := matches[2]
		if len(localPart) > 2 {
			return localPart[:1] + "****" + localPart[len(localPart)-1:] + "@" + domain
		}
		return "****@" + domain  // default to a basic mask if the local part is too short
	}
	return "****"  // entirely redacted in case of invalid email format
}

// MaskPhone partially masks a phone number, handling short numbers.
func MaskPhone(phone string) string {
	if len(phone) > 4 {
		return "*****" + phone[len(phone)-4:]
	}
	return "*****"  // fully masked if the phone number is too short or invalid
}

// GenerateEmail creates an email template with masked information for the recipient and logs.
func GenerateEmail(template, userName, email, phone string) string {
	maskedEmail := MaskEmail(email)
	maskedPhone := MaskPhone(phone)
	return fmt.Sprintf(template, userName, maskedEmail, maskedPhone)
}

// LogDebugInformation logs the essential details of generated emails with sensitive fields masked.
func LogDebugInformation(userName, email, phone string) {
	maskedEmail := MaskEmail(email)
	maskedPhone := MaskPhone(phone)
	logMessage := fmt.Sprintf("Generated email for user: %s, email: %s, phone: %s", userName, maskedEmail, maskedPhone)
	log.Println(logMessage)
}

func main() {
	userName := "John Doe"
	email := "john.doe@example.com"
	phone := "1234567890"
	language := "fr"  // Simulate a user preference for French

	// Retrieve the appropriate template based on language preference
	template, exists := EmailTemplates[language]
	if !exists {
		template = EmailTemplates["en"]  // Fallback to English if the preferred language is not supported
	}

	// Generate the email with masked details
	emailContent := GenerateEmail(template, userName, email, phone)
	fmt.Println("Email Content:\n", emailContent)

	// Log the debug information
	LogDebugInformation(userName, email, phone)
}
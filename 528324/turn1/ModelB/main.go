package main

import (
	"fmt"
	"github.com/coreos/go-log/log"
	"github.com/pkg/errors"
	"strings"
)

// MaskSensitiveData masks sensitive data in a given string.
func MaskSensitiveData(input string, mask string) string {
	// Define a list of sensitive data patterns to be masked.
	sensitiveDataPatterns := []string{
		"@.*?\\.", // Email addresses
		"\\d{3}-\\d{3}-\\d{4}", // Phone numbers
	}

	// Iterate through each pattern and replace it with the mask.
	for _, pattern := range sensitiveDataPatterns {
		re := regexp.MustCompile(pattern)
		input = re.ReplaceAllString(input, mask)
	}

	return input
}

// GenerateEmailTemplate generates an email template with masked sensitive data.
func GenerateEmailTemplate(template string, data map[string]string) (string, error) {
	// Create a new logger with a custom prefix.
	logger := log.New("email_template")

	// Validate the template and data.
	if template == "" {
		return "", errors.New("template cannot be empty")
	}
	if data == nil {
		return "", errors.New("data cannot be nil")
	}

	// Mask sensitive data in the template.
	maskedTemplate := MaskSensitiveData(template, "***")

	// Create a new buffer to store the formatted template.
	var buf bytes.Buffer

	// Execute the template with the provided data.
	err := tmpl.Execute(&buf, data)
	if err != nil {
		logger.Errorf("failed to execute template: %v", err)
		return "", errors.Wrap(err, "failed to execute template")
	}

	// Return the formatted template.
	return buf.String(), nil
}

func main() {
	// Define an email template with placeholders.
	template := `
Dear {{.UserName}},

Thank you for reaching out to our customer support. Your account details are as follows:

Account ID: {{.AccountID}}
Email Address: {{.EmailAddress}}
Phone Number: {{.PhoneNumber}}

We will respond to your query as soon as possible.

Best regards,

Customer Support
`

	// Define data for the template.
	data := map[string]string{
		"UserName":     "John Doe",
		"AccountID":    "123456",
		"EmailAddress": "johndoe@example.com",
		"PhoneNumber":  "123-456-7890",
	}

	// Generate the email template with masked sensitive data.
	email, err := GenerateEmailTemplate(template, data)
	if err != nil {
		fmt.Printf("Error generating email template: %v\n", err)
		return
	}

	// Print the generated email template.
	fmt.Println(email)
}
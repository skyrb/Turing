package main  
import (  
        "fmt"
        "log"
        "regexp"
        "strings"

        "golang.org/x/text/language"
        "golang.org/x/text/message"
)

// MaskEmail partially masks an email address based on the provided locale.
func MaskEmail(email string, locale string) string {
        re := regexp.MustCompile(`(^[^@]+)@([^@]+)$`)
        matches := re.FindStringSubmatch(email)
        if len(matches) == 3 {
                localPart := matches[1]
                domain := matches[2]
                if len(localPart) > 2 {
                        return localPart[:1] + "****" + localPart[len(localPart)-1:] + "@" + domain
                }
                return "****@" + domain
        }
        return "****"
}

// MaskPhone partially masks a phone number based on the provided locale.
func MaskPhone(phone string, locale string) string {
        if len(phone) > 4 {
                return "*****" + phone[len(phone)-4:]
        }
        return "*****"
}

// LoadLocalizedTemplate loads a localized template for the given locale.
func LoadLocalizedTemplate(locale string) string {
        // In a real application, you would load templates from a file or database.
        // For simplicity, we'll use a map here.
        templates := map[string]string{
                "en": "Hello %[1]s,\n\nThank you for contacting support. We have received your request and will get back to you shortly.\n\nYour registered email is: %[2]s\nYour contact number ends with: %[3]s\n\nBest regards,\nCustomer Support",
                "fr": "Bonjour %[1]s,\n\nMerci de nous avoir contactés. Nous avons reçu votre demande et nous vous répondrons bientôt.\n\nVotre adresse e-mail enregistrée est : %[2]s\nLe dernier chiffre de votre numéro de téléphone est : %[3]s\n\nMeilleures salutations,\nSupport Client",
        }

        template, ok := templates[locale]
        if !ok {
                // Default to English template if the locale is not found
                template = templates["en"]
        }
        return template
}

// GenerateEmail generates an email with masked information for the given locale.
func GenerateEmail(userName, email, phone, locale string) string {
        template := LoadLocalizedTemplate(locale)
        maskedEmail := MaskEmail(email, locale)
        maskedPhone := MaskPhone(phone, locale)
        p := message.NewPrinter(language.Make(locale))
        return p.Sprintf(template, userName, maskedEmail, maskedPhone)
}

// LogDebugInformation logs the essential details of generated emails with sensitive fields masked.
func LogDebugInformation(userName, email, phone, locale string) {
        maskedEmail := MaskEmail(email, locale)
        maskedPhone := MaskPhone(phone, locale)
        logMessage := fmt.Sprintf("Generated email for user: %s, email: %s, phone: %s (Locale: %s)", userName, maskedEmail, maskedPhone, locale)
        log.Println(logMessage)
}

func main() {
        // Example usage with different locales
        userName := "John Doe"
        email := "john.doe@example.com"
        phone := "1234567890"

        locales := []string{"en", "fr"}
        for _, locale := range locales {
                // Generate the email with masked details for the current locale
                emailContent := GenerateEmail(userName, email, phone, locale)
                fmt.Println("Email Content in", locale, ":\n", emailContent)

                // Log the debug information
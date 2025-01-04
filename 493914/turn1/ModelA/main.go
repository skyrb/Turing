package main

import "fmt"

type EmailSender interface {
	SendEmail(email Email)
}

type SimpleEmailSender struct{}

func (s *SimpleEmailSender) SendEmail(email Email) {
	// Simulate sending an email
	if email.To == "correct@example.com" {
		fmt.Println("Email sent successfully.")
	} else {
		fmt.Println("Failed to send email.")
	}
}

type Email struct {
	To       string
	Subject  string
	Body     string
	OnSuccess func()
	OnFailure func()
}

type EmailBuilder struct {
	email *Email
}

func NewEmailBuilder() *EmailBuilder {
	return &EmailBuilder{email: &Email{}}
}

func (b *EmailBuilder) WithTo(to string) *EmailBuilder {
	b.email.To = to
	return b
}

func (b *EmailBuilder) WithSubject(subject string) *EmailBuilder {
	b.email.Subject = subject
	return b
}

func (b *EmailBuilder) WithBody(body string) *EmailBuilder {
	b.email.Body = body
	return b
}

func (b *EmailBuilder) WithOnSuccess(onSuccess func()) *EmailBuilder {
	b.email.OnSuccess = onSuccess
	return b
}

func (b *EmailBuilder) WithOnFailure(onFailure func()) *EmailBuilder {
	b.email.OnFailure = onFailure
	return b
}

func (b *EmailBuilder) Build() Email {
	return *b.email
}

func main() {
	sender := &SimpleEmailSender{}

	email := NewEmailBuilder().WithTo("correct@example.com").
		WithSubject("Test Email").
		WithBody("This is a test email.").
		WithOnSuccess(func() {
			fmt.Println("Email sent successfully.")
		}).
		WithOnFailure(func() {
			fmt.Println("Failed to send email.")
		}).
		Build()

	sender.SendEmail(email)

	email = NewEmailBuilder().
		WithTo("incorrect@example.com").
		WithSubject("Test Email").
		WithBody("This is a test email.").
		WithOnSuccess(func() {
			fmt.Println("Email sent successfully.")
		}).
		WithOnFailure(func() {
			fmt.Println("Failed to send email.")
		}).
		Build()

	sender.SendEmail(email)
}
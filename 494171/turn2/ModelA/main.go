package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func main() {
	// Create a new Fyne application
	myApp := app.New()
	defer myApp.Quit() // Ensure application resources are cleaned up on exit

	// Create the main application window
	mainWindow := myApp.NewWindow("Form Example")
	defer mainWindow.Close() // Ensure window resources are cleaned up on exit

	// Set up input fields
	nameInput := widget.NewEntry()
	emailInput := widget.NewEntry()
	nameInput.SetPlaceHolder("Name")
	emailInput.SetPlaceHolder("Email")

	// Create a submit button
	submitButton := widget.NewButton("Submit", func() {
		submitForm(nameInput.Text, emailInput.Text)
	})

	// Display the form with a grid layout
	mainWindow.SetContent(container.NewGridWithColumns(1,
		container.NewHorizontalBox(
			widget.NewLabel("Name:"),
			nameInput,
		),
		container.NewHorizontalBox(
			widget.NewLabel("Email:"),
			emailInput,
		),
		container.NewCenterWidget(submitButton),
	))

	// Show the main window and start the application
	mainWindow.ShowAndRun()
}

func submitForm(name, email string) {
	// Check for empty fields
	if name == "" || email == "" {
		dialog.NewError("Fields Required", "Please fill out all required fields.").Show()
		return
	}

	// Acknowledge successful form submission
	dialog.NewInformation("Form Submitted", "Your form has been successfully submitted.").Show()

	// Simulate processing time or any post-submission action
	// (In this example, we simulate with a defer statement)
	defer func() {
		dialog.NewInformation("Processing Complete", "Form processing is complete.").Show()
	}()
}
package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"log"
)

func main() {
	myApp := app.New()
	defer myApp.Quit()

	mainWindow := myApp.NewWindow("GUI with Lifecycle Handling")
	defer mainWindow.Close()

	// Counter to keep track of the number of clicks
	clickCount := 0

	// Create a label to display the click count
	clickCountLabel := widget.NewLabel("Click Count: 0")

	// Create a button to increment the click count
	incrementButton := widget.NewButton("Increment", func() {
		clickCount++
		clickCountLabel.SetText("Click Count: " + string(clickCount))
	})

	// Create a button to open a confirmation dialog
	confirmButton := widget.NewButton("Confirm", func() {
		confirmDialog := dialog.NewConfirm("Confirm Action", "Are you sure?", func(confirmed bool) {
			if confirmed {
				dialog.NewInformation("Confirmed", "Action confirmed!", mainWindow).Show()
			} else {
				dialog.NewInformation("Cancelled", "Action cancelled.", mainWindow).Show()
			}
		}, mainWindow)
		confirmDialog.Show()
	})

	// Handle the main window closing event
	mainWindow.SetOnClose(func() {
		// Prompt the user to confirm before closing
		confirmDialog := dialog.NewConfirm("Close Application", "Are you sure you want to exit?", func(confirmed bool) {
			if confirmed {
				log.Println("Application closed by user.")
				myApp.Quit() // Close the application
			} else {
				mainWindow.Show() // Reopen the main window
			}
		}, mainWindow)
		confirmDialog.Show()
	})

	// Create a vertical box to hold the widgets
	mainWindow.SetContent(container.NewVBox(
		clickCountLabel,
		incrementButton,
		confirmButton,
	))

	mainWindow.ShowAndRun()
}
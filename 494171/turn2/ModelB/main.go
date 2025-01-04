

package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Lifecycle Event Demo")

	// Custom variable to keep track of the last selected status
	lastSelected := "None"

	// Create a radio group with some options
	options := []string{"Option 1", "Option 2", "Option 3"}
	radioGroup := widget.NewRadioGroup(options, func(value string) {
		lastSelected = value
		fmt.Println("Radio Group Selected:", value)
	})
	// Wrap the radio group in a scroll container to handle more options
	radioScroll := container.NewScroll(radioGroup)

	// Create a button to display the selected option
	showButton := widget.NewButton("Show Selected", func() {
		dialog.ShowMessage("Selected Option", lastSelected, w)
	})

	// Create a button to close the window programatically
	closeButton := widget.NewButton("Close Window", func() {
		w.Close() // Closes the window
	})

	// Register a lifecycle event handler for WindowClose to show a confirmation dialog
	w.SetOnClose(func() {
		response := dialog.ShowConfirm("Confirm", "Are you sure you want to close the window?", w)
		if response {
			w.SetClosed(true) // Proceed with closing the window
			// Any cleanup operations or deferred tasks should be performed here
			fmt.Println("Window Closed.")
		}
	})

	// Wrap the UI elements in a horizontal box
	content := container.NewVBox(
		radioScroll,
		layout.NewSpacer(), // Add some vertical space
		container.NewHBox(showButton, layout.NewSpacer(), closeButton),
	)

	// Apply the content to the window
	w.SetContent(content)

	// Add defer statement to handle cleanup tasks (e.g., saving settings)
	defer func() {
		fmt.Println("Performing cleanup tasks...")
		// Save user settings or preferences here before application exit
	}()

	// Show the window and start the application
	w.ShowAndRun()
}  

 
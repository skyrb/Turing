package main

import (
	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/container"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
)

func main() {
	// Create a new application instance
	a := app.New()
	defer a.Stop() // Ensure the application stops properly

	// Create the main window
	w := a.NewWindow("Hello Fyne!")
	defer w.Close() // Ensure the window is closed properly

	// Create a container to hold the widget
	cont := container.NewBox(layout.NewGridLayout(1, 2))

	// Create a button widget
	btn := widget.NewButton("Click Me", func() {
		// Create a dialog window
		dialog := fyne.NewDialog("Dialog Example", w)
		defer dialog.Close() // Ensure the dialog is closed properly

		// Create a container to hold the dialog content
		dialogCont := container.NewBox(layout.NewCenterLayout())

		// Add a label to the dialog
		dialogCont.AddObject(widget.NewLabel("You clicked the button!"))

		// Add a close button to the dialog
		dialogCont.AddObject(widget.NewButton("Close", func() {
			dialog.Hide()
		}))

		// Set the dialog content
		dialog.SetContent(dialogCont)

		// Show the dialog
		dialog.Show()
	})

	// Create a label widget
	label := widget.NewLabel("Hello, World!")

	// Add widgets to the main window container
	cont.AddObject(label)
	cont.AddObject(btn)

	// Set the window content
	w.SetContent(cont)

	// Show the main window
	w.ShowAndRun()
}

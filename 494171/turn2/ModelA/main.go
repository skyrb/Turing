package main

import (
	"fmt"
	"github.com/fyne-io/fyne/v2"
	"github.com/fyne-io/fyne/v2/app"
	"github.com/fyne-io/fyne/v2/container"
	"github.com/fyne-io/fyne/v2/dialog"
	"github.com/fyne-io/fyne/v2/widget"
)

var clickCount int

func main() {
	clickCount = 0
	a := app.New()
	w := a.NewWindow("Click Counter")

	// Set the main content
	w.SetContent(container.NewVBox(
		widget.NewLabelWithText("Click Count: "+fmt.Sprint(clickCount)),
		widget.NewButton("Click Me", func() {
			clickCount++
			w.Resize(w.MinSize()) // Ensure the label updates its size if needed
		}),
	))

	// Register a window close handler
	w.OnClose = func(ev interface{}) {
		if dialog.ShowConfirm("Confirm", "Are you sure you want to exit?", w) {
			// Window should be closed explicitly after confirmation
			w.Close()
		} else {
			// Cancel the close action
			ev.(fyne.Cancelable).Cancel()
		}
	}

	// Defer statement to display a closing message even if user cancels
	defer func() {
		fmt.Println("Application closing.")
	}()

	// Show the window and start the application
	w.ShowAndRun()
}
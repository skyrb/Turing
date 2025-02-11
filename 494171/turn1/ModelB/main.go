

package main  
import (  
    "fmt"
    "github.com/fyne-io/fyne/v2"
    "github.com/fyne-io/fyne/v2/app"
    "github.com/fyne-io/fyne/v2/container"
    "github.com/fyne-io/fyne/v2/dialog"
    "github.com/fyne-io/fyne/v2/widget"
)

func main() {  
    a := app.New()
    w := a.NewWindow("Fyne Demo")

    // Define a button with a click event handler
    btn := widget.NewButton("Click Me", func() {
        dialog.ShowInformation("Message", "Button clicked!", w)
    })

    // Create a content container for the button
    content := container.NewCenter(btn)

    // Set the content of the window
    w.SetContent(content)

    // Add a defer statement to close the window when the main function exits
    defer w.Close()

    // Show the window and start the application
    w.ShowAndRun()
}

  
package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	app := app.New()
	window := app.NewWindow("Hello")

	helloLabel := widget.NewLabel("Hello Fyne!")
	window.SetContent(container.NewVBox(
		helloLabel,
		widget.NewButton("Hi!", func() {
			helloLabel.SetText("Welcome :)")
		}),
	))

	window.ShowAndRun()
}

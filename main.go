package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/theme"
	"github.com/FeryrArdacon/datasaver-fyne/view"
)

func main() {
	application := app.NewWithID("FeryrArdacon.datasaver-fyne")
	application.SetIcon(theme.FyneLogo())
	window := application.NewWindow("Hello")

	window.SetContent(view.CreateRecordList())

	window.ShowAndRun()
}

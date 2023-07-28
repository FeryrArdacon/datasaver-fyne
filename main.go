package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/theme"
	"github.com/FeryrArdacon/datasaver-fyne/files"
	"github.com/FeryrArdacon/datasaver-fyne/view"
)

func main() {
	application := app.NewWithID("FeryrArdacon.datasaver-fyne")
	application.SetIcon(theme.FyneLogo())
	application.Settings().SetTheme(theme.LightTheme())
	window := application.NewWindow("Datensicherung")

	var records = []*files.Record{
		files.NewRecord("/my/first/source", "/my/first/target"),
		files.NewRecord("/my/second/source", "/my/second/target"),
	}

	window.SetContent(view.CreateRecordList(records))
	window.Resize(fyne.NewSize(860, 640))

	window.ShowAndRun()
}

package view

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func CreateRecordList() fyne.CanvasObject {
	helloLabel := widget.NewLabel("Hello Fyne!")
	return container.NewVBox(
		helloLabel,
		widget.NewButton("Hi!", func() {
			helloLabel.SetText("Welcome :)")
		}),
	)
}

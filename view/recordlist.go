package view

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/FeryrArdacon/datasaver-fyne/files"
)

func CreateRecordList(records []*files.Record) fyne.CanvasObject {
	columnTitles := []string{"Ordner", "Sicherung"}

	table := widget.NewTable(
		func() (int, int) {
			return len(records) + 1, 2
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("")
		},
		func(cellID widget.TableCellID, cellItem fyne.CanvasObject) {
			cellLabel := cellItem.(*widget.Label)

			if cellID.Row == 0 {
				cellLabel.SetText(columnTitles[cellID.Col])
				cellLabel.TextStyle.Bold = true
				return
			}

			recordID := cellID.Row - 1
			cellLabel.SetText(records[recordID].ToTuple()[cellID.Col])
		},
	)

	table.SetColumnWidth(0, 420)
	table.SetColumnWidth(1, 420)

	startButton := widget.NewButton("Sicherung starten", func() {})

	startButton.Importance = widget.HighImportance

	header := container.New(
		layout.NewGridLayout(2),
		widget.NewButton("Hinzufügen", func() {}),
		startButton,
	)

	footer := container.NewHBox()

	content := container.New(
		layout.NewBorderLayout(header, footer, nil, nil),
		header,
		footer,
		table,
	)

	return content
}

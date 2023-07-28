package view

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
	"github.com/FeryrArdacon/datasaver-fyne/files"
)

func CreateRecordList(records []*files.Record) fyne.CanvasObject {
	table := widget.NewTable(
		func() (int, int) {
			return len(records), 2
		},
		func() fyne.CanvasObject {
			cellLabel := widget.NewLabel("")
			cellLabel.Resize(fyne.NewSize(300, 24))
			return cellLabel
		},
		func(cellID widget.TableCellID, cellItem fyne.CanvasObject) {
			cellLabel := cellItem.(*widget.Label)
			cellLabel.SetText(records[cellID.Row].ToTuple()[cellID.Col])
			cellLabel.Resize(fyne.NewSize(300, 24))
		},
	)

	table.SetColumnWidth(0, 300)
	table.SetColumnWidth(1, 300)

	return table
}

package gui

import (
	models "tiles/internal/models"

	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func InitializeGUI(ui *models.UI, p models.Piano) {
	ui.Columns = 15
	ui.Rows = 15

	UpdateGrid(ui, p)

	coloraButton := widget.NewButton("Colora", func() {
		showColoraDialogTakeParam(ui.Window, p)
		UpdateGrid(ui, p)
	})

	ui.Window.SetContent(container.NewHBox(
		ui.Grid,
		coloraButton,
	))
}

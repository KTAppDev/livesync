package topbar

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func DisplayTopBar() fyne.CanvasObject {
	Livesync := canvas.NewText("🎹 LIVE SYNC", color.White)
	Livesync.TextSize = 36.0
	Livesync.Alignment = fyne.TextAlignCenter
	Livesync.TextStyle = fyne.TextStyle{Bold: true}

	topBar := widget.NewToolbar(
		widget.NewToolbarAction(theme.HelpIcon(), func() {}),
		widget.NewToolbarAction(theme.AccountIcon(), func() {}),
	)
	bar := container.NewHBox(container.NewPadded(Livesync), layout.NewSpacer(), topBar)
	blueRectangle := canvas.NewRectangle(&color.Black)
	blueRectangle.SetMinSize(fyne.NewSize(800, 60))

	return container.NewStack(blueRectangle, bar)

}

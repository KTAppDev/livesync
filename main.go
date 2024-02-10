package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
	"github.com/ktappdev/filesync/getDirs"
	"github.com/ktappdev/filesync/getFiles"
	"github.com/ktappdev/filesync/monitorFiles"
)

func main() {
	a := app.New()
	w := a.NewWindow("Hello World")

	w.SetContent(widget.NewLabel("Hello World!"))
	directory := "/Users/kentaylor/Downloads/"
	getDirs.GetDirectories(directory)
	getFiles.GetFiles(directory)
	monitorFiles.MonitorFiles(directory)

	w.ShowAndRun()
}

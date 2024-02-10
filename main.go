package main

import (
	"fyne.io/fyne/v2/app"
	"github.com/ktappdev/filesync/data"
	"github.com/ktappdev/filesync/getDirs"
	"github.com/ktappdev/filesync/getFiles"
	"github.com/ktappdev/filesync/monitorFiles"
	"github.com/ktappdev/filesync/ui"
)

func main() {
	a := app.New()
	files := data.GetFakeFiles()
	w := ui.NewFileManagerUI(a, files)

	directory := "/Users/kentaylor/Downloads/"
	getDirs.GetDirectories(directory)
	getFiles.GetFiles(directory)
	go monitorFiles.MonitorFiles(directory)

	w.Run()
}

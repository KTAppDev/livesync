package main

import (
	"fyne.io/fyne/v2/app"
	"github.com/ktappdev/filesync/getDirs"
	"github.com/ktappdev/filesync/getFiles"
	"github.com/ktappdev/filesync/models"
	// "github.com/ktappdev/filesync/monitorFiles"
	"github.com/ktappdev/filesync/ui"
	"time"
)

func main() {
	a := app.New()
	files := []models.FileInfo{
		{"Lil King Project", 122, 147.00, "Hip-Hop", "WIP", "C#", "S", "02/06/24", time.Now(), time.Now()},
		{"Some Name", 123, 120.0, "Genre", "Status", "Key", "Grade", "Release Date", time.Now(), time.Now()},
	}
	w := ui.NewFileManagerUI(a, files)

	directory := "/Users/kentaylor/Downloads/"
	getDirs.GetDirectories(directory)
	getFiles.GetFiles(directory)
	// monitorFiles.MonitorFiles(directory)

	w.Run()
}

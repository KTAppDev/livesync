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
		{"Soulful Journey", 232, 109.44, "Jazz", "Upcoming", "Bb", "S", "02/06/24", time.Now(), time.Now()},
		{"Retro Vibes", 232, 157.41, "Electronic", "Released", "Cb", "D", "07/08/24", time.Now(), time.Now()},
		{"Soulful Journey", 111, 150.37, "Electronic", "Upcoming", "E#", "B", "01/01/24", time.Now(), time.Now()},
		{"Lil King Project", 736, 143.54, "Classical", "Upcoming", "Cb", "D", "01/01/24", time.Now(), time.Now()},
	}
	w := ui.NewFileManagerUI(a, files)

	directory := "/Users/kentaylor/Downloads/"
	getDirs.GetDirectories(directory)
	getFiles.GetFiles(directory)
	// monitorFiles.MonitorFiles(directory)

	w.Run()
}

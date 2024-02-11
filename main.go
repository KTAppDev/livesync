package main

import (
	"fmt"
	"time"

	"fyne.io/fyne/v2/app"
	"github.com/ktappdev/filesync/data"
	"github.com/ktappdev/filesync/getDirs"
	"github.com/ktappdev/filesync/getFiles"
	"github.com/ktappdev/filesync/monitorFiles"
	"github.com/ktappdev/filesync/ui"
)

func main() {
	files := data.GetFakeFiles()
	a := app.New()
	w := ui.NewFileManagerUI(a, files)

	directory := "/Users/kentaylor/Downloads/"
	getDirs.GetDirectories(directory)
	allFiles, err := getFiles.GetFiles(directory)
	if err != nil {
		fmt.Println(err)
	}
	for _, fileInfo := range allFiles {
		fmt.Printf("File: %s, Size: %d, Permissions: %s, Modified: %s\n",
			fileInfo.Path, fileInfo.Size, fileInfo.Permissions, fileInfo.Modified.Format(time.RFC3339))
	}

	go monitorFiles.MonitorFiles(directory)

	w.Run()
}

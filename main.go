package main

import (
	"database/sql"
	"fmt"
	"log"

	"fyne.io/fyne/v2/app"
	// "github.com/ktappdev/filesync/data"
	// "github.com/ktappdev/filesync/getDirs"
	"github.com/ktappdev/filesync/database"
	"github.com/ktappdev/filesync/getFiles"
	"github.com/ktappdev/filesync/monitorFiles"
	"github.com/ktappdev/filesync/ui"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	directory := "/Users/kentaylor/Downloads/"
	// Open a connection to the SQLite database file
	db, err := sql.Open("sqlite3", "./file_manager.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Initialize the database
	err = database.InitDB(db)
	if err != nil {
		log.Fatal("Error initializing database: ", err)
	}

	// getDirs.GetDirectories(directory)
	allFiles, err := getFiles.GetFiles(directory)
	if err != nil {
		fmt.Println(err)
	}

	go monitorFiles.MonitorFiles(directory)

	a := app.New()
	w := ui.NewFileManagerUI(a, allFiles)

	w.Run()
}

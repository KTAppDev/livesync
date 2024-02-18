package main

import (
	"database/sql"
	"log"

	"fyne.io/fyne/v2/app"
	"github.com/ktappdev/filesync/database"
	"github.com/ktappdev/filesync/getFiles"
	"github.com/ktappdev/filesync/ui"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	directory := "/Users/kentaylor/Downloads/"

	// Open a connection to the SQLite database file
	db, err := sql.Open("sqlite3", "./file_manager.db")
	if err != nil {
		log.Fatal("Error opening database connection:", err)
	}
	defer db.Close()

	// Initialize the database and create the table if it doesn't exist
	err = database.InitDB(db)
	if err != nil {
		log.Fatal("Error initializing database:", err)
	}

	// Retrieve files from the specified directory
	allFiles, err := getFiles.GetFiles(directory)
	if err != nil {
		log.Fatal("Error retrieving files:", err)
	}

	// Insert or update files in the database
	err = database.InsertFilesIntoDB(db, allFiles)
	if err != nil {
		log.Fatal("Error inserting files into database:", err)
	}

	// Retrieve all files from the database
	allDbFiles, err := database.GetAllFilesFromDB(db)
	if err != nil {
		log.Fatal("Error retrieving files from database:", err)
	}

	// Initialize the Fyne app and UI
	a := app.New()
	w := ui.NewFileManagerUI(a, allDbFiles)

	// Run the Fyne app
	w.Run()
}

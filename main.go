package main

import (
	"database/sql"
	"log"
	"os/user"

	"fyne.io/fyne/v2/app"
	"github.com/ktappdev/filesync/database"
	"github.com/ktappdev/filesync/getFiles"
	"github.com/ktappdev/filesync/parser"

	"github.com/ktappdev/filesync/logging"
	"github.com/ktappdev/filesync/theme"
	"github.com/ktappdev/filesync/ui"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	alsSourcePath := "/Users/kentaylor/developer/go-projects/livesync/ap/ap/ableton12.als"

	value, err := parser.ExtractALS(alsSourcePath)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Tempo Value: %s BPM\n", value)
	///////

	usr, err := user.Current()
	if err != nil {
		log.Fatal("Error getting executable path:", err)
	}

	homeDir := usr.HomeDir

	directory := homeDir + "/Desktop/"
	logging.Setup(homeDir + "/livesync")

	// Open a connection to the SQLite database file
	db, err := sql.Open("sqlite3", homeDir+"/livesync/file_manager.db")
	if err != nil {
	}
	defer db.Close()

	// Initialize the database and create the table if it doesn't exist
	err = database.InitDB(db)
	if err != nil {
		log.Fatal("Error initializing database:", err)
	}

	allFiles, err := getFiles.GetFiles(directory)
	if err != nil {
		log.Fatal("Error retrieving files:", err)
	}

	err = database.InsertFilesIntoDB(db, allFiles)
	if err != nil {
		log.Fatal("Error inserting files into database:", err)
	}

	allDbFiles, err := database.GetAllFilesFromDB(db)
	if err != nil {
		log.Fatal("Error retrieving files from database:", err)
	}

	a := app.New()
	a.Settings().SetTheme(theme.NewMyTheme())
	w := ui.NewFileManagerUI(a, allDbFiles)

	w.Run()
}

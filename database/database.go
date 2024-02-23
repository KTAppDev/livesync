package database

import (
	"database/sql"
)

// InitDB initializes the database by creating the "files" and "settings" tables if they do not exist.
func InitDB(db *sql.DB) error {
	// Initialize the "files" table
	err := InitFilesTable(db)
	if err != nil {
		return err
	}

	// Initialize the "settings" table
	err = InitSettingsTable(db)
	if err != nil {
		return err
	}

	return nil
}

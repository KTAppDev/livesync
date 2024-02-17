package database

import (
	"database/sql"
	"log"

	"github.com/ktappdev/filesync/models"
)

func InitDB(db *sql.DB) error {
	// Create the 'files' table
	createTableSQL := `CREATE TABLE IF NOT EXISTS files (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        name TEXT,
        size INTEGER,
        bpm REAL,
        genre TEXT,
        status TEXT,
        key TEXT,
        grade TEXT,
        release_date TEXT,
        created_at TEXT,
        modified_at TEXT
    );`
	_, err := db.Exec(createTableSQL)
	if err != nil {
		return err
	}
	log.Println("Database and table created successfully.")
	return nil
}

// Function to load files from the database
func LoadFilesFromDB(db *sql.DB) ([]models.FileInfo, error) {
	var files []models.FileInfo
	rows, err := db.Query("SELECT name, size, bpm, genre, status, key, grade, release_date, created_at, modified_at FROM files")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var file models.FileInfo
		err := rows.Scan(&file.Name, &file.Size, &file.BPM, &file.Genre, &file.Status, &file.Key, &file.Grade, &file.ReleaseDate, &file.CreatedAt, &file.Modified)
		if err != nil {
			return nil, err
		}
		files = append(files, file)
	}
	return files, nil
}

// Function to insert a file record into the 'files' table
func InsertFileIntoDB(db *sql.DB, file models.FileInfo) error {
	insertSQL := `
		INSERT INTO files (name, size, bpm, genre, status, key, grade, release_date, created_at, modified_at)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?);`

	_, err := db.Exec(
		insertSQL,
		file.Name, file.Size, file.BPM, file.Genre, file.Status,
		file.Key, file.Grade, file.ReleaseDate, file.CreatedAt, file.Modified,
	)
	if err != nil {
		return err
	}

	log.Println("File inserted successfully.")
	return nil
}

package database

import (
	"database/sql"
	"log"
	"time"

	"github.com/ktappdev/filesync/models"
)

// InitDB initializes the database by creating the "files" table if it does not exist.
func InitDB(db *sql.DB) error {
	// SQL statement to create the "files" table
	createTableSQL := `
		CREATE TABLE IF NOT EXISTS files (
			name TEXT UNIQUE,
			size INTEGER,
			bpm REAL,
			genre TEXT,
			status TEXT,
			key TEXT,
			grade TEXT,
			release_date TEXT, -- Use TEXT for simplicity, consider a proper DATE type in your database
			created_at DATETIME,
			updated_at DATETIME,
			path TEXT,
			permissions INTEGER, -- Use INTEGER for FileMode
			modified_at DATETIME,
			PRIMARY KEY (name)
		);
	`

	// Execute the SQL statement to create the table
	_, err := db.Exec(createTableSQL)
	if err != nil {
		return err
	}

	return nil
}

// GetAllFilesFromDB retrieves all rows from the "files" table and returns a slice of models.FileInfo.
func GetAllFilesFromDB(db *sql.DB) ([]models.FileInfo, error) {
	// SQL query to retrieve all fields from the "files" table
	query := `SELECT name, size, bpm, genre, status, key, grade, release_date, created_at, path, modified_at FROM files;`

	// Execute the query
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var files []models.FileInfo

	// Iterate through the result set
	for rows.Next() {
		var file models.FileInfo

		// Scan the row into the FileInfo struct
		err := rows.Scan(&file.Name, &file.Size, &file.BPM, &file.Genre, &file.Status, &file.Key, &file.Grade, &file.ReleaseDate, &file.CreatedAt, &file.Path, &file.ModifiedAt)
		if err != nil {
			// log.Println("Error scanning row:", err)
			continue
		}

		// Append the FileInfo struct to the files slice
		files = append(files, file)
	}

	// Check for any errors during iteration
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return files, nil
}

// InsertFilesIntoDB inserts or updates a slice of models.FileInfo into the "files" table.
func InsertFilesIntoDB(db *sql.DB, files []models.FileInfo) error {
	tx, err := db.Begin() // Start a transaction for efficiency
	if err != nil {
		return err
	}

	defer func() {
		if err != nil {
			tx.Rollback() // Rollback if any error occurs
		} else {
			tx.Commit() // Commit if all insertions/updates are successful
		}
	}()

	// Check database system and adapt placeholder for duplicate key check
	checkDuplicateSQL := `SELECT name, size, created_at FROM files WHERE name = ? AND size = ? AND created_at = ? LIMIT 1;`

	// Prepared statement for insertion/update
	insertSQL := `
		INSERT OR REPLACE INTO files (name, size, bpm, genre, status, key, grade, release_date, created_at, updated_at, path, permissions, modified_at)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
	`

	for _, file := range files {
		// Check if file already exists with the same name, size, and created_at
		row := tx.QueryRow(checkDuplicateSQL, file.Name, file.Size, file.CreatedAt)
		var existingName, existingSize, existingCreatedAt interface{}
		err = row.Scan(&existingName, &existingSize, &existingCreatedAt)

		if err != nil {
			// Check if the error is "no rows in result set"
			if err == sql.ErrNoRows {
				// File doesn't exist, insert
				_, err = tx.Exec(insertSQL, file.Name, file.Size, file.BPM, file.Genre, file.Status, file.Key, file.Grade, file.ReleaseDate, file.CreatedAt, file.UpdatedAt, file.Path, file.Permissions, file.ModifiedAt)
				if err != nil {
					return err
				}
				// log.Println("File", file.Name, "inserted successfully.")
			} else {
				// Handle other errors
				return err
			}
		} else {
			// File exists, update if necessary
			if file.ModifiedAt.After(existingCreatedAt.(time.Time)) {
				_, err = tx.Exec(insertSQL, file.Name, file.Size, file.BPM, file.Genre, file.Status, file.Key, file.Grade, file.ReleaseDate, file.CreatedAt, file.UpdatedAt, file.Path, file.Permissions, file.ModifiedAt)
				if err != nil {
					return err
				}
				log.Printf("File %s already exists, updated with newer version.", file.Name)
			} else {
				// log.Printf("File %s already exists with the same content, skipping update.", file.Name)
			}
		}
	}

	return nil
}

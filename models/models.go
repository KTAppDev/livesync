package models

import (
	"io/fs"
	"time"
)

// FileInfo represents information about a file.
type FileInfo struct {
	Name        string      // File name
	Size        int64       // File size in bytes
	BPM         float64     // Beats per minute
	Genre       string      // Genre of the file
	Status      string      // Status of the file
	Key         string      // Key of the file
	Grade       string      // Grade of the file
	ReleaseDate string      // Release date of the file
	CreatedAt   time.Time   // Creation time of the file
	UpdatedAt   time.Time   // Last update time of the file
	Path        string      // File path
	Permissions fs.FileMode // File permissions
	Modified    time.Time   // Last modified time
}

// NewFileInfoWithDefaults creates a new FileInfo instance with default values.
func (fi *FileInfo) NewFileInfoWithDefaults() {
	fi.Name = ""
	fi.Size = 0
	fi.BPM = 0
	fi.Genre = ""
	fi.Status = ""
	fi.Key = "C#"
	fi.Grade = ""
	fi.ReleaseDate = "This is the release date"
	fi.CreatedAt = time.Now()
	fi.UpdatedAt = time.Now()
	fi.Path = ""
	fi.Permissions = 0
	fi.Modified = time.Now()
}

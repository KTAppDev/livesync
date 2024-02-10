package data

import (
	"github.com/ktappdev/filesync/models"
	"time"
)

// GetFiles returns a slice of FileInfo.
func GetFakeFiles() []models.FileInfo {
	files := []models.FileInfo{
		{
			Name:        "Lil King Project",
			Size:        122,
			BPM:         147.00,
			Genre:       "Hip-Hop",
			Status:      "WIP",
			Key:         "C#",
			Grade:       "S",
			ReleaseDate: "02/06/24",
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
		{
			Name:        "Some Name",
			Size:        123,
			BPM:         120.0,
			Genre:       "Genre",
			Status:      "Status",
			Key:         "Key",
			Grade:       "Grade",
			ReleaseDate: "Release Date",
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
		{
			Name:        "Soulful Journey",
			Size:        232,
			BPM:         109.44,
			Genre:       "Jazz",
			Status:      "Upcoming",
			Key:         "Bb",
			Grade:       "S",
			ReleaseDate: "02/06/24",
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
		{
			Name:        "Retro Vibes",
			Size:        232,
			BPM:         157.41,
			Genre:       "Electronic",
			Status:      "Released",
			Key:         "Cb",
			Grade:       "D",
			ReleaseDate: "07/08/24",
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
		{
			Name:        "Soulful Journey",
			Size:        111,
			BPM:         150.37,
			Genre:       "Electronic",
			Status:      "Upcoming",
			Key:         "E#",
			Grade:       "B",
			ReleaseDate: "01/01/24",
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
		{
			Name:        "Lil King Project",
			Size:        736,
			BPM:         143.54,
			Genre:       "Classical",
			Status:      "Upcoming",
			Key:         "Cb",
			Grade:       "D",
			ReleaseDate: "01/01/24",
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
	}

	return files
}

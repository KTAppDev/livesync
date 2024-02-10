package models

import "time"

type FileInfo struct {
	Name        string
	Size        int
	BPM         float64
	Genre       string
	Status      string
	Key         string
	Grade       string
	ReleaseDate string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

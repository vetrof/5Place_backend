package models

import "time"

type City struct {
	ID   int
	Name string
	Geom string
}

// Place — структура результата
type Place struct {
	ID       int
	CityName string
	Name     string
	Geom     string
	Desc     string
	Distance float64
}

type Photo struct {
	ID           int       `json:"id"`
	PlaceID      int       `json:"place_id"`
	Filename     string    `json:"filename"`
	OriginalName string    `json:"original_name"`
	FilePath     string    `json:"file_path"`
	FileSize     int64     `json:"file_size"`
	MimeType     string    `json:"mime_type"`
	UploadedAt   time.Time `json:"uploaded_at"`
	Description  string    `json:"description"`
}

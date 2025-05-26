package models

type City struct {
	ID     int
	Name   string
	Geom   string
	Points int
}

// Place — структура результата
type Place struct {
	ID       int
	CityName string
	Name     string
	Geom     string
	Desc     string
	Distance float64
	Photos   []string
}

type Photo struct {
	ID          int    `json:"id"`
	PlaceID     int    `json:"place_id"`
	FileLink    string `json:"file_path"`
	Description string `json:"description"`
}

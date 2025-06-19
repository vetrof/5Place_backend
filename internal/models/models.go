package models

type City struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Geom   string `json:"geom"`
	Points int    `json:"points"`
}

// Place — структура результата
type Place struct {
	ID       int      `json:"id"`
	CityName string   `json:"cityName"`
	Name     string   `json:"name"`
	Geom     string   `json:"geom"`
	Desc     string   `json:"desc"`
	Distance float64  `json:"distance"`
	Photos   []string `json:"photos"`
}

type Photo struct {
	ID          int    `json:"id"`
	PlaceID     int    `json:"placeId"`
	FileLink    string `json:"fileLink"`
	Description string `json:"description"`
}

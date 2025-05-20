package dto

type PlaceResponse struct {
	ID       int     `json:"id"`
	CityName string  `json:"city_name,omitempty"`
	Name     string  `json:"name"`
	Desc     string  `json:"description,omitempty"`
	Lat      float64 `json:"lat"`
	Long     float64 `json:"long"`
	Geom     string  `json:"geom"`
	Distance float64 `json:"distance_meters"`
}

package models

// Country представляет страну
type Country struct {
	ID       int    `json:"id" example:"1"`
	Name     string `json:"name" example:"Kazakhstan"`
	Currency string `json:"currency" example:"USD"`
}

// City представляет город
type City struct {
	ID      int    `json:"id" example:"1"`
	Country string `json:"country" example:"Kazakhstan"`
	Name    string `json:"name" example:"Almaty"`
	Geom    string `json:"geom" example:"POINT(76.8512 43.2220)"`
	Points  int    `json:"points" example:"150"`
}

// Place представляет место/достопримечательность
type Place struct {
	ID       int      `json:"id" example:"1"`
	CityName string   `json:"cityName" example:"Almaty"`
	Name     string   `json:"name" example:"Kok-Tobe Hill"`
	Geom     string   `json:"geom" example:"POINT(76.9572 43.2316)"`
	Desc     string   `json:"desc" example:"Famous hill with panoramic view of Almaty city"`
	Distance float64  `json:"distance" example:"1245.67"`
	Type     string   `json:"type" example:"monument"`
	Price    *int     `json:"price" example:"42"`
	Currency string   `json:"currency" example:"USD"`
	Photos   []string `json:"photos" example:"photo1.jpg,photo2.jpg"`
}

// Photo представляет фотографию места
type Photo struct {
	ID          int    `json:"id" example:"1"`
	PlaceID     int    `json:"placeId" example:"1"`
	FileLink    string `json:"fileLink" example:"https://example.com/photos/place_1_photo_1.jpg"`
	Description string `json:"description" example:"Beautiful sunset view from Kok-Tobe Hill"`
}

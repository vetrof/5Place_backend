package mocks

import (
	"5Place/internal/place/models"
)

type FakeRepository struct{}

func NewFakeRepository() *FakeRepository {
	return &FakeRepository{}
}

func (r *FakeRepository) GetNearPlaces(lat, long float64, limit int, radius float64) ([]models.Place, error) {
	return []models.Place{
		{
			ID:       1,
			CityName: "Астана",
			Name:     "central park",
			Geom:     "POINT(71.419953 51.154506)",
			Desc:     "центральный парк Астаны",
			Distance: 150.25,
			Photos:   []string{"https://astana.citypass.kz/wp-content/uploads/7db97aa358c9dcf7b27cd405bceba5e3.jpeg"},
		},
		{
			ID:       2,
			CityName: "Астана",
			Name:     "Independence Square",
			Geom:     "POINT(71.429745 51.128479)",
			Desc:     "центральная площадь",
			Distance: 300.00,
			Photos:   []string{"https://media-cdn.tripadvisor.com/media/photo-s/0b/89/fb/fc/caption.jpg"},
		},
	}, nil
}

func (r *FakeRepository) GetCountries() ([]models.Country, error) {
	return []models.Country{
		{
			ID:   1,
			Name: "kz",
		},
		{
			ID:   2,
			Name: "uz",
		},
	}, nil
}

func (r *FakeRepository) GetAllCities(country_id int) ([]models.City, error) {
	return []models.City{
		{
			ID:     1,
			Name:   "Астана",
			Geom:   "POINT(71.4304 51.1284)",
			Points: 2,
		},
		{
			ID:     2,
			Name:   "Алматы",
			Geom:   "POINT(76.8860 43.2389)",
			Points: 1,
		},
	}, nil
}

func (r *FakeRepository) GetPhotosByPlaceID(placeID int) ([]string, error) {
	photos := map[int][]string{
		1: {"https://astana.citypass.kz/wp-content/uploads/7db97aa358c9dcf7b27cd405bceba5e3.jpeg"},
		2: {"https://media-cdn.tripadvisor.com/media/photo-s/0b/89/fb/fc/caption.jpg"},
	}
	return photos[placeID], nil
}

// todo
// GetAllCityPlaces выводит все места города
func (db *FakeRepository) GetAllCityPlaces(id int) ([]models.Place, error) {
	return []models.Place{
		{
			ID:       1,
			CityName: "Астана",
			Name:     "central park",
			Geom:     "POINT(71.419953 51.154506)",
			Desc:     "центральный парк Астаны",
			Distance: 150.25,
			Photos:   []string{"https://astana.citypass.kz/wp-content/uploads/7db97aa358c9dcf7b27cd405bceba5e3.jpeg"},
		},
		{
			ID:       2,
			CityName: "Астана",
			Name:     "Independence Square",
			Geom:     "POINT(71.429745 51.128479)",
			Desc:     "центральная площадь",
			Distance: 300.00,
			Photos:   []string{"https://media-cdn.tripadvisor.com/media/photo-s/0b/89/fb/fc/caption.jpg"},
		},
	}, nil
}

// GetAllCityPlaces выводит все места города
func (db *FakeRepository) GetPlaceDetail(id int) (models.Place, error) {
	return models.Place{

		ID:       1,
		CityName: "Астана",
		Name:     "central park",
		Geom:     "POINT(71.419953 51.154506)",
		Desc:     "центральный парк Астаны",
		Distance: 150.25,
		Photos:   []string{"https://astana.citypass.kz/wp-content/uploads/7db97aa358c9dcf7b27cd405bceba5e3.jpeg"},
	}, nil
}

func (db *FakeRepository) GetRandomPlaces(countryId *int64, cityId *int64) ([]models.Place, error) {
	return []models.Place{
		{
			ID:       1,
			CityName: "Астана",
			Name:     "central park",
			Geom:     "POINT(71.419953 51.154506)",
			Desc:     "центральный парк Астаны",
			Distance: 150.25,
			Photos:   []string{"https://astana.citypass.kz/wp-content/uploads/7db97aa358c9dcf7b27cd405bceba5e3.jpeg"},
		},
	}, nil
}

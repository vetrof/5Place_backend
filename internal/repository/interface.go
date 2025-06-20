package repository

import "5Place/internal/models"

type Repository interface {
	GetNearPlaces(lat, long float64, limit int, radius float64) ([]models.Place, error)
	GetAllCities(country_id int) ([]models.City, error)
	GetAllCityPlaces(id int) ([]models.Place, error)
	GetPlaceDetail(id int) ([]models.Place, error)
	GetCountries() ([]models.Country, error)
}

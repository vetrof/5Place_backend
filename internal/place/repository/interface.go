package repository

import (
	"5Place/internal/place/models"
)

type Repository interface {
	GetNearPlaces(lat, long float64, limit int, radius float64) ([]models.Place, error)
	GetAllCities(country_id int) ([]models.City, error)
	GetAllCityPlaces(id int) ([]models.Place, error)
	GetPlaceDetail(id int, lat, long float64) (models.Place, error)
	GetCountries() ([]models.Country, error)
	GetRandomPlaces(countryId *int64, cityId *int64) ([]models.Place, error)
	RepoFavoritesPlaces() ([]models.Place, error)
}

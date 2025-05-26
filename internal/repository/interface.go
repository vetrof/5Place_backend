package repository

import "5Place/internal/models"

type Repository interface {
	GetNearPlaces(lat, long float64) ([]models.Place, error)
	GetAllCities() ([]models.City, error)
}

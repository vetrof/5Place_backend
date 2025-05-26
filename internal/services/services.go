package services

import (
	"5Place/internal/models"
	"5Place/internal/repository"
	"log"
)

// DB — глобальный интерфейсный репозиторий
var DB repository.Repository

// InitServices инициализирует сервисный слой
func InitServices(repo repository.Repository) {
	DB = repo
}

// FindNearbyPlaces находит ближайшие места по координатам
func FindNearbyPlaces(lat, long float64) []models.Place {
	if DB == nil {
		log.Println("Error: repository not initialized")
		return nil
	}

	places, err := DB.GetNearPlaces(lat, long)
	if err != nil {
		log.Printf("Error finding nearby places: %v", err)
		return nil
	}

	return places
}

// GetAllCities возвращает список городов
func GetAllCities() []models.City {
	if DB == nil {
		log.Println("Error: repository not initialized")
		return nil
	}

	cities, err := DB.GetAllCities()
	if err != nil {
		log.Printf("Error finding all cities: %v", err)
		return nil
	}

	return cities
}

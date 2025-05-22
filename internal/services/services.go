package services

import (
	"5Place/internal/models"
	"5Place/internal/repository"
	"log"
)

// DB глобальная переменная для хранения экземпляра репозитория
var DB *repository.PostgresDB

// InitServices инициализирует сервисный слой с репозиторием
func InitServices(repo *repository.PostgresDB) {
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

// FindNearbyPlaces находит ближайшие места по координатам
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

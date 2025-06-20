package services

import (
	"5Place/internal/place/models"
	"5Place/internal/place/repository"
	"log"
)

// DB — глобальный интерфейсный репозиторий
var DB repository.Repository

// InitServices инициализирует сервисный слой
func InitServices(repo repository.Repository) {
	DB = repo
}

// GetAllCities возвращает список стран
func GetCountries() []models.Country {
	if DB == nil {
		log.Println("Error: repository not initialized")
		return nil
	}

	countries, err := DB.GetCountries()
	if err != nil {
		log.Printf("Error finding all cities: %v", err)
		return nil
	}

	return countries
}

// GetAllCities возвращает список городов
func GetAllCities(country_id int) []models.City {
	if DB == nil {
		log.Println("Error: repository not initialized")
		return nil
	}

	cities, err := DB.GetAllCities(country_id)
	if err != nil {
		log.Printf("Error finding all cities: %v", err)
		return nil
	}

	return cities
}

// FindNearbyPlaces находит ближайшие места по координатам
func FindNearbyPlaces(lat, long float64, limit int, radius float64) []models.Place {
	if DB == nil {
		log.Println("Error: repository not initialized")
		return nil
	}

	places, err := DB.GetNearPlaces(lat, long, limit, radius)
	if err != nil {
		log.Printf("Error finding nearby places: %v", err)
		return nil
	}

	return places
}

// Get All Places for city возвращает список мест для города
func CityPlaces(id int) []models.Place {
	if DB == nil {
		log.Println("Error: repository not initialized")
		return nil
	}

	places, err := DB.GetAllCityPlaces(id)
	if err != nil {
		log.Printf("Error finding nearby places: %v", err)
		return nil
	}

	return places
}

// Get All Places for city возвращает список мест для города
func PlaceDetail(id int) []models.Place {
	if DB == nil {
		log.Println("Error: repository not initialized")
		return nil
	}

	places, err := DB.GetPlaceDetail(id)
	if err != nil {
		log.Printf("Error finding nearby places: %v", err)
		return nil
	}

	return places
}

func RandomPlaces(countryId *int64, cityId *int64) []models.Place {
	if DB == nil {
		log.Println("Error: repository not initialized")
		return nil
	}

	if cityId != nil {
		countryId = nil
	}

	places, err := DB.GetRandomPlaces(countryId, cityId)
	if err != nil {
		log.Printf("Error finding nearby places: %v", err)
		return nil
	}

	return places
}

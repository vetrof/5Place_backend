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

// GetCountries возвращает список стран
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
func GetAllCities(countryId int) []models.City {
	if DB == nil {
		log.Println("Error: repository not initialized")
		return nil
	}

	cities, err := DB.GetAllCities(countryId)
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

// CityPlaces Get All Places for city возвращает список мест для города
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

// PlaceDetail Get All Places for city возвращает список мест для города
func PlaceDetail(id int, lat, long float64) *models.Place {
	if DB == nil {
		log.Println("Error: repository not initialized")
		return nil
	}

	place, err := DB.GetPlaceDetail(id, lat, long)
	if err != nil {
		log.Printf("Error finding place: %v", err)
		return nil
	}

	return &place
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

func FavoritePlaces(userId int) []models.Place {
	if DB == nil {
		log.Println("Error: repository not initialized")
		return nil
	}

	places, err := DB.RepoFavoritesPlaces(userId)
	if err != nil {
		log.Printf("Error finding nearby places: %v", err)
		return nil
	}

	return places
}

func AddFavoritePlaces(userId int, placeId int) []models.Place {
	if DB == nil {
		log.Println("Error: repository not initialized")
		return nil
	}

	places, err := DB.RepoAddFavoritesPlaces(userId, placeId)
	if err != nil {
		log.Printf("Error finding nearby places: %v", err)
		return nil
	}

	return places
}

func RepoDeleteFavoritesPlaces(userId int, placeId int) []models.Place {
	if DB == nil {
		log.Println("Error: repository not initialized")
		return nil
	}

	places, err := DB.RepoDeleteFavoritesPlaces(userId, placeId)
	if err != nil {
		log.Printf("Error finding nearby places: %v", err)
		return nil
	}

	return places
}

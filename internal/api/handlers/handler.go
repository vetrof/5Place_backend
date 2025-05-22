package handlers

import (
	"5Place/internal/services"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

// NearPlace принимает реквест, извлекает координаты из кверипараметров,
// передает из в бизнес-логику и принимаем в ответ список ближайших мест
func NearPlace(w http.ResponseWriter, r *http.Request) {

	// берем координаты из квери-параметров
	latStr := r.URL.Query().Get("lat")
	longStr := r.URL.Query().Get("long")

	// из строки в числа
	lat, err := strconv.ParseFloat(latStr, 64)
	if err != nil {
		http.Error(w, "Invalid lat parameter", http.StatusBadRequest)
		return
	}
	long, err := strconv.ParseFloat(longStr, 64)
	if err != nil {
		http.Error(w, "Invalid long parameter", http.StatusBadRequest)
		return
	}

	log.Println("Near place request", "lat=", lat, "long=", long)

	// передаем координаты в сервисный слой и ожидаем список мест
	places := services.FindNearbyPlaces(lat, long)

	// Сериализация и отправка ответа напрямую
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(places); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}

}

func AllCities(w http.ResponseWriter, r *http.Request) {

	// передаем координаты в сервисный слой и ожидаем список мест
	cities := services.GetAllCities()

	// Сериализация и отправка ответа напрямую
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(cities); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}

}

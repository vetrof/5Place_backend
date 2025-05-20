package handlers

import (
	"5Place/internal/dto"
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

	// Преобразуем доменные сущности в DTO
	var resp []dto.PlaceResponse
	for _, p := range places {
		resp = append(resp, dto.PlaceResponse{
			ID:       p.ID,
			Name:     p.Name,
			CityName: p.CityName,
		})
	}

	// Сериализация и отправка ответа
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}

}

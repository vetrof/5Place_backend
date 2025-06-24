package handlers

import (
	"5Place/internal/place/models"
	"5Place/internal/place/services"
	"encoding/json"
	"github.com/go-chi/chi"
	"log"
	"net/http"
	"strconv"
)

type ResponseGeneric[T any, M any] struct {
	Data T `json:"data"`
	Meta M `json:"meta"`
}

type ResponseMeta struct {
	Count        int         `json:"count"`
	Limit        int         `json:"limit"`
	SearchRadius float64     `json:"searchRadius"`
	Center       Coordinates `json:"center"`
}

type Coordinates struct {
	Lat float64 `json:"lat"`
	Lon float64 `json:"lon"`
}

// ErrorResponse представляет структуру ошибки
type ErrorResponse struct {
	Error string `json:"error" example:"Invalid parameter"`
}

// Countries godoc
// @Summary Получить список стран
// @Description Возвращает список всех доступных стран
// @Tags countries
// @Accept json
// @Produce json
// @Success 200 {object} ResponseGeneric[[]models.Country, ResponseMeta]
// @Failure 500 {object} ErrorResponse
// @Router /place/countries [get]
func Countries(w http.ResponseWriter, r *http.Request) {

	countries := services.GetCountries()

	response := ResponseGeneric[[]models.Country, ResponseMeta]{
		Data: countries,
		Meta: ResponseMeta{},
	}

	// Сериализация и отправка ответа напрямую
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}

// Cities godoc
// @Summary Получить города по стране
// @Description Возвращает список городов для указанной страны
// @Tags cities
// @Accept json
// @Produce json
// @Param country_id path int true "ID страны" example(1)
// @Success 200 {object} ResponseGeneric[[]models.City, ResponseMeta]
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /place/cities/country/{country_id} [get]
func Cities(w http.ResponseWriter, r *http.Request) {

	idStr := chi.URLParam(r, "country_id")
	id, err := strconv.Atoi(idStr) // конвертируем в int
	if err != nil {
		http.Error(w, "Неверный ID", http.StatusBadRequest)
		return
	}

	cities := services.GetAllCities(id)

	response := ResponseGeneric[[]models.City, ResponseMeta]{
		Data: cities,
		Meta: ResponseMeta{},
	}

	// Сериализация и отправка ответа напрямую
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}

// NearPlaces godoc
// @Summary Найти места поблизости
// @Description Возвращает места в указанном радиусе от координат
// @Tags places
// @Accept json
// @Produce json
// @Param lat query number true "Широта" example(43.2220)
// @Param long query number true "Долгота" example(76.8512)
// @Param limit query int false "Максимальное количество результатов (1-100)" default(10) example(10)
// @Param radius query number false "Радиус поиска в метрах (1-50000)" default(5000) example(5000)
// @Success 200 {object} ResponseGeneric[[]models.Place, ResponseMeta]
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /place/places/near [get]
func NearPlaces(w http.ResponseWriter, r *http.Request) {

	// берем координаты из квери-параметров
	latStr := r.URL.Query().Get("lat")
	longStr := r.URL.Query().Get("long")

	// из строки в числа
	lat, err := strconv.ParseFloat(latStr, 64)
	if err != nil {
		http.Error(w, "Invalid lat parameter", http.StatusBadRequest)
		return
	}
	lon, err := strconv.ParseFloat(longStr, 64)
	if err != nil {
		http.Error(w, "Invalid long parameter", http.StatusBadRequest)
		return
	}

	limit := 10 // значение по умолчанию
	if limitStr := r.URL.Query().Get("limit"); limitStr != "" {
		if parsedLimit, err := strconv.Atoi(limitStr); err == nil && parsedLimit > 0 && parsedLimit <= 100 {
			limit = parsedLimit
		}
	}

	radius := 5000.0 // 5км по умолчанию
	if radiusStr := r.URL.Query().Get("radius"); radiusStr != "" {
		if parsedRadius, err := strconv.ParseFloat(radiusStr, 64); err == nil && parsedRadius > 0 && parsedRadius <= 50000 {
			radius = parsedRadius
		}
	}

	log.Printf("Near place request: lat=%f, lon=%f, limit=%d, radius=%f", lat, lon, limit, radius)

	// передаем координаты в сервисный слой и ожидаем список мест
	places := services.FindNearbyPlaces(lat, lon, limit, radius)

	response := ResponseGeneric[[]models.Place, ResponseMeta]{
		Data: places,
		Meta: ResponseMeta{
			Count:        len(places),
			Limit:        limit,
			SearchRadius: radius,
			Center: Coordinates{
				Lat: lat,
				Lon: lon,
			},
		},
	}

	// Сериализация и отправка ответа напрямую
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		log.Printf("Failed to encode response: %v", err)
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return // добавить return
	}

}

// RandomPlaces godoc
// @Summary Получить случайные места
// @Description Возвращает случайные места с возможностью фильтрации по стране или городу
// @Tags places
// @Accept json
// @Produce json
// @Param country query int false "ID страны для фильтрации" example(1)
// @Param city query int false "ID города для фильтрации" example(1)
// @Success 200 {object} ResponseGeneric[[]models.Place, ResponseMeta]
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /place/places/random [get]
func RandomPlaces(w http.ResponseWriter, r *http.Request) {

	var countryId *int64
	var cityId *int64

	countryIdStr := r.URL.Query().Get("country")
	if countryIdStr != "" {
		val, err := strconv.ParseInt(countryIdStr, 10, 64)
		if err != nil {
			http.Error(w, "Invalid country parameter", http.StatusBadRequest)
			return
		}
		countryId = &val
	}

	cityIdStr := r.URL.Query().Get("city")
	if cityIdStr != "" {
		val, err := strconv.ParseInt(cityIdStr, 10, 64)
		if err != nil {
			http.Error(w, "Invalid city parameter", http.StatusBadRequest)
			return
		}
		cityId = &val
	}

	// передаем координаты в сервисный слой и ожидаем список мест
	places := services.RandomPlaces(countryId, cityId)

	response := ResponseGeneric[[]models.Place, ResponseMeta]{
		Data: places,
	}

	// Сериализация и отправка ответа напрямую
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		log.Printf("Failed to encode response: %v", err)
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return // добавить return
	}

}

// PlaceDetail godoc
// @Summary Получить детали места
// @Description Возвращает подробную информацию о месте по его ID
// @Tags places
// @Accept json
// @Produce json
// @Param place_id path int true "ID места" example(1)
// @Success 200 {object} ResponseGeneric[[]models.Place, ResponseMeta]
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /place/places/{place_id} [get]
func PlaceDetail(w http.ResponseWriter, r *http.Request) {

	idStr := chi.URLParam(r, "place_id")
	id, err := strconv.Atoi(idStr) // конвертируем в int
	if err != nil {
		http.Error(w, "Неверный ID", http.StatusBadRequest)
		return
	}

	// передаем координаты в сервисный слой и ожидаем список мест
	cityPlaces := services.PlaceDetail(id)

	response := ResponseGeneric[[]models.Place, ResponseMeta]{
		Data: cityPlaces,
		Meta: ResponseMeta{},
	}

	// Сериализация и отправка ответа напрямую
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}

}

// CityPlaces godoc
// @Summary Получить места в городе
// @Description Возвращает список всех мест в указанном городе
// @Tags places
// @Accept json
// @Produce json
// @Param city_id path int true "ID города" example(1)
// @Success 200 {object} ResponseGeneric[[]models.Place, ResponseMeta]
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /place/places/city/{city_id} [get]
func CityPlaces(w http.ResponseWriter, r *http.Request) {

	idStr := chi.URLParam(r, "city_id")
	id, err := strconv.Atoi(idStr) // конвертируем в int
	if err != nil {
		http.Error(w, "Неверный ID", http.StatusBadRequest)
		return
	}

	// передаем координаты в сервисный слой и ожидаем список мест
	cityPlaces := services.CityPlaces(id)

	response := ResponseGeneric[[]models.Place, ResponseMeta]{
		Data: cityPlaces,
		Meta: ResponseMeta{},
	}

	// Сериализация и отправка ответа напрямую
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}

}

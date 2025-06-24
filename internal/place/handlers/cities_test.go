package handlers_test

import (
	"5Place/internal/place/handlers"
	"5Place/internal/place/models"
	"5Place/internal/place/services"
	"context"
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"net/http"
	"net/http/httptest"
	"testing"
)

type fakeRepo struct{}

func (f *fakeRepo) GetAllCities(countryID int) ([]models.City, error) {
	return []models.City{
		{ID: 1, Name: "Almaty"},
		{ID: 2, Name: "Astana"},
	}, nil
}

func (f *fakeRepo) GetCountries() ([]models.Country, error) { return nil, nil }
func (f *fakeRepo) GetNearPlaces(float64, float64, int, float64) ([]models.Place, error) {
	return nil, nil
}
func (f *fakeRepo) GetAllCityPlaces(int) ([]models.Place, error)           { return nil, nil }
func (f *fakeRepo) GetPlaceDetail(int) ([]models.Place, error)             { return nil, nil }
func (f *fakeRepo) GetRandomPlaces(*int64, *int64) ([]models.Place, error) { return nil, nil }

func TestCities(t *testing.T) {
	services.InitServices(&fakeRepo{})

	req := httptest.NewRequest("GET", "/place/cities/country/1", nil)
	w := httptest.NewRecorder()

	// Вставляем RouteContext вручную, как это делает chi в реальном роутере
	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("country_id", "1")
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

	handlers.Cities(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("Ожидался статус 200, но был %d", w.Code)
	}

	var result handlers.ResponseGeneric[[]models.City, handlers.ResponseMeta]
	if err := json.NewDecoder(w.Body).Decode(&result); err != nil {
		t.Fatalf("Ошибка при декодировании JSON: %v", err)
	}

	if len(result.Data) != 2 {
		t.Errorf("Ожидалось 2 города, но получили %d", len(result.Data))
	}
}

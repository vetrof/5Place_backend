package routers

import (
	"5Place/internal/handlers"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func Router() chi.Router {

	// Router init
	router := chi.NewRouter()
	router.Use(middleware.Logger)

	// Public paths
	//router.Get("/countries", handlers.Countries)
	router.Get("/countries", handlers.Countries)
	router.Get("/cities/country/{country_id}", handlers.Cities)
	router.Get("/places/near", handlers.NearPlace)
	router.Get("/places/{place_id}", handlers.PlaceDetail)
	router.Get("/places/city/{city_id}", handlers.CityPlaces)

	// TODO
	//router.Get("/cities/{city_id}", handlers.CityDetail) // детали города
	//router.Get("/places/search", handlers.SearchPlaces) // поиск мест

	return router
}

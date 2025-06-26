package router

import (
	"5Place/internal/place/handlers"
	"github.com/go-chi/chi/v5"
)

func Router() chi.Router {

	// Router init
	router := chi.NewRouter()

	// Public paths
	router.Get("/countries", handlers.Countries)
	router.Get("/cities/country/{country_id}", handlers.Cities)
	router.Get("/near", handlers.NearPlaces)
	router.Get("/random", handlers.RandomPlaces)
	router.Get("/detail/{place_id}", handlers.PlaceDetail)
	router.Get("/city/{city_id}", handlers.CityPlaces)
	router.Get("/favorite", handlers.FavoritePlaces)
	router.Post("/favorite/{place_id}", handlers.FavoritePlaces)

	// TODO
	//router.Get("/cities/{city_id}", handlers.CityDetail) // детали города
	//router.Get("/places/search", handlers.SearchPlaces) // поиск мест

	return router
}

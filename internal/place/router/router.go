package router

import (
	"5Place/internal/auth"
	"5Place/internal/place/handlers"
	"github.com/go-chi/chi/v5"
)

func Router(jwtConfig *auth.JWTConfig) chi.Router {

	// Router init
	router := chi.NewRouter()

	// Public paths
	router.Get("/countries", handlers.Countries)
	router.Get("/cities/country/{country_id}", handlers.Cities)
	router.Get("/near", handlers.NearPlaces)
	router.Get("/random", handlers.RandomPlaces)
	router.Get("/detail/{place_id}", handlers.PlaceDetail)
	router.Get("/city/{city_id}", handlers.CityPlaces)

	// Защищённые пути:
	router.Group(func(r chi.Router) {
		r.Use(auth.JWTMiddleware(jwtConfig))
		r.Get("/favorite", handlers.FavoritePlaces)
		r.Post("/favorite/{place_id}", handlers.FavoritePlaces)
		r.Delete("/favorite/{place_id}", handlers.FavoritePlaces)
	})

	return router
}

package router

import (
	"5Place/internal/api/handlers"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func Router() chi.Router {

	// Router init
	router := chi.NewRouter()
	router.Use(middleware.Logger)

	// Public paths
	router.Get("/near_place", handlers.NearPlace)

	return router
}

package auth

import (
	"database/sql"

	"github.com/go-chi/chi/v5"
)

// SetupAuthRoutes настраивает роуты для авторизации
func SetupAuthRoutes(r chi.Router, db *sql.DB, jwtConfig *JWTConfig) {
	// Создаем репозиторий, сервис и handler
	repo := NewRepository(db)
	service := NewService(repo, jwtConfig)
	handler := NewHandler(service)

	// Публичные роуты (без авторизации)
	r.Route("/auth", func(r chi.Router) {
		r.Post("/register", handler.Register)
		r.Post("/login", handler.Login)

		// Защищенные роуты (с авторизацией)
		r.Group(func(r chi.Router) {
			r.Use(JWTMiddleware(jwtConfig))
			r.Get("/profile", handler.Profile)
		})
	})
}

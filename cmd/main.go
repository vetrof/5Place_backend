// @title 5Place API
// @version 1.0
// @description API для приложения 5Place - поиск и управление местами
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /

// @schemes http https

// @securityDefinitions.apikey Bearer
// @in header
// @name Authorization
// @description Type "Bearer" followed by a space and JWT token.

// @tag.name auth
// @tag.description Операции авторизации и аутентификации

// @tag.name countries
// @tag.description Операции со странами

// @tag.name cities
// @tag.description Операции с городами

// @tag.name places
// @tag.description Операции с местами и достопримечательностями

package main

import (
	"5Place/internal/auth"
	"5Place/internal/config/utils"
	repository2 "5Place/internal/place/repository"
	"5Place/internal/place/repository/mocks"
	placeRouter "5Place/internal/place/router"
	"5Place/internal/place/services"
	userRouter "5Place/internal/user/router"
	"database/sql"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"
	httpSwagger "github.com/swaggo/http-swagger"
	"log"
	"net/http"
	"os"

	_ "5Place/docs" // импорт сгенерированных swagger файлов
)

func main() {
	// подгружаем переменные из .env
	err := godotenv.Load(".env")
	if err != nil {
		log.Printf("Warning: Error loading .env file: %v", err)
	}

	// инициализация JWT конфигурации
	jwtConfig := &auth.JWTConfig{
		SecretKey:       utils.GetEnvOrDefault("JWT_SECRET", "your-super-secret-jwt-key-change-in-production"),
		ExpirationHours: utils.GetEnvIntOrDefault("JWT_EXPIRATION_HOURS", 24),
	}

	// инициализация репозитория
	var repo repository2.Repository // интерфейс, который реализуют и PostgresDB, и FakeRepository
	var db *sql.DB                  // для передачи в auth роуты

	// проверяем переменную окружения REPO
	// если она равна "fake", то используем мок репозиторий, иначе - реальный
	r := os.Getenv("REPO")
	if r == "fake" {
		repo = mocks.NewFakeRepository()
		log.Println("Fake repository initialized")
		// Для fake repo создаем пустое подключение к БД (JWT auth будет работать только с реальной БД)
		log.Println("Warning: JWT auth will not work with fake repository")
	} else {
		pgRepo, err := repository2.NewPostgresDB()
		if err != nil {
			log.Fatalf("Failed to initialize repository: %v", err)
		}
		defer pgRepo.Close()
		log.Println("Postgres repository initialized")
		repo = pgRepo

		// Получаем подключение к БД для auth
		db = pgRepo.GetDB()
	}

	// инициализация сервисного слоя
	services.InitServices(repo)
	log.Println("Services initialized successfully")

	// создаем chi router
	router := chi.NewRouter()

	// базовые middleware
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(middleware.RequestID)

	// настройка CORS для фронтенда
	router.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

			if r.Method == "OPTIONS" {
				w.WriteHeader(http.StatusOK)
				return
			}

			next.ServeHTTP(w, r)
		})
	})

	// настройка auth роутов (только если есть подключение к БД)
	// добавляем роут для авторизации
	if db != nil {
		auth.SetupAuthRoutes(router, db, jwtConfig)
		log.Println("Auth routes initialized")
	}

	// роуты по доменам
	router.Route("/api/v1", func(r chi.Router) {
		r.Mount("/place", placeRouter.Router())
		r.Mount("/user", userRouter.Router())
	})

	// Swagger
	router.Handle("/swagger/*", httpSwagger.WrapHandler)

	// Health check endpoint
	router.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":"OK","message":"Server is running"}`))
	})

	// API info endpoint
	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message":"5Place API","version":"1.0","swagger":"/swagger/"}`))
	})

	port := utils.GetEnvOrDefault("PORT", "5555")
	log.Printf("Starting server at port %s", port)
	log.Printf("Swagger documentation available at: http://localhost:%s/swagger/", port)
	log.Printf("Health check available at: http://localhost:%s/health", port)
	if db != nil {
		log.Printf("Auth endpoints available at: http://localhost:%s/auth/", port)
	}

	// Server init
	if err := http.ListenAndServe(":"+port, router); err != nil {
		log.Fatal("Error starting server:", err)
	}
}

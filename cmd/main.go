// @title 5Place API
// @version 1.0
// @description API для приложения 5Place - поиск и управление местами
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:5555
// @BasePath /

// @schemes http https

// @tag.name countries
// @tag.description Операции со странами

// @tag.name cities
// @tag.description Операции с городами

// @tag.name places
// @tag.description Операции с местами и достопримечательностями

package main

import (
	repository2 "5Place/internal/place/repository"
	"5Place/internal/place/repository/mocks"
	placeRouter "5Place/internal/place/router"
	userRouter "5Place/internal/place/router"
	"5Place/internal/place/services"
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

	// инициализация репозитория
	var repo repository2.Repository // интерфейс, который реализуют и PostgresDB, и FakeRepository

	// проверяем переменную окружения REPO
	// если она равна "fake", то используем мок репозиторий, иначе - реальный
	r := os.Getenv("REPO")
	if r == "fake" {
		repo = mocks.NewFakeRepository()
		log.Println("Fake repository initialized")
	} else {
		pgRepo, err := repository2.NewPostgresDB()
		if err != nil {
			log.Fatalf("Failed to initialize repository: %v", err)
		}
		defer pgRepo.Close()
		log.Println("Postgres repository initialized")
		repo = pgRepo
	}

	// инициализация сервисного слоя и репозитория
	services.InitServices(repo)
	log.Println("Services initialized successfully")

	// объединяем роутеры
	mux := http.NewServeMux()
	mux.Handle("/place/", http.StripPrefix("/place", placeRouter.Router()))
	mux.Handle("/user/", http.StripPrefix("/user", userRouter.Router()))

	// Swagger endpoint
	mux.Handle("/swagger/", httpSwagger.WrapHandler)

	// Health check endpoint
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	// Server init
	port := os.Getenv("PORT")
	if port == "" {
		port = "5555"
	}
	log.Printf("Starting server at port %s", port)
	log.Printf("Swagger documentation available at: http://localhost:%s/swagger/", port)
	if err := http.ListenAndServe(":"+port, mux); err != nil {
		log.Fatal("Error starting server:", err)
	}
}

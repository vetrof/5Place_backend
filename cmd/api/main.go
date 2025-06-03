package main

import (
	"5Place/internal/repository"
	"5Place/internal/repository/mocks"
	"5Place/internal/services"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"

	"5Place/internal/api/routers"
)

func main() {
	// подгружаем переменные из .env
	err := godotenv.Load(".env")
	if err != nil {
		log.Printf("Warning: Error loading .env file: %v", err)
	}

	// инициализация репозитория
	var repo repository.Repository // интерфейс, который реализуют и PostgresDB, и FakeRepository

	// проверяем переменную окружения REPO
	// если она равна "fake", то используем мок репозиторий, иначе - реальный
	r := os.Getenv("REPO")
	if r == "fake" {
		repo = mocks.NewFakeRepository()
		log.Println("Fake repository initialized")
	} else {
		pgRepo, err := repository.NewPostgresDB()
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

	// get rout
	router := routers.Router()

	// Server init
	port := os.Getenv("PORT")
	if port == "" {
		port = "5555"
	}
	log.Println("Starting server at port", port)
	if err := http.ListenAndServe(":"+port, router); err != nil {
		log.Fatal("Error starting server:", err)
	}
}

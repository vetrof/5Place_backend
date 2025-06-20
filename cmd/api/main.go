package main

import (
	repository2 "5Place/internal/place/repository"
	"5Place/internal/place/repository/mocks"
	"5Place/internal/place/services"
	"5Place/internal/routers"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
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

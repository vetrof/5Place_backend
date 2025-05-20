package main

import (
	"5Place/internal/repository"
	"5Place/internal/services"
	"log"
	"net/http"
	"os"

	"5Place/internal/api/routers"
)

func main() {

	// db init
	repo, err := repository.NewPostgresDB()
	if err != nil {
		log.Fatalf("Failed to initialize repository: %v", err)
	}
	defer repo.Close()
	log.Println("Repository initialized successfully")

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

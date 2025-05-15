package main

import (
	"5Place/internal/repository"
	"log"
)

func main() {

	// db init
	// Инициализация репозитория
	repo, err := interfaces.NewPostgresDB()
	if err != nil {
		log.Fatalf("Failed to initialize repository: %v", err)
	}
	defer repo.Close()

	log.Println("Repository initialized successfully")

	//// Server init
	//r := router.NewRouter()

	//port := os.Getenv("PORT")
	//if port == "" {
	//	port = "5555"
	//}
	//
	//log.Println("Starting server at port", port)
	//if err := http.ListenAndServe(":"+port, r); err != nil {
	//	log.Fatal("Error starting server:", err)
	//}
}

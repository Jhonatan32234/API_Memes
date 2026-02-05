package main

import (
	"log"
	"net/http"
	"os"

	"api_memes/internal/api/routes"
	"api_memes/internal/shared"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
        log.Println("No .env file found")
    }
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}
	db, err := shared.NewPostgres()
	if err != nil {
		log.Fatal(err)
	}

	r := chi.NewRouter()
	routes.RegisterRoutes(r, db)

	log.Println("API running on port " + port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}

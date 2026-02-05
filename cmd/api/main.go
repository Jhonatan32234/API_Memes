package main

import (
	"log"
	"net/http"

	"estructura_base/internal/api/routes"
	"estructura_base/internal/shared"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"

)

func main() {
	if err := godotenv.Load(); err != nil {
        log.Println("No .env file found")
    }
	db, err := shared.NewPostgres()
	if err != nil {
		log.Fatal(err)
	}

	r := chi.NewRouter()
	routes.RegisterRoutes(r, db)

	log.Println("API running on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}

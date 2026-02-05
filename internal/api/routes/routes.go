package routes

import (
	"database/sql"

	"estructura_base/internal/api/handlers"
	"estructura_base/internal/users"

	"github.com/go-chi/chi/v5"
)

func RegisterRoutes(r chi.Router, db *sql.DB) {
	userService := users.NewService(db)
	userHandler := handlers.NewUserHandler(userService)

	r.Post("/users", userHandler.Create)
	r.Get("/users", userHandler.GetAll)
    r.Get("/users/{id}", userHandler.GetByID)
}

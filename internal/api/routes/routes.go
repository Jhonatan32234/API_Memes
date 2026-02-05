package routes

import (
	"database/sql"

	"api_memes/internal/api/handlers"
	"api_memes/internal/memes"
	"api_memes/internal/users"

	"github.com/go-chi/chi/v5"
)

func RegisterRoutes(r chi.Router, db *sql.DB) {
    userService := users.NewService(db)
    userHandler := handlers.NewUserHandler(userService)
    memeService := memes.NewService(db)
    memeHandler := handlers.NewMemeHandler(memeService)

    r.Post("/register", userHandler.Create)
    r.Post("/login", userHandler.Login)
    r.Get("/memes", memeHandler.GetAll)

    r.Group(func(r chi.Router) {
        r.Use(AuthMiddleware)
        r.Post("/memes", memeHandler.Create)
    })
}

package handlers

import (
	"api_memes/internal/shared"
	"api_memes/internal/users"
	"encoding/json"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/golang-jwt/jwt/v5"
)

type UserHandler struct {
	service *users.Service
}

func NewUserHandler(service *users.Service) *UserHandler {
	return &UserHandler{service: service}
}

func (h *UserHandler) Login(w http.ResponseWriter, r *http.Request) {
    var dto users.CreateUserDTO // Reutilizamos el DTO de Email/Password
    if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
        http.Error(w, "invalid json", http.StatusBadRequest)
        return
    }

    user, err := h.service.Login(dto.Email, dto.Password)
    if err != nil {
        http.Error(w, "unauthorized", http.StatusUnauthorized)
        return
    }

    // Crear el Token
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
        "user_id": user.ID,
        "exp":     time.Now().Add(time.Hour * 72).Unix(),
    })

    tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
    if err != nil {
        http.Error(w, "error generating token", http.StatusInternalServerError)
        return
    }

    json.NewEncoder(w).Encode(map[string]string{"token": tokenString})
}

func (h *UserHandler) Create(w http.ResponseWriter, r *http.Request) {
	var dto users.CreateUserDTO
	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
		http.Error(w, "invalid json", http.StatusBadRequest)
		return
	}
	user, err := h.service.Create(dto)
	if err != nil {
		switch err {
		case shared.ErrValidation:
			http.Error(w, err.Error(), http.StatusBadRequest)
		case shared.ErrDuplicate:
			http.Error(w, err.Error(), http.StatusConflict)
		default:
			http.Error(w, "internal error", http.StatusInternalServerError)
		}
		return
	}
	response := users.UserResponse{
	ID:        user.ID,
	Email:     user.Email,
	CreatedAt: user.CreatedAt,
    }
    
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(response)

}

func (h *UserHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	usersList, err := h.service.GetAll()
	if err != nil {
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}
    var response []users.UserResponse
    for _, u := range usersList {
    	response = append(response, users.UserResponse{
    		ID:        u.ID,
    		Email:     u.Email,
    		CreatedAt: u.CreatedAt,
    	})
    }
    
    json.NewEncoder(w).Encode(response)

}

func (h *UserHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	user, err := h.service.GetByID(id)
	if err != nil {
		http.Error(w, "not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(users.UserResponse{
	ID:        user.ID,
	Email:     user.Email,
	CreatedAt: user.CreatedAt,
    })

}

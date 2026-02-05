package handlers

import (
	"encoding/json"
	"estructura_base/internal/shared"
	"estructura_base/internal/users"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type UserHandler struct {
	service *users.Service
}

func NewUserHandler(service *users.Service) *UserHandler {
	return &UserHandler{service: service}
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

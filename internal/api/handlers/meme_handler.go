package handlers

import (
	"api_memes/internal/memes"
	"encoding/json"
	"net/http"
)

type MemeHandler struct {
    service *memes.Service
}

func NewMemeHandler(s *memes.Service) *MemeHandler {
    return &MemeHandler{service: s}
}

func (h *MemeHandler) GetAll(w http.ResponseWriter, r *http.Request) {
    list, err := h.service.GetAll()
    if err != nil {
        http.Error(w, "Error al obtener memes", http.StatusInternalServerError)
        return
    }
    
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(list)
}

func (h *MemeHandler) Create(w http.ResponseWriter, r *http.Request) {
    userID, ok := r.Context().Value("user_id").(float64)
    if !ok {
        http.Error(w, "No autorizado", http.StatusUnauthorized)
        return
    }

    var dto memes.CreateMemeDTO
    if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
        http.Error(w, "JSON inválido", http.StatusBadRequest)
        return
    }

    dto.AuthorID = int64(userID)
    
    meme, err := h.service.Create(dto)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(meme)
}
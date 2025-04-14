package handlers

import (
	"VoteGolang/internals/data"
	"VoteGolang/internals/usecases"
	"VoteGolang/pkg/domain"
	"encoding/json"
	"net/http"
	"strconv"
)

type GeneralNewsHandler struct {
	UseCase      *usecases.GeneralNewsUseCase
	TokenManager *domain.JwtToken
}

func NewGeneralNewsHandler(uc *usecases.GeneralNewsUseCase, tokenManager *domain.JwtToken) *GeneralNewsHandler {
	return &GeneralNewsHandler{
		UseCase:      uc,
		TokenManager: tokenManager,
	}
}

func (h *GeneralNewsHandler) Create(w http.ResponseWriter, r *http.Request) {
	var news data.GeneralNews
	if err := json.NewDecoder(r.Body).Decode(&news); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}
	if err := h.UseCase.Create(&news); err != nil {
		http.Error(w, "Failed to create news", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(news)
}

func (h *GeneralNewsHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	newsList, err := h.UseCase.GetAll()
	if err != nil {
		http.Error(w, "Failed to fetch news", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(newsList)
}

func (h *GeneralNewsHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	idParam := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	news, err := h.UseCase.GetByID(uint(id))
	if err != nil {
		http.Error(w, "News not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(news)
}

func (h *GeneralNewsHandler) Update(w http.ResponseWriter, r *http.Request) {
	var news data.GeneralNews
	if err := json.NewDecoder(r.Body).Decode(&news); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}
	if err := h.UseCase.Update(&news); err != nil {
		http.Error(w, "Failed to update news", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(news)
}

func (h *GeneralNewsHandler) Delete(w http.ResponseWriter, r *http.Request) {
	idParam := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}
	if err := h.UseCase.Delete(uint(id)); err != nil {
		http.Error(w, "Failed to delete news", http.StatusInternalServerError)
		return
	}
	w.Write([]byte("News deleted successfully"))
}

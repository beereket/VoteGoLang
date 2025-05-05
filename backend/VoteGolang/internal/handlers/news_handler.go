package handlers

import (
	"VoteGolang/internal/domain"
	"VoteGolang/internal/service"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type NewsHandler struct {
	Service *service.NewsService
}

func NewNewsHandler(s *service.NewsService) *NewsHandler {
	return &NewsHandler{Service: s}
}

func (h *NewsHandler) List(w http.ResponseWriter, r *http.Request) {
	pageStr := r.URL.Query().Get("page")
	limitStr := r.URL.Query().Get("limit")

	page, limit := 1, 10
	if p, err := strconv.Atoi(pageStr); err == nil && p > 0 {
		page = p
	}
	if l, err := strconv.Atoi(limitStr); err == nil && l > 0 {
		limit = l
	}

	news, err := h.Service.List(page, limit)
	if err != nil {
		http.Error(w, "❌ Failed to fetch news", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(news)
}

func (h *NewsHandler) Create(w http.ResponseWriter, r *http.Request) {
	var input domain.CreateNewsInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "❌ Invalid input", http.StatusBadRequest)
		return
	}
	if err := h.Service.Create(input); err != nil {
		http.Error(w, "❌ Failed to create news", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(map[string]string{"message": "✅ News article created!"})
}

func (h *NewsHandler) Delete(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "❌ Invalid news ID", http.StatusBadRequest)
		return
	}
	if err := h.Service.Delete(id); err != nil {
		http.Error(w, "❌ Failed to delete news", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(map[string]string{"message": "✅ News deleted successfully"})
}

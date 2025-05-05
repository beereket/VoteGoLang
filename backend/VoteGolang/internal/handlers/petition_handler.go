package handlers

import (
	"VoteGolang/internal/domain"
	"VoteGolang/internal/service"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type PetitionHandler struct {
	Service *service.PetitionService
}

func NewPetitionHandler(s *service.PetitionService) *PetitionHandler {
	return &PetitionHandler{Service: s}
}

func (h *PetitionHandler) List(w http.ResponseWriter, r *http.Request) {
	petitions, err := h.Service.List()
	if err != nil {
		http.Error(w, "❌ Failed to fetch petitions", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(petitions)
}

func (h *PetitionHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		http.Error(w, "❌ Invalid petition ID", http.StatusBadRequest)
		return
	}
	petition, err := h.Service.GetByID(id)
	if err != nil {
		http.Error(w, "❌ Petition not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(petition)
}

func (h *PetitionHandler) Create(w http.ResponseWriter, r *http.Request) {
	var req domain.Petition
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "❌ Invalid input", http.StatusBadRequest)
		return
	}
	req.UserID = 1
	req.Status = "open"

	err := h.Service.Create(req)
	if err != nil {
		http.Error(w, "❌ Failed to create petition", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(map[string]string{"message": "✅ Petition created successfully"})
}

func (h *PetitionHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		http.Error(w, "❌ Invalid petition ID", http.StatusBadRequest)
		return
	}
	if err := h.Service.Delete(id); err != nil {
		http.Error(w, "❌ Failed to delete petition", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(map[string]string{"message": "✅ Petition deleted successfully"})
}

func (h *PetitionHandler) Vote(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value("user_id").(int)
	if !ok {
		http.Error(w, "❌ Unauthorized", http.StatusUnauthorized)
		return
	}

	var req domain.PetitionVoteRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "❌ Invalid request body", http.StatusBadRequest)
		return
	}

	if err := h.Service.Vote(userID, req); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"message": "✅ Vote recorded successfully"})
}

func (h *PetitionHandler) GetResults(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		http.Error(w, "❌ Invalid petition ID", http.StatusBadRequest)
		return
	}

	results, err := h.Service.GetResults(id)
	if err != nil {
		http.Error(w, "❌ Failed to fetch petition results", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(results)
}

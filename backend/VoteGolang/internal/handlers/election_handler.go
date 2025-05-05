package handlers

import (
	"VoteGolang/internal/domain"
	"VoteGolang/internal/service"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type ElectionHandler struct {
	Service *service.ElectionService
}

func NewElectionHandler(s *service.ElectionService) *ElectionHandler {
	return &ElectionHandler{Service: s}
}

func (h *ElectionHandler) List(w http.ResponseWriter, r *http.Request) {
	elections, err := h.Service.List()
	if err != nil {
		http.Error(w, "❌ Failed to fetch elections", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(elections)
}

func (h *ElectionHandler) Get(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		http.Error(w, "❌ Invalid election ID", http.StatusBadRequest)
		return
	}
	election, err := h.Service.Get(id)
	if err != nil {
		http.Error(w, "❌ Election not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(election)
}

func (h *ElectionHandler) Create(w http.ResponseWriter, r *http.Request) {
	var input domain.Election
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "❌ Invalid input", http.StatusBadRequest)
		return
	}
	if err := h.Service.Create(input); err != nil {
		http.Error(w, "❌ Failed to create election", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(map[string]string{"message": "✅ Election created successfully"})
}

func (h *ElectionHandler) Update(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		http.Error(w, "❌ Invalid election ID", http.StatusBadRequest)
		return
	}
	var input domain.Election
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "❌ Invalid input", http.StatusBadRequest)
		return
	}
	if err := h.Service.Update(id, input); err != nil {
		http.Error(w, "❌ Failed to update election", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(map[string]string{"message": "✅ Election updated successfully"})
}

func (h *ElectionHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		http.Error(w, "❌ Invalid election ID", http.StatusBadRequest)
		return
	}
	if err := h.Service.Delete(id); err != nil {
		http.Error(w, "❌ Failed to delete election", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(map[string]string{"message": "✅ Election deleted successfully"})
}

func (h *ElectionHandler) GetWithCandidates(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["election_id"])
	if err != nil {
		http.Error(w, "❌ Invalid election ID", http.StatusBadRequest)
		return
	}

	data, err := h.Service.GetWithCandidates(id)
	if err != nil {
		http.Error(w, "❌ Election not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(data)
}

func (h *ElectionHandler) GetResults(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		http.Error(w, "❌ Invalid election ID", http.StatusBadRequest)
		return
	}

	results, err := h.Service.GetResults(id)
	if err != nil {
		http.Error(w, "❌ Failed to fetch results", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(results)
}

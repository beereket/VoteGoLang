package handlers

import (
	"VoteGolang/internal/domain"
	"VoteGolang/internal/service"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type CandidateHandler struct {
	Service *service.CandidateService
}

func NewCandidateHandler(s *service.CandidateService) *CandidateHandler {
	return &CandidateHandler{Service: s}
}

func (h *CandidateHandler) ListByElection(w http.ResponseWriter, r *http.Request) {
	electionID := mux.Vars(r)["election_id"]
	candidates, err := h.Service.ListByElection(electionID)
	if err != nil {
		http.Error(w, "❌ Failed to list candidates", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(candidates)
}

func (h *CandidateHandler) Create(w http.ResponseWriter, r *http.Request) {
	var input domain.Candidate
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "❌ Invalid input", http.StatusBadRequest)
		return
	}
	if err := h.Service.Create(input); err != nil {
		http.Error(w, "❌ Failed to create candidate", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(map[string]string{"message": "✅ Candidate created"})
}

func (h *CandidateHandler) Update(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		http.Error(w, "❌ Invalid ID", http.StatusBadRequest)
		return
	}
	var input domain.Candidate
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "❌ Invalid input", http.StatusBadRequest)
		return
	}
	if err := h.Service.Update(id, input); err != nil {
		http.Error(w, "❌ Failed to update", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(map[string]string{"message": "✅ Candidate updated"})
}

func (h *CandidateHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		http.Error(w, "❌ Invalid ID", http.StatusBadRequest)
		return
	}
	if err := h.Service.Delete(id); err != nil {
		http.Error(w, "❌ Failed to delete", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(map[string]string{"message": "✅ Deleted"})
}

func (h *CandidateHandler) Vote(w http.ResponseWriter, r *http.Request) {
	uid, ok := r.Context().Value("user_id").(int)
	if !ok {
		http.Error(w, "❌ Unauthorized", http.StatusUnauthorized)
		return
	}
	var req domain.VoteRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "❌ Invalid body", http.StatusBadRequest)
		return
	}
	if err := h.Service.CastVote(uid, req); err != nil {
		http.Error(w, err.Error(), http.StatusForbidden)
		return
	}
	json.NewEncoder(w).Encode(map[string]string{"message": "✅ Vote cast"})
}

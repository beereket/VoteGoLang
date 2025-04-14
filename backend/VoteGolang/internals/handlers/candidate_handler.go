package handlers

import (
	"VoteGolang/internals/services/auth"
	"VoteGolang/internals/usecases"
	"VoteGolang/pkg/domain"
	"encoding/json"
	"github.com/dgrijalva/jwt-go"
	"log"
	"net/http"
)

type CandidateHandler struct {
	UseCase      *usecases.CandidateUseCase
	TokenManager *domain.JwtToken
}

func NewCandidateHandler(uc *usecases.CandidateUseCase, tokenManager *domain.JwtToken) *CandidateHandler {
	return &CandidateHandler{
		UseCase:      uc,
		TokenManager: tokenManager,
	}
}

func (h *CandidateHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	typ := r.URL.Query().Get("type")
	if typ == "" {
		http.Error(w, "type is required", http.StatusBadRequest)
		return
	}

	candidates, err := h.UseCase.GetAllByType(typ)
	if err != nil {
		http.Error(w, "failed to get candidates", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(candidates)
}

func (h *CandidateHandler) Vote(w http.ResponseWriter, r *http.Request) {
	token, err := auth.ExtractTokenFromRequest(r)
	if err != nil {
		http.Error(w, "Authorization token missing", http.StatusUnauthorized)
		return
	}

	payload := &domain.JwtClaims{}
	_, err = jwt.ParseWithClaims(token, payload, func(t *jwt.Token) (interface{}, error) {
		return h.TokenManager.Secret, nil
	})
	if err != nil {
		http.Error(w, "Invalid token", http.StatusUnauthorized)
		return
	}

	userID := payload.UserID
	if userID == 0 {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	if userID == 0 {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	log.Printf("User ID from token: %s", userID)

	type voteRequest struct {
		CandidateID   uint   `json:"candidate_id"`
		CandidateType string `json:"candidate_type"`
	}

	var req voteRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request format", http.StatusBadRequest)
		return
	}

	err = h.UseCase.Vote(req.CandidateID, userID, req.CandidateType)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Write([]byte("Vote successful"))
}

package handlers

import (
	"VoteGolang/internals/data"
	"VoteGolang/internals/services/auth"
	"VoteGolang/internals/usecases"
	"VoteGolang/pkg/domain"
	"encoding/json"
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"strconv"
)

type PetitionHandler struct {
	usecase      usecases.PetitionUseCase
	TokenManager *domain.JwtToken
}

func NewPetitionHandler(usecase usecases.PetitionUseCase, tokenManager *domain.JwtToken) *PetitionHandler {
	return &PetitionHandler{
		usecase:      usecase,
		TokenManager: tokenManager,
	}
}

func (h *PetitionHandler) CreatePetition(w http.ResponseWriter, r *http.Request) {
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

	var p data.Petition
	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	p.UserID = userID

	if err := h.usecase.CreatePetition(&p); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(p)
}

func (h *PetitionHandler) GetAllPetitions(w http.ResponseWriter, r *http.Request) {
	petitions, err := h.usecase.GetAllPetitions()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(petitions)
}

func (h *PetitionHandler) GetPetitionByID(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	petition, err := h.usecase.GetPetitionByID(uint(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(petition)
}

func (h *PetitionHandler) Vote(w http.ResponseWriter, r *http.Request) {
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

	var input struct {
		PetitionID uint   `json:"petition_id"`
		VoteType   string `json:"vote_type"`
	}
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.usecase.Vote(userID, input.PetitionID, input.VoteType)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *PetitionHandler) DeletePetition(w http.ResponseWriter, r *http.Request) {
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

	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	err = h.usecase.DeletePetition(uint(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

package auth

import (
	"VoteGolang/internals/data"
	"VoteGolang/internals/usecases"
	"VoteGolang/pkg/domain"
	"encoding/json"
	"net/http"
)

type AuthHandler struct {
	authUseCase  *usecases.AuthUseCase
	tokenManager domain.TokenManager
}

func NewAuthHandler(authUseCase *usecases.AuthUseCase, tokenManager domain.TokenManager) *AuthHandler {
	return &AuthHandler{
		authUseCase:  authUseCase,
		tokenManager: tokenManager,
	}
}

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	token, err := h.authUseCase.Login(req.Username, req.Password)
	if err != nil {
		http.Error(w, "Unauthorized: "+err.Error(), http.StatusUnauthorized)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"token": token})
}

func (h *AuthHandler) Register(w http.ResponseWriter, r *http.Request) {
	var req data.User

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	err := h.authUseCase.Register(&req)
	if err != nil {
		http.Error(w, "Failed to register user: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "User registered successfully"})
}

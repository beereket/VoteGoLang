package handlers

import (
	"VoteGolang/internal/domain"
	"VoteGolang/internal/service"
	"encoding/json"
	"net/http"
)

type UserHandler struct {
	Service *service.UserService
}

func NewUserHandler(s *service.UserService) *UserHandler {
	return &UserHandler{Service: s}
}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user domain.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "❌ Invalid input", http.StatusBadRequest)
		return
	}
	_, err := h.Service.Register(user, false)
	if err != nil {
		http.Error(w, "❌ Failed to create user", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(map[string]string{"message": "✅ User created"})
}

func (h *UserHandler) LoginUser(w http.ResponseWriter, r *http.Request) {
	var input domain.LoginInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "❌ Invalid input", http.StatusBadRequest)
		return
	}
	token, role, err := h.Service.Login(input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	json.NewEncoder(w).Encode(map[string]string{
		"token": token,
		"role":  role,
	})
}

func (h *UserHandler) CreateAdminUser(w http.ResponseWriter, r *http.Request) {
	var user domain.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "❌ Invalid input", http.StatusBadRequest)
		return
	}
	_, err := h.Service.Register(user, true)
	if err != nil {
		http.Error(w, "❌ Failed to create admin", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(map[string]string{"message": "✅ Admin created"})
}

func (h *UserHandler) ListAdminUsers(w http.ResponseWriter, r *http.Request) {
	admins, err := h.Service.ListAdmins()
	if err != nil {
		http.Error(w, "❌ Failed to fetch admins", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(admins)
}

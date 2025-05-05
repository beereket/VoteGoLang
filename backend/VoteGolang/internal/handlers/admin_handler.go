package handlers

import (
	"VoteGolang/internal/service"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type AdminHandler struct {
	Service *service.AdminService
}

func NewAdminHandler(s *service.AdminService) *AdminHandler {
	return &AdminHandler{Service: s}
}

func (h *AdminHandler) GetDashboard(w http.ResponseWriter, r *http.Request) {
	data, err := h.Service.GetDashboard()
	if err != nil {
		http.Error(w, "❌ Failed to get dashboard", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(data)
}

func (h *AdminHandler) GetVotesPerDay(w http.ResponseWriter, r *http.Request) {
	data, err := h.Service.GetVotesPerDay()
	if err != nil {
		http.Error(w, "❌ Failed to fetch votes per day", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(data)
}

func (h *AdminHandler) GetUserRegistrationsPerWeek(w http.ResponseWriter, r *http.Request) {
	data, err := h.Service.GetUserRegistrationsPerWeek()
	if err != nil {
		http.Error(w, "❌ Failed to fetch user registrations", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(data)
}

func (h *AdminHandler) BanUser(w http.ResponseWriter, r *http.Request) {
	id, err := getIDFromVars(w, r)
	if err != nil {
		return
	}
	if err := h.Service.BanUser(id); err != nil {
		http.Error(w, "❌ Failed to ban user", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(map[string]string{"message": "✅ User banned successfully"})
}

func (h *AdminHandler) UnbanUser(w http.ResponseWriter, r *http.Request) {
	id, err := getIDFromVars(w, r)
	if err != nil {
		return
	}
	if err := h.Service.UnbanUser(id); err != nil {
		http.Error(w, "❌ Failed to unban user", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(map[string]string{"message": "✅ User unbanned successfully"})
}

func (h *AdminHandler) ListAllUsers(w http.ResponseWriter, r *http.Request) {
	users, err := h.Service.ListAllUsers()
	if err != nil {
		http.Error(w, "❌ Failed to fetch users", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(users)
}

func getIDFromVars(w http.ResponseWriter, r *http.Request) (int, error) {
	vars := mux.Vars(r)
	idStr := vars["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "❌ Invalid ID", http.StatusBadRequest)
		return 0, err
	}
	return id, nil
}

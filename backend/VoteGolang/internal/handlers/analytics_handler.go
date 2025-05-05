package handlers

import (
	"VoteGolang/internal/service"
	"encoding/json"
	"net/http"
)

type AnalyticsHandler struct {
	Service *service.AnalyticsService
}

func NewAnalyticsHandler(s *service.AnalyticsService) *AnalyticsHandler {
	return &AnalyticsHandler{Service: s}
}

func (h *AnalyticsHandler) GetTopCandidatesPerParty(w http.ResponseWriter, r *http.Request) {
	data, err := h.Service.GetTopCandidatesPerParty()
	if err != nil {
		http.Error(w, "❌ Failed to get data", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(data)
}

func (h *AnalyticsHandler) GetPartyAnalytics(w http.ResponseWriter, r *http.Request) {
	data, err := h.Service.GetPartyAnalyticsWithPercentage()
	if err != nil {
		http.Error(w, "❌ Failed to get analytics", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(data)
}

func (h *AnalyticsHandler) GetTotalUsers(w http.ResponseWriter, r *http.Request) {
	h.getTotal("users", w)
}

func (h *AnalyticsHandler) GetTotalVotes(w http.ResponseWriter, r *http.Request) {
	h.getTotal("votes", w)
}

func (h *AnalyticsHandler) GetTotalPetitions(w http.ResponseWriter, r *http.Request) {
	h.getTotal("petitions", w)
}

func (h *AnalyticsHandler) GetTotalCandidates(w http.ResponseWriter, r *http.Request) {
	h.getTotal("candidates", w)
}

func (h *AnalyticsHandler) GetPartyVotes(w http.ResponseWriter, r *http.Request) {
	data, err := h.Service.GetPartyVotes()
	if err != nil {
		http.Error(w, "❌ Failed to fetch party votes", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(data)
}

func (h *AnalyticsHandler) getTotal(table string, w http.ResponseWriter) {
	count, err := h.Service.GetTotal(table)
	if err != nil {
		http.Error(w, "❌ Failed to fetch total", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(map[string]int{"total": count})
}

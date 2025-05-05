package handlers

import (
	"VoteGolang/internal/service"
	"encoding/json"
	"net/http"
)

type ActivityLogHandler struct {
	Service *service.ActivityLogService
}

func NewActivityLogHandler(s *service.ActivityLogService) *ActivityLogHandler {
	return &ActivityLogHandler{Service: s}
}

func (h *ActivityLogHandler) List(w http.ResponseWriter, r *http.Request) {
	logs, err := h.Service.List()
	if err != nil {
		http.Error(w, "❌ Failed to fetch activity logs", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(logs)
}

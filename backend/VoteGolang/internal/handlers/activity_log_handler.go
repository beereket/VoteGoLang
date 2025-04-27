package handlers

import (
	"VoteGolang/internal/database"
	"encoding/json"
	"net/http"
)

type ActivityLog struct {
	ID          int    `json:"id"`
	UserID      int    `json:"user_id"`
	Username    string `json:"username"`
	Action      string `json:"action"`
	Description string `json:"description"`
	CreatedAt   string `json:"created_at"`
}

func ListActivityLogs(w http.ResponseWriter, r *http.Request) {
	rows, err := database.DB.Query(`
		SELECT al.id, al.user_id, u.username, al.action, al.description, al.created_at
		FROM activity_logs al
		LEFT JOIN users u ON al.user_id = u.id
		ORDER BY al.created_at DESC
	`)
	if err != nil {
		http.Error(w, "❌ Failed to fetch activity logs", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var logs []ActivityLog
	for rows.Next() {
		var l ActivityLog
		err := rows.Scan(&l.ID, &l.UserID, &l.Username, &l.Action, &l.Description, &l.CreatedAt)
		if err != nil {
			http.Error(w, "❌ Error reading activity logs", http.StatusInternalServerError)
			return
		}
		logs = append(logs, l)
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(logs)
}

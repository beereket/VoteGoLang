package handlers

import (
	"VoteGolang/internal/service"
	"encoding/json"
	"fmt"
	"net/http"

	"VoteGolang/internal/database"
)

type VoteRequest struct {
	CandidateID   int    `json:"candidate_id"`
	CandidateType string `json:"candidate_type"`
}

func VoteForCandidate(w http.ResponseWriter, r *http.Request) {
	userIDInterface := r.Context().Value("user_id")
	if userIDInterface == nil {
		http.Error(w, "❌ Unauthorized", http.StatusUnauthorized)
		return
	}
	userID := userIDInterface.(int)
	fmt.Println(userID)

	var req VoteRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "❌ Invalid request body", http.StatusBadRequest)
		return
	}

	var count int
	err := database.DB.QueryRow(`
		SELECT COUNT(*)
		FROM votes
		WHERE user_id = ? AND candidate_type = ?
	`, userID, req.CandidateType).Scan(&count)

	if err != nil {
		http.Error(w, "❌ Failed to check previous vote", http.StatusInternalServerError)
		return
	}

	if count > 0 {
		http.Error(w, "⚠️ You have already voted for this election", http.StatusForbidden)
		return
	}

	// Insert vote
	_, err = database.DB.Exec(`
		INSERT INTO votes (user_id, candidate_id, candidate_type, created_at, updated_at)
		VALUES (?, ?, ?, NOW(), NOW())
	`, userID, req.CandidateID, req.CandidateType)

	if err != nil {
		http.Error(w, "❌ Failed to save vote", http.StatusInternalServerError)
		return
	}

	logErr := service.CreateActivityLog(userID, "vote", fmt.Sprintf("Voted for candidate ID %d in %s election", req.CandidateID, req.CandidateType))
	if logErr != nil {
		fmt.Println("⚠️ Failed to log activity:", logErr)
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "✅ Vote cast successfully",
	})
}

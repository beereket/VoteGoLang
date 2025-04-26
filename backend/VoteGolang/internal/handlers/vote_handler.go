package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"time"

	"VoteGolang/internal/database"
)

type VoteInput struct {
	CandidateID   int    `json:"candidate_id"`
	CandidateType string `json:"candidate_type"`
}

func VoteForCandidate(w http.ResponseWriter, r *http.Request) {
	var input VoteInput

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "❌ Invalid input", http.StatusBadRequest)
		return
	}

	username := r.Context().Value("username").(string)

	// Get user ID from username
	var userID int
	err := database.DB.QueryRow("SELECT id FROM users WHERE username = ?", username).Scan(&userID)
	if err == sql.ErrNoRows {
		http.Error(w, "❌ User not found", http.StatusUnauthorized)
		return
	} else if err != nil {
		http.Error(w, "❌ Server error", http.StatusInternalServerError)
		return
	}

	var exists int
	err = database.DB.QueryRow(`
		SELECT COUNT(*) FROM votes
		WHERE user_id = ? AND candidate_id = ? AND deleted_at IS NULL
	`, userID, input.CandidateID).Scan(&exists)

	if err != nil {
		http.Error(w, "❌ Server error", http.StatusInternalServerError)
		return
	}

	if exists > 0 {
		http.Error(w, "❌ You have already voted for this candidate", http.StatusBadRequest)
		return
	}

	_, err = database.DB.Exec(`
		INSERT INTO votes (user_id, candidate_id, candidate_type, created_at, updated_at)
		VALUES (?, ?, ?, ?, ?)
	`, userID, input.CandidateID, input.CandidateType, time.Now(), time.Now())

	if err != nil {
		http.Error(w, "❌ Failed to cast vote", http.StatusInternalServerError)
		return
	}

	_, err = database.DB.Exec(`
		UPDATE candidates
		SET votes = votes + 1
		WHERE id = ?
	`, input.CandidateID)

	if err != nil {
		http.Error(w, "❌ Failed to update candidate votes", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "✅ Vote cast successfully!",
	})
}

type CandidateResult struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Party string `json:"party"`
	Votes int    `json:"votes"`
}

func ListElectionResults(w http.ResponseWriter, r *http.Request) {
	rows, err := database.DB.Query(`
		SELECT id, name, party, votes
		FROM candidates
		WHERE deleted_at IS NULL
		ORDER BY votes DESC
	`)
	if err != nil {
		http.Error(w, "❌ Failed to fetch results", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var results []CandidateResult

	for rows.Next() {
		var c CandidateResult
		err := rows.Scan(&c.ID, &c.Name, &c.Party, &c.Votes)
		if err != nil {
			http.Error(w, "❌ Error reading results", http.StatusInternalServerError)
			return
		}
		results = append(results, c)
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(results)
}

type UserVote struct {
	CandidateName  string `json:"candidate_name"`
	CandidatePhoto string `json:"candidate_photo"`
	CandidateType  string `json:"candidate_type"`
}

func GetUserVotingHistory(w http.ResponseWriter, r *http.Request) {
	// Get username from JWT token (middleware)
	username := r.Context().Value("username").(string)

	// Find user id
	var userID int
	err := database.DB.QueryRow("SELECT id FROM users WHERE username = ?", username).Scan(&userID)
	if err != nil {
		http.Error(w, "❌ User not found", http.StatusUnauthorized)
		return
	}

	rows, err := database.DB.Query(`
		SELECT c.name, c.photo, v.candidate_type
		FROM votes v
		JOIN candidates c ON v.candidate_id = c.id
		WHERE v.user_id = ? AND v.deleted_at IS NULL
	`, userID)

	if err != nil {
		http.Error(w, "❌ Failed to fetch voting history", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var history []UserVote

	for rows.Next() {
		var uv UserVote
		if err := rows.Scan(&uv.CandidateName, &uv.CandidatePhoto, &uv.CandidateType); err != nil {
			http.Error(w, "❌ Error reading voting history", http.StatusInternalServerError)
			return
		}
		history = append(history, uv)
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(history)
}

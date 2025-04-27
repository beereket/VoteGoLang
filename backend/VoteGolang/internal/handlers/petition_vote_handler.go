package handlers

import (
	"encoding/json"
	"net/http"

	"VoteGolang/internal/database"
	"VoteGolang/internal/service"
)

var petitionVotingService = &service.PetitionVotingService{}

type PetitionVoteInput struct {
	PetitionID int    `json:"petition_id"`
	VoteType   string `json:"vote_type"`
}

func VoteOnPetition(w http.ResponseWriter, r *http.Request) {
	var input PetitionVoteInput

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "❌ Invalid input", http.StatusBadRequest)
		return
	}

	username := r.Context().Value("username").(string)

	var userID int
	err := database.DB.QueryRow("SELECT id FROM users WHERE username = ?", username).Scan(&userID)
	if err != nil {
		http.Error(w, "❌ User not found", http.StatusUnauthorized)
		return
	}

	err = petitionVotingService.CastPetitionVote(userID, input.PetitionID, input.VoteType)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "✅ Petition vote cast successfully!",
	})
}

type PetitionVoteResult struct {
	Title        string `json:"title"`
	VotesInFavor int    `json:"votes_in_favor"`
	VotesAgainst int    `json:"votes_against"`
}

func GetPetitionVoteResult(w http.ResponseWriter, r *http.Request) {
	rows, err := database.DB.Query(`
		SELECT title, vote_in_favor, vote_against
		FROM petitions
		WHERE deleted_at IS NULL
		ORDER BY created_at DESC
	`)
	if err != nil {
		http.Error(w, "❌ Failed to fetch petition results", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var results []PetitionVoteResult

	for rows.Next() {
		var p PetitionVoteResult
		err := rows.Scan(&p.Title, &p.VotesInFavor, &p.VotesAgainst)
		if err != nil {
			http.Error(w, "❌ Error reading petition results", http.StatusInternalServerError)
			return
		}
		results = append(results, p)
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(results)
}

type PetitionUserVote struct {
	PetitionTitle string `json:"petition_title"`
	VoteType      string `json:"vote_type"`
}

func GetUserPetitionVotingHistory(w http.ResponseWriter, r *http.Request) {
	username := r.Context().Value("username").(string)

	var userID int
	err := database.DB.QueryRow("SELECT id FROM users WHERE username = ?", username).Scan(&userID)
	if err != nil {
		http.Error(w, "❌ User not found", http.StatusUnauthorized)
		return
	}

	rows, err := database.DB.Query(`
		SELECT p.title, v.vote_type
		FROM petition_votes v
		JOIN petitions p ON v.petition_id = p.id
		WHERE v.user_id = ? AND v.deleted_at IS NULL
	`, userID)

	if err != nil {
		http.Error(w, "❌ Failed to fetch voting history", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var history []PetitionUserVote

	for rows.Next() {
		var uv PetitionUserVote
		if err := rows.Scan(&uv.PetitionTitle, &uv.VoteType); err != nil {
			http.Error(w, "❌ Error reading voting history", http.StatusInternalServerError)
			return
		}
		history = append(history, uv)
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(history)
}

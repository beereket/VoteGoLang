package handlers

import (
	"VoteGolang/internal/database"
	"VoteGolang/internal/models"
	"database/sql"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"time"
)

func ListPetitions(w http.ResponseWriter, r *http.Request) {
	rows, err := database.DB.Query(`
		SELECT id, title, description, photo, vote_in_favor, vote_against, status, deadline
		FROM petitions
		ORDER BY created_at DESC
	`)
	if err != nil {
		http.Error(w, "❌ Failed to fetch petitions", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var petitions []models.Petition

	for rows.Next() {
		var p models.Petition
		var deadline sql.NullTime

		err := rows.Scan(
			&p.ID, &p.Title, &p.Description, &p.Photo,
			&p.VotesInFavor, &p.VotesAgainst, &p.Status, &deadline,
		)
		if err != nil {
			http.Error(w, "❌ Failed reading petition", http.StatusInternalServerError)
			return
		}

		if deadline.Valid {
			p.Deadline = &deadline.Time
		} else {
			p.Deadline = nil
		}

		petitions = append(petitions, p)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(petitions)
}

type PetitionVoteRequest struct {
	PetitionID int    `json:"petition_id"`
	VoteType   string `json:"vote_type"`
}

func VoteOnPetition(w http.ResponseWriter, r *http.Request) {
	userIDInterface := r.Context().Value("user_id")
	if userIDInterface == nil {
		http.Error(w, "❌ Unauthorized", http.StatusUnauthorized)
		return
	}
	userID := userIDInterface.(int)

	var req PetitionVoteRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "❌ Invalid request body", http.StatusBadRequest)
		return
	}

	// Check if user already voted for this petition
	var existingCount int
	err := database.DB.QueryRow(`
		SELECT COUNT(*)
		FROM petition_votes
		WHERE user_id = ? AND petition_id = ?
	`, userID, req.PetitionID).Scan(&existingCount)
	if err != nil {
		http.Error(w, "❌ Failed to check previous votes", http.StatusInternalServerError)
		return
	}

	if existingCount > 0 {
		http.Error(w, "⚠️ You have already voted for this petition", http.StatusForbidden)
		return
	}

	// Insert the vote
	_, err = database.DB.Exec(`
		INSERT INTO petition_votes (user_id, petition_id, vote_type, created_at, updated_at)
		VALUES (?, ?, ?, ?, ?)
	`, userID, req.PetitionID, req.VoteType, time.Now(), time.Now())

	if err != nil {
		http.Error(w, "❌ Failed to record vote", http.StatusInternalServerError)
		return
	}

	// Update petition vote counters
	if req.VoteType == "favor" {
		_, err = database.DB.Exec(`
			UPDATE petitions
			SET vote_in_favor = vote_in_favor + 1
			WHERE id = ?
		`, req.PetitionID)
	} else if req.VoteType == "against" {
		_, err = database.DB.Exec(`
			UPDATE petitions
			SET vote_against = vote_against + 1
			WHERE id = ?
		`, req.PetitionID)
	} else {
		http.Error(w, "❌ Invalid vote type", http.StatusBadRequest)
		return
	}

	if err != nil {
		http.Error(w, "❌ Failed to update petition votes", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "✅ Vote recorded successfully",
	})
}

func CreatePetition(w http.ResponseWriter, r *http.Request) {
	type CreatePetitionRequest struct {
		Title       string     `json:"title"`
		Description string     `json:"description"`
		Photo       string     `json:"photo"`
		Deadline    *time.Time `json:"deadline,omitempty"` // optional
	}

	var req CreatePetitionRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "❌ Invalid input", http.StatusBadRequest)
		return
	}

	if req.Title == "" || req.Description == "" {
		http.Error(w, "❌ Title and Description are required", http.StatusBadRequest)
		return
	}

	_, err := database.DB.Exec(`
		INSERT INTO petitions (user_id, title, description, photo, status, deadline, vote_in_favor, vote_against, created_at, updated_at)
		VALUES (?, ?, ?, ?, 'open', ?, 0, 0, NOW(), NOW())
	`, 1, req.Title, req.Description, req.Photo, req.Deadline) // 👈 Admin user_id = 1 or real admin ID

	if err != nil {
		http.Error(w, "❌ Failed to create petition", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "✅ Petition created successfully",
	})
}

func ClosePetition(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	petitionID := vars["id"]

	_, err := database.DB.Exec(`
		UPDATE petitions
		SET status = 'closed'
		WHERE id = ?
	`, petitionID)

	if err != nil {
		http.Error(w, "❌ Failed to close petition", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "✅ Petition closed successfully",
	})
}

func DeletePetition(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	petitionIDStr := vars["id"]
	petitionID, err := strconv.Atoi(petitionIDStr)
	if err != nil {
		http.Error(w, "❌ Invalid petition ID", http.StatusBadRequest)
		return
	}

	// Optional: check if petition exists first

	_, err = database.DB.Exec(`DELETE FROM petitions WHERE id = ?`, petitionID)
	if err != nil {
		http.Error(w, "❌ Failed to delete petition", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "✅ Petition deleted successfully!",
	})
}

func GetPetitionByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idParam := vars["id"]

	petitionID, err := strconv.Atoi(idParam)
	if err != nil {
		http.Error(w, "❌ Invalid petition ID", http.StatusBadRequest)
		return
	}

	var (
		id          int
		userID      int
		title       string
		photo       string
		description string
		voteInFavor int
		voteAgainst int
		createdAt   string
	)

	err = database.DB.QueryRow(`
		SELECT id, user_id, title, photo, description, vote_in_favor, vote_against, created_at
		FROM petitions
		WHERE id = ?
	`, petitionID).Scan(&id, &userID, &title, &photo, &description, &voteInFavor, &voteAgainst, &createdAt)

	if err != nil {
		http.Error(w, "❌ Petition not found", http.StatusNotFound)
		return
	}

	response := map[string]interface{}{
		"id":            id,
		"user_id":       userID,
		"title":         title,
		"photo":         photo,
		"description":   description,
		"vote_in_favor": voteInFavor,
		"vote_against":  voteAgainst,
		"created_at":    createdAt,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// petitions/:id/results
func GetPetitionResults(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	petitionID := vars["id"]

	var (
		votesInFavor int
		votesAgainst int
	)

	err := database.DB.QueryRow(`
		SELECT vote_in_favor, vote_against
		FROM petitions
		WHERE id = ?
	`, petitionID).Scan(&votesInFavor, &votesAgainst)

	if err != nil {
		http.Error(w, "❌ Failed to fetch petition results", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"vote_in_favor": votesInFavor,
		"vote_against":  votesAgainst,
	})
}

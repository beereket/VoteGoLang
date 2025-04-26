package handlers

import (
	"VoteGolang/internal/models"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"time"

	"VoteGolang/internal/database"
)

func ListPetitions(w http.ResponseWriter, r *http.Request) {
	rows, err := database.DB.Query(`
		SELECT id, title, photo, description, vote_in_favor, vote_against, created_at
		FROM petitions
		WHERE deleted_at IS NULL
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
		err := rows.Scan(&p.ID, &p.Title, &p.Photo, &p.Description, &p.VoteInFavor, &p.VoteAgainst, &p.CreatedAt)
		if err != nil {
			http.Error(w, "❌ Error reading petition", http.StatusInternalServerError)
			return
		}
		petitions = append(petitions, p)
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(petitions)
}

type VotePetitionInput struct {
	PetitionID int    `json:"petition_id"`
	VoteType   string `json:"vote_type"` // "favor" or "against"
}

func VoteOnPetition(w http.ResponseWriter, r *http.Request) {
	var input VotePetitionInput

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "❌ Invalid input", http.StatusBadRequest)
		return
	}

	var field string
	if input.VoteType == "favor" {
		field = "vote_in_favor"
	} else if input.VoteType == "against" {
		field = "vote_against"
	} else {
		http.Error(w, "❌ Invalid vote type", http.StatusBadRequest)
		return
	}

	_, err := database.DB.Exec(`
		UPDATE petitions
		SET `+field+` = `+field+` + 1, updated_at = NOW()
		WHERE id = ? AND deleted_at IS NULL
	`, input.PetitionID)

	if err != nil {
		http.Error(w, "❌ Failed to vote on petition", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "✅ Voted on petition successfully!",
	})
}

type PetitionInput struct {
	Title       string `json:"title"`
	Photo       string `json:"photo"`
	Description string `json:"description"`
}

func CreatePetition(w http.ResponseWriter, r *http.Request) {
	var input PetitionInput

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "❌ Invalid input", http.StatusBadRequest)
		return
	}

	// Get username from JWT token
	username := r.Context().Value("username").(string)

	// Find user ID
	var userID int
	err := database.DB.QueryRow("SELECT id FROM users WHERE username = ?", username).Scan(&userID)
	if err != nil {
		http.Error(w, "❌ User not found", http.StatusUnauthorized)
		return
	}

	_, err = database.DB.Exec(`
		INSERT INTO petitions (user_id, title, photo, description, created_at, updated_at)
		VALUES (?, ?, ?, ?, ?, ?)
	`, userID, input.Title, input.Photo, input.Description, time.Now(), time.Now())

	if err != nil {
		http.Error(w, "❌ Failed to create petition", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "✅ Petition created!"})
}

func DeletePetition(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "❌ Invalid petition ID", http.StatusBadRequest)
		return
	}

	// Soft delete the petition
	_, err = database.DB.Exec(`
		UPDATE petitions
		SET deleted_at = NOW()
		WHERE id = ? AND deleted_at IS NULL
	`, id)

	if err != nil {
		http.Error(w, "❌ Failed to delete petition", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "✅ Petition deleted (soft deleted)",
	})
}

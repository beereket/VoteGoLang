package handlers

import (
	"VoteGolang/internal/database"
	"VoteGolang/internal/models"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"time"
)

func ListCandidates(w http.ResponseWriter, r *http.Request) {
	rows, err := database.DB.Query(`
		SELECT id, name, photo, education, age, party, region, type
		FROM candidates
		WHERE deleted_at IS NULL
	`)
	if err != nil {
		http.Error(w, "❌ Failed to fetch candidates", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var candidates []models.Candidate

	for rows.Next() {
		var c models.Candidate
		err := rows.Scan(&c.ID, &c.Name, &c.Photo, &c.Education, &c.Age, &c.Party, &c.Region, &c.Type)
		if err != nil {
			http.Error(w, "❌ Error reading candidate", http.StatusInternalServerError)
			return
		}
		candidates = append(candidates, c)
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(candidates)
}

func CreateCandidate(w http.ResponseWriter, r *http.Request) {
	var candidate models.Candidate

	if err := json.NewDecoder(r.Body).Decode(&candidate); err != nil {
		http.Error(w, "❌ Invalid candidate input", http.StatusBadRequest)
		return
	}

	query := `
	INSERT INTO candidates (name, photo, education, age, party, region, votes, type, created_at, updated_at)
	VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
	`

	_, err := database.DB.Exec(query,
		candidate.Name,
		candidate.Photo,
		candidate.Education,
		candidate.Age,
		candidate.Party,
		candidate.Region,
		0, // votes start at 0
		candidate.Type,
		time.Now(),
		time.Now(),
	)

	if err != nil {
		http.Error(w, "❌ Failed to create candidate", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "✅ Candidate created"})
}

func UpdateCandidate(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "❌ Invalid candidate ID", http.StatusBadRequest)
		return
	}

	var input struct {
		Name      string `json:"name"`
		Photo     string `json:"photo"`
		Education string `json:"education"`
		Age       int    `json:"age"`
		Party     string `json:"party"`
		Region    string `json:"region"`
		Type      string `json:"type"`
	}

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "❌ Invalid input", http.StatusBadRequest)
		return
	}

	query := `
		UPDATE candidates
		SET name = ?, photo = ?, education = ?, age = ?, party = ?, region = ?, type = ?, updated_at = ?
		WHERE id = ? AND deleted_at IS NULL
	`

	_, err = database.DB.Exec(query,
		input.Name,
		input.Photo,
		input.Education,
		input.Age,
		input.Party,
		input.Region,
		input.Type,
		time.Now(),
		id,
	)
	if err != nil {
		http.Error(w, "❌ Failed to update candidate", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "✅ Candidate updated successfully",
	})
}

func DeleteCandidate(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "❌ Invalid candidate ID", http.StatusBadRequest)
		return
	}

	query := `
	UPDATE candidates
	SET deleted_at = NOW()
	WHERE id = ? AND deleted_at IS NULL
	`

	_, err = database.DB.Exec(query, id)
	if err != nil {
		http.Error(w, "❌ Failed to delete candidate", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "✅ Candidate deleted (soft)"})
}

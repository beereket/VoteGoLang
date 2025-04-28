package handlers

import (
	"VoteGolang/internal/database"
	"VoteGolang/internal/models"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func ListCandidatesByElection(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	electionID := vars["election_id"]

	rows, err := database.DB.Query(`
		SELECT id, name, party, region, age, type, votes
		FROM candidates
		WHERE election_id = ?
	`, electionID)
	if err != nil {
		http.Error(w, "❌ Failed to fetch candidates", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var candidates []models.Candidate

	for rows.Next() {
		var c models.Candidate
		err := rows.Scan(&c.ID, &c.Name, &c.Party, &c.Region, &c.Age, &c.Type, &c.Votes)
		if err != nil {
			http.Error(w, "❌ Error reading candidate", http.StatusInternalServerError)
			return
		}
		candidates = append(candidates, c)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(candidates)
}

func CreateCandidate(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Name       string `json:"name"`
		Photo      string `json:"photo"`
		Education  string `json:"education"`
		Age        int    `json:"age"`
		Party      string `json:"party"`
		Region     string `json:"region"`
		Type       string `json:"type"`
		ElectionID int    `json:"election_id"` // VERY IMPORTANT
	}

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "❌ Invalid input", http.StatusBadRequest)
		return
	}

	_, err := database.DB.Exec(`
		INSERT INTO candidates (name, photo, education, age, party, region, type, election_id, created_at, updated_at)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, NOW(), NOW())
	`, input.Name, input.Photo, input.Education, input.Age, input.Party, input.Region, input.Type, input.ElectionID)

	if err != nil {
		http.Error(w, "❌ Failed to create candidate", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "✅ Candidate created successfully",
	})
}
func UpdateCandidate(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "❌ Invalid candidate ID", http.StatusBadRequest)
		return
	}

	var input models.Candidate
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "❌ Invalid input", http.StatusBadRequest)
		return
	}

	_, err = database.DB.Exec(`
		UPDATE candidates
		SET name = ?, party = ?, region = ?, age = ?, type = ?, updated_at = NOW()
		WHERE id = ?
	`, input.Name, input.Party, input.Region, input.Age, input.Type, id)

	if err != nil {
		http.Error(w, "❌ Failed to update candidate", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "✅ Candidate updated successfully"})
}

func DeleteCandidate(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "❌ Invalid candidate ID", http.StatusBadRequest)
		return
	}

	_, err = database.DB.Exec(`DELETE FROM candidates WHERE id = ?`, id)
	if err != nil {
		http.Error(w, "❌ Failed to delete candidate", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "✅ Candidate deleted successfully"})
}

package handlers

import (
	"VoteGolang/internal/database"
	"VoteGolang/internal/models"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

func ListElections(w http.ResponseWriter, r *http.Request) {
	rows, err := database.DB.Query(`
		SELECT id, name, description, election_date, created_at, updated_at
		FROM elections
		ORDER BY election_date ASC
	`)
	if err != nil {
		http.Error(w, "❌ Failed to fetch elections", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var elections []models.Election

	for rows.Next() {
		var e models.Election
		err := rows.Scan(&e.ID, &e.Name, &e.Description, &e.ElectionDate, &e.CreatedAt, &e.UpdatedAt)
		if err != nil {
			http.Error(w, "❌ Error reading election", http.StatusInternalServerError)
			return
		}
		elections = append(elections, e)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(elections)
}

func GetElectionDetail(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	electionID := vars["id"]

	// Fetch election info
	var election models.Election
	err := database.DB.QueryRow(`
		SELECT id, name, description, election_date, created_at, updated_at
		FROM elections
		WHERE id = ?
	`, electionID).Scan(&election.ID, &election.Name, &election.Description, &election.ElectionDate, &election.CreatedAt, &election.UpdatedAt)

	if err != nil {
		http.Error(w, "❌ Election not found", http.StatusNotFound)
		return
	}

	// Fetch candidates linked to election
	candidateRows, err := database.DB.Query(`
		SELECT id, name, photo, education, age, party, region, votes, election_id, type, created_at, updated_at
		FROM candidates
		WHERE election_id = ?
	`, electionID)
	if err != nil {
		http.Error(w, "❌ Failed to fetch candidates", http.StatusInternalServerError)
		return
	}
	defer candidateRows.Close()

	var candidates []models.Candidate
	for candidateRows.Next() {
		var c models.Candidate
		err := candidateRows.Scan(
			&c.ID,
			&c.Name,
			&c.Photo,
			&c.Education,
			&c.Age,
			&c.Party,
			&c.Region,
			&c.Votes,
			&c.ElectionID,
			&c.Type,
			&c.CreatedAt,
			&c.UpdatedAt,
		)
		if err != nil {
			http.Error(w, "❌ Error reading candidate", http.StatusInternalServerError)
			return
		}
		candidates = append(candidates, c)
	}

	userIDInterface := r.Context().Value("user_id")
	var alreadyVoted bool
	if userIDInterface != nil {
		userID := userIDInterface.(int)

		var count int
		err := database.DB.QueryRow(`
			SELECT COUNT(*)
			FROM votes
			WHERE user_id = ? AND election_id = ?
		`, userID, electionID).Scan(&count)

		if err == nil && count > 0 {
			alreadyVoted = true
		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"election":     election,
		"candidates":   candidates,
		"alreadyVoted": alreadyVoted,
	})
}

func GetElectionWithCandidates(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	electionID := vars["election_id"]

	var election models.Election
	err := database.DB.QueryRow(`
        SELECT id, name, description, election_date
        FROM elections
        WHERE id = ?
    `, electionID).Scan(&election.ID, &election.Name, &election.Description, &election.ElectionDate)

	if err != nil {
		http.Error(w, "❌ Election not found", http.StatusNotFound)
		return
	}

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
	json.NewEncoder(w).Encode(map[string]interface{}{
		"election":   election,
		"candidates": candidates,
	})
}

func GetElectionResults(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	electionID := vars["id"]

	rows, err := database.DB.Query(`
		SELECT name, party, votes
		FROM candidates
		WHERE election_id = ?
		ORDER BY votes DESC
	`, electionID)
	if err != nil {
		http.Error(w, "❌ Failed to fetch election results", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	type Result struct {
		Name  string `json:"name"`
		Party string `json:"party"`
		Votes int    `json:"votes"`
	}

	var results []Result
	for rows.Next() {
		var r Result
		err := rows.Scan(&r.Name, &r.Party, &r.Votes)
		if err != nil {
			http.Error(w, "❌ Error scanning result", http.StatusInternalServerError)
			return
		}
		results = append(results, r)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(results)
}

package handlers

import (
	"VoteGolang/internal/database"
	"encoding/json"
	"net/http"
)

// 📊 Total Users
func GetTotalUsers(w http.ResponseWriter, r *http.Request) {
	var total int
	err := database.DB.QueryRow("SELECT COUNT(*) FROM users WHERE deleted_at IS NULL").Scan(&total)
	if err != nil {
		http.Error(w, "❌ Failed to fetch users", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(map[string]int{"total": total})
}

// 📊 Total Votes
func GetTotalVotes(w http.ResponseWriter, r *http.Request) {
	var total int
	err := database.DB.QueryRow("SELECT COUNT(*) FROM votes WHERE deleted_at IS NULL").Scan(&total)
	if err != nil {
		http.Error(w, "❌ Failed to fetch votes", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(map[string]int{"total": total})
}

// 📊 Total Petitions
func GetTotalPetitions(w http.ResponseWriter, r *http.Request) {
	var total int
	err := database.DB.QueryRow("SELECT COUNT(*) FROM petitions WHERE deleted_at IS NULL").Scan(&total)
	if err != nil {
		http.Error(w, "❌ Failed to fetch petitions", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(map[string]int{"total": total})
}

// 📊 Total Candidates
func GetTotalCandidates(w http.ResponseWriter, r *http.Request) {
	var total int
	err := database.DB.QueryRow("SELECT COUNT(*) FROM candidates WHERE deleted_at IS NULL").Scan(&total)
	if err != nil {
		http.Error(w, "❌ Failed to fetch candidates", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(map[string]int{"total": total})
}

// 📊 Party Votes (for chart)
func GetPartyVotes(w http.ResponseWriter, r *http.Request) {
	rows, err := database.DB.Query(`
		SELECT party, SUM(votes) as total_votes
		FROM candidates
		WHERE deleted_at IS NULL
		GROUP BY party
		ORDER BY total_votes DESC
	`)
	if err != nil {
		http.Error(w, "❌ Failed to fetch party votes", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	type PartyVote struct {
		Party string `json:"party"`
		Votes int    `json:"votes"`
	}

	var results []PartyVote

	for rows.Next() {
		var pv PartyVote
		if err := rows.Scan(&pv.Party, &pv.Votes); err != nil {
			http.Error(w, "❌ Error reading party votes", http.StatusInternalServerError)
			return
		}
		results = append(results, pv)
	}

	json.NewEncoder(w).Encode(results)
}

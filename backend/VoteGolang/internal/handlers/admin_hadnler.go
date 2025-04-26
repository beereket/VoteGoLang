package handlers

import (
	"encoding/json"
	"net/http"

	"VoteGolang/internal/database"
)

type DashboardData struct {
	TotalUsers      int            `json:"total_users"`
	TotalVotes      int            `json:"total_votes"`
	TotalCandidates int            `json:"total_candidates"`
	TopCandidates   []TopCandidate `json:"top_candidates"`
}

type TopCandidate struct {
	Name  string `json:"name"`
	Votes int    `json:"votes"`
}

func GetAdminDashboard(w http.ResponseWriter, r *http.Request) {
	var data DashboardData

	// Count total users
	err := database.DB.QueryRow("SELECT COUNT(*) FROM users WHERE deleted_at IS NULL").Scan(&data.TotalUsers)
	if err != nil {
		http.Error(w, "❌ Failed to count users", http.StatusInternalServerError)
		return
	}

	// Count total votes
	err = database.DB.QueryRow("SELECT COUNT(*) FROM votes WHERE deleted_at IS NULL").Scan(&data.TotalVotes)
	if err != nil {
		http.Error(w, "❌ Failed to count votes", http.StatusInternalServerError)
		return
	}

	// Count total candidates
	err = database.DB.QueryRow("SELECT COUNT(*) FROM candidates WHERE deleted_at IS NULL").Scan(&data.TotalCandidates)
	if err != nil {
		http.Error(w, "❌ Failed to count candidates", http.StatusInternalServerError)
		return
	}

	// Get top 3 candidates by votes
	rows, err := database.DB.Query(`
		SELECT name, votes
		FROM candidates
		WHERE deleted_at IS NULL
		ORDER BY votes DESC
		LIMIT 3
	`)
	if err != nil {
		http.Error(w, "❌ Failed to get top candidates", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var c TopCandidate
		err := rows.Scan(&c.Name, &c.Votes)
		if err != nil {
			http.Error(w, "❌ Error reading candidate", http.StatusInternalServerError)
			return
		}
		data.TopCandidates = append(data.TopCandidates, c)
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data)
}

func GetVotesPerDay(w http.ResponseWriter, r *http.Request) {
	rows, err := database.DB.Query(`
		SELECT DATE(created_at) as vote_date, COUNT(*) as vote_count
		FROM votes
		WHERE deleted_at IS NULL
		GROUP BY vote_date
		ORDER BY vote_date DESC
		LIMIT 7
	`)
	if err != nil {
		http.Error(w, "❌ Failed to fetch votes per day", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	type VoteDay struct {
		Date  string `json:"date"`
		Count int    `json:"count"`
	}

	var votes []VoteDay

	for rows.Next() {
		var v VoteDay
		if err := rows.Scan(&v.Date, &v.Count); err != nil {
			http.Error(w, "❌ Error reading votes", http.StatusInternalServerError)
			return
		}
		votes = append(votes, v)
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(votes)
}

func GetUserRegistrationsPerWeek(w http.ResponseWriter, r *http.Request) {
	rows, err := database.DB.Query(`
		SELECT YEARWEEK(created_at, 1) as year_week, COUNT(*) as user_count
		FROM users
		WHERE deleted_at IS NULL
		GROUP BY year_week
		ORDER BY year_week DESC
		LIMIT 4
	`)
	if err != nil {
		http.Error(w, "❌ Failed to fetch users per week", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	type UserWeek struct {
		YearWeek string `json:"year_week"`
		Count    int    `json:"count"`
	}

	var users []UserWeek

	for rows.Next() {
		var u UserWeek
		if err := rows.Scan(&u.YearWeek, &u.Count); err != nil {
			http.Error(w, "❌ Error reading users", http.StatusInternalServerError)
			return
		}
		users = append(users, u)
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(users)
}

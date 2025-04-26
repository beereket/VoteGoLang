package handlers

import (
	"encoding/json"
	"net/http"

	"VoteGolang/internal/database"
)

type CandidateAnalytics struct {
	CandidateName string `json:"candidate_name"`
	CandidateType string `json:"candidate_type"`
	Votes         int    `json:"votes"`
}

type ElectionAnalytics struct {
	TotalVotes int                  `json:"total_votes"`
	TotalUsers int                  `json:"total_users"`
	Candidates []CandidateAnalytics `json:"candidates"`
}

func GetElectionAnalytics(w http.ResponseWriter, r *http.Request) {
	var data ElectionAnalytics

	// Total Votes
	err := database.DB.QueryRow("SELECT COUNT(*) FROM votes WHERE deleted_at IS NULL").Scan(&data.TotalVotes)
	if err != nil {
		http.Error(w, "❌ Failed to count votes", http.StatusInternalServerError)
		return
	}

	// Total Users
	err = database.DB.QueryRow("SELECT COUNT(*) FROM users WHERE deleted_at IS NULL").Scan(&data.TotalUsers)
	if err != nil {
		http.Error(w, "❌ Failed to count users", http.StatusInternalServerError)
		return
	}

	// Votes per candidate
	rows, err := database.DB.Query(`
		SELECT name, type, votes
		FROM candidates
		WHERE deleted_at IS NULL
		ORDER BY votes DESC
	`)
	if err != nil {
		http.Error(w, "❌ Failed to fetch candidates", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var c CandidateAnalytics
		err := rows.Scan(&c.CandidateName, &c.CandidateType, &c.Votes)
		if err != nil {
			http.Error(w, "❌ Error reading candidate", http.StatusInternalServerError)
			return
		}
		data.Candidates = append(data.Candidates, c)
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data)
}

type CandidateDetails struct {
	Name   string `json:"name"`
	Type   string `json:"type"`
	Age    int    `json:"age"`
	Party  string `json:"party"`
	Region string `json:"region"`
	Votes  int    `json:"votes"`
}

func GetCandidateAnalytics(w http.ResponseWriter, r *http.Request) {
	rows, err := database.DB.Query(`
		SELECT name, type, age, party, region, votes
		FROM candidates
		WHERE deleted_at IS NULL
		ORDER BY votes DESC
	`)
	if err != nil {
		http.Error(w, "❌ Failed to fetch candidate analytics", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var candidates []CandidateDetails

	for rows.Next() {
		var c CandidateDetails
		err := rows.Scan(&c.Name, &c.Type, &c.Age, &c.Party, &c.Region, &c.Votes)
		if err != nil {
			http.Error(w, "❌ Error reading candidate data", http.StatusInternalServerError)
			return
		}
		candidates = append(candidates, c)
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(candidates)
}

type PartyAnalytics struct {
	PartyName  string `json:"party_name"`
	TotalVotes int    `json:"total_votes"`
}

func GetPartyAnalytics(w http.ResponseWriter, r *http.Request) {
	rows, err := database.DB.Query(`
		SELECT party, SUM(votes) as total_votes
		FROM candidates
		WHERE deleted_at IS NULL
		GROUP BY party
		ORDER BY total_votes DESC
	`)
	if err != nil {
		http.Error(w, "❌ Failed to fetch party analytics", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var parties []PartyAnalytics
	totalVotes := 0

	for rows.Next() {
		var p PartyAnalytics
		err := rows.Scan(&p.PartyName, &p.TotalVotes)
		if err != nil {
			http.Error(w, "❌ Error reading party data", http.StatusInternalServerError)
			return
		}
		parties = append(parties, p)
		totalVotes += p.TotalVotes
	}

	// Calculate percentages
	type PartyAnalyticsWithPercent struct {
		PartyName  string  `json:"party_name"`
		TotalVotes int     `json:"total_votes"`
		Percentage float64 `json:"percentage"`
	}

	var result []PartyAnalyticsWithPercent
	for _, p := range parties {
		percent := (float64(p.TotalVotes) / float64(totalVotes)) * 100
		result = append(result, PartyAnalyticsWithPercent{
			PartyName:  p.PartyName,
			TotalVotes: p.TotalVotes,
			Percentage: percent,
		})
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)
}

type PartyPercentAnalytics struct {
	PartyName  string  `json:"party_name"`
	TotalVotes int     `json:"total_votes"`
	Percentage float64 `json:"percentage"`
}

func GetPartyPercentageAnalytics(w http.ResponseWriter, r *http.Request) {
	rows, err := database.DB.Query(`
		SELECT party, SUM(votes) as total_votes
		FROM candidates
		WHERE deleted_at IS NULL
		GROUP BY party
	`)
	if err != nil {
		http.Error(w, "❌ Failed to fetch party analytics", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var parties []PartyPercentAnalytics
	totalVotes := 0

	// First: read all party votes
	for rows.Next() {
		var p PartyPercentAnalytics
		err := rows.Scan(&p.PartyName, &p.TotalVotes)
		if err != nil {
			http.Error(w, "❌ Error reading party data", http.StatusInternalServerError)
			return
		}
		parties = append(parties, p)
		totalVotes += p.TotalVotes
	}

	// Second: calculate % for each party
	for i := range parties {
		if totalVotes > 0 {
			parties[i].Percentage = (float64(parties[i].TotalVotes) / float64(totalVotes)) * 100
		} else {
			parties[i].Percentage = 0
		}
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(parties)
}

type TopCandidate struct {
	PartyName     string `json:"party_name"`
	CandidateName string `json:"candidate_name"`
	TotalVotes    int    `json:"total_votes"`
}

func GetTopCandidatesPerParty(w http.ResponseWriter, r *http.Request) {
	rows, err := database.DB.Query(`
		SELECT party, name, votes
		FROM candidates
		WHERE deleted_at IS NULL
		ORDER BY party ASC, votes DESC
	`)
	if err != nil {
		http.Error(w, "❌ Failed to fetch top candidates", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	partySeen := make(map[string]bool)
	var topCandidates []TopCandidate

	for rows.Next() {
		var party, name string
		var votes int

		err := rows.Scan(&party, &name, &votes)
		if err != nil {
			http.Error(w, "❌ Error reading candidate data", http.StatusInternalServerError)
			return
		}

		// If party not already added ➔ this is top candidate
		if !partySeen[party] {
			topCandidates = append(topCandidates, TopCandidate{
				PartyName:     party,
				CandidateName: name,
				TotalVotes:    votes,
			})
			partySeen[party] = true
		}
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(topCandidates)
}

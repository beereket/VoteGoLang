package service

import (
	"VoteGolang/internal/database"
	"VoteGolang/internal/domain"
	"fmt"
)

type AnalyticsService struct{}

func NewAnalyticsService() *AnalyticsService {
	return &AnalyticsService{}
}

func (s *AnalyticsService) GetTotal(table string) (int, error) {
	var count int
	err := database.DB.QueryRow(fmt.Sprintf("SELECT COUNT(*) FROM %s WHERE deleted_at IS NULL", table)).Scan(&count)
	return count, err
}

func (s *AnalyticsService) GetPartyVotes() ([]domain.PartyVote, error) {
	rows, err := database.DB.Query(`
		SELECT party, SUM(votes) as total_votes
		FROM candidates
		WHERE deleted_at IS NULL
		GROUP BY party
		ORDER BY total_votes DESC
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []domain.PartyVote
	for rows.Next() {
		var pv domain.PartyVote
		if err := rows.Scan(&pv.Party, &pv.Votes); err != nil {
			return nil, err
		}
		result = append(result, pv)
	}
	return result, nil
}

func (s *AnalyticsService) GetPartyAnalyticsWithPercentage() ([]domain.PartyAnalytics, error) {
	partyVotes, err := s.GetPartyVotes()
	if err != nil {
		return nil, err
	}
	total := 0
	for _, p := range partyVotes {
		total += p.Votes
	}

	var result []domain.PartyAnalytics
	for _, p := range partyVotes {
		percent := (float64(p.Votes) / float64(total)) * 100
		result = append(result, domain.PartyAnalytics{
			PartyName:  p.Party,
			TotalVotes: p.Votes,
			Percentage: percent,
		})
	}
	return result, nil
}

func (s *AnalyticsService) GetTopCandidatesPerParty() ([]domain.TopCandidateAnalytics, error) {
	rows, err := database.DB.Query(`
		SELECT party, name, votes
		FROM candidates
		WHERE deleted_at IS NULL
		ORDER BY party ASC, votes DESC
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	seen := map[string]bool{}
	var result []domain.TopCandidateAnalytics
	for rows.Next() {
		var party, name string
		var votes int
		if err := rows.Scan(&party, &name, &votes); err != nil {
			return nil, err
		}
		if !seen[party] {
			result = append(result, domain.TopCandidateAnalytics{
				PartyName:     party,
				CandidateName: name,
				TotalVotes:    votes,
			})
			seen[party] = true
		}
	}
	return result, nil
}

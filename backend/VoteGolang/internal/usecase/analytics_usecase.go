package usecase

import "VoteGolang/internal/domain"

type AnalyticsUseCase interface {
	GetTopCandidatesPerParty() ([]domain.TopCandidateAnalytics, error)
	GetPartyAnalyticsWithPercentage() ([]domain.PartyAnalytics, error)
	GetTotal(table string) (int, error)
	GetPartyVotes() ([]domain.PartyVote, error)
}

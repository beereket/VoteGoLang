package usecase

import "VoteGolang/internal/domain"

type CandidateUseCase interface {
	ListByElection(electionID string) ([]domain.Candidate, error)
	Create(candidate domain.Candidate) error
	Update(id int, candidate domain.Candidate) error
	Delete(id int) error
	CastVote(userID int, vote domain.VoteRequest) error
}

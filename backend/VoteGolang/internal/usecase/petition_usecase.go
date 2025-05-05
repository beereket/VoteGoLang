package usecase

import "VoteGolang/internal/domain"

type PetitionUseCase interface {
	List() ([]domain.Petition, error)
	GetByID(id int) (domain.Petition, error)
	Create(p domain.Petition) error
	Delete(id int) error
	Vote(userID int, req domain.PetitionVoteRequest) error
	GetResults(id int) (map[string]int, error)
}

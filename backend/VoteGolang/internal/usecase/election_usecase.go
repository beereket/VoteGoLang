package usecase

import "VoteGolang/internal/domain"

type ElectionUseCase interface {
	List() ([]domain.Election, error)
	Get(id int) (domain.Election, error)
	Create(e domain.Election) error
	Update(id int, e domain.Election) error
	Delete(id int) error
}

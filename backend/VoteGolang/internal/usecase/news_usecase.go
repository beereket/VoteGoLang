package usecase

import "VoteGolang/internal/domain"

type NewsUseCase interface {
	List(page, limit int) ([]domain.News, error)
	Create(input domain.CreateNewsInput) error
	Delete(id int) error
}

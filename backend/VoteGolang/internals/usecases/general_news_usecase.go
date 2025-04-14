package usecases

import (
	"VoteGolang/internals/data"
	"VoteGolang/internals/repositories"
)

type GeneralNewsUseCase struct {
	Repo repositories.GeneralNewsRepository
}

func NewGeneralNewsUseCase(repo repositories.GeneralNewsRepository) *GeneralNewsUseCase {
	return &GeneralNewsUseCase{Repo: repo}
}

func (uc *GeneralNewsUseCase) Create(news *data.GeneralNews) error {
	return uc.Repo.Create(news)
}

func (uc *GeneralNewsUseCase) GetAll() ([]data.GeneralNews, error) {
	return uc.Repo.GetAll()
}

func (uc *GeneralNewsUseCase) GetByID(id uint) (*data.GeneralNews, error) {
	return uc.Repo.GetByID(id)
}

func (uc *GeneralNewsUseCase) Update(news *data.GeneralNews) error {
	return uc.Repo.Update(news)
}

func (uc *GeneralNewsUseCase) Delete(id uint) error {
	return uc.Repo.Delete(id)
}

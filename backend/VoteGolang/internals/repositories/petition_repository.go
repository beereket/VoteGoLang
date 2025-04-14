package repositories

import (
	"VoteGolang/internals/data"
	"gorm.io/gorm"
)

type PetitionRepository interface {
	Create(petition *data.Petition) error
	GetAll() ([]data.Petition, error)
	GetByID(id uint) (*data.Petition, error)
	VoteInFavor(id uint) error
	VoteAgainst(id uint) error
	Delete(id uint) error
}

type petitionRepository struct {
	db *gorm.DB
}

func NewPetitionRepository(db *gorm.DB) PetitionRepository {
	return &petitionRepository{db: db}
}

func (r *petitionRepository) Create(petition *data.Petition) error {
	return r.db.Create(petition).Error
}

func (r *petitionRepository) GetAll() ([]data.Petition, error) {
	var petitions []data.Petition
	err := r.db.Find(&petitions).Error
	return petitions, err
}

func (r *petitionRepository) GetByID(id uint) (*data.Petition, error) {
	var petition data.Petition
	err := r.db.First(&petition, id).Error
	if err != nil {
		return nil, err
	}
	return &petition, nil
}

func (r *petitionRepository) VoteInFavor(id uint) error {
	return r.db.Model(&data.Petition{}).
		Where("id = ?", id).
		Update("votes_in_favor", gorm.Expr("votes_in_favor + ?", 1)).
		Error
}

func (r *petitionRepository) VoteAgainst(id uint) error {
	return r.db.Model(&data.Petition{}).
		Where("id = ?", id).
		Update("votes_against", gorm.Expr("votes_against + ?", 1)).
		Error
}

func (r *petitionRepository) Delete(id uint) error {
	return r.db.Delete(&data.Petition{}, id).Error
}

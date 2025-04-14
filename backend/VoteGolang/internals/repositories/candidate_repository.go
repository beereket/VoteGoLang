package repositories

import (
	"VoteGolang/internals/data"
	"gorm.io/gorm"
)

type CandidateRepository interface {
	GetAllByType(candidateType string) ([]data.Candidate, error)
	GetByID(id uint) (*data.Candidate, error)
	IncrementVote(id uint) error
}

type candidateRepository struct {
	db *gorm.DB
}

func NewCandidateRepository(db *gorm.DB) CandidateRepository {
	return &candidateRepository{db: db}
}

func (r *candidateRepository) GetAllByType(candidateType string) ([]data.Candidate, error) {
	var candidates []data.Candidate
	err := r.db.Where("type = ?", candidateType).Find(&candidates).Error
	return candidates, err
}

func (r *candidateRepository) GetByID(id uint) (*data.Candidate, error) {
	var candidate data.Candidate
	err := r.db.First(&candidate, id).Error
	if err != nil {
		return nil, err
	}
	return &candidate, nil
}

func (r *candidateRepository) IncrementVote(id uint) error {
	return r.db.Model(&data.Candidate{}).
		Where("id = ?", id).
		UpdateColumn("votes", gorm.Expr("votes + ?", 1)).Error
}

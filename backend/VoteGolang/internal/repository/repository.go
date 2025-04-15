package repository

import (
	"VoteGolang/internal/data"
	"gorm.io/gorm"
)

type VoteRepository interface {
	FindAllVotes(db *gorm.DB) ([]data.Vote, error)
	DeleteVoteByID(db *gorm.DB, id uint) error
}

type UserRepository interface {
	FindAllUsers(db *gorm.DB) ([]data.User, error)
}

type PetitionRepository interface {
	FindAllPetitions(db *gorm.DB) ([]data.Petition, error)
	CreatePetition(db *gorm.DB, petition *data.Petition) error
	DeletePetitionByID(db *gorm.DB, id uint) error
}

type CandidateRepository interface {
	FindAllCandidates(db *gorm.DB) ([]data.Candidate, error)
}

// VoteRepository implementation
type voteRepo struct{}

func (r *voteRepo) FindAllVotes(db *gorm.DB) ([]data.Vote, error) {
	var votes []data.Vote
	err := db.Find(&votes).Error
	return votes, err
}

func (r *voteRepo) DeleteVoteByID(db *gorm.DB, id uint) error {
	return db.Delete(&data.Vote{}, id).Error
}

// UserRepository implementation
type userRepo struct{}

func (r *userRepo) FindAllUsers(db *gorm.DB) ([]data.User, error) {
	var users []data.User
	err := db.Find(&users).Error
	return users, err
}

// PetitionRepository implementation
type petitionRepo struct{}

func (r *petitionRepo) FindAllPetitions(db *gorm.DB) ([]data.Petition, error) {
	var petitions []data.Petition
	err := db.Find(&petitions).Error
	return petitions, err
}

func (r *petitionRepo) CreatePetition(db *gorm.DB, petition *data.Petition) error {
	return db.Create(petition).Error
}

func (r *petitionRepo) DeletePetitionByID(db *gorm.DB, id uint) error {
	return db.Delete(&data.Petition{}, id).Error
}

// CandidateRepository implementation
type candidateRepo struct{}

func (r *candidateRepo) FindAllCandidates(db *gorm.DB) ([]data.Candidate, error) {
	var candidates []data.Candidate
	err := db.Find(&candidates).Error
	return candidates, err
}

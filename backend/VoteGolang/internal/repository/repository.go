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

type VoteRepo struct{}

func (r *VoteRepo) FindAllVotes(db *gorm.DB) ([]data.Vote, error) {
	var votes []data.Vote
	err := db.Find(&votes).Error
	return votes, err
}

func (r *VoteRepo) DeleteVoteByID(db *gorm.DB, id uint) error {
	return db.Delete(&data.Vote{}, id).Error
}

type UserRepo struct{}

func (r *UserRepo) FindAllUsers(db *gorm.DB) ([]data.User, error) {
	var users []data.User
	err := db.Find(&users).Error
	return users, err
}

type PetitionRepo struct{}

func (r *PetitionRepo) FindAllPetitions(db *gorm.DB) ([]data.Petition, error) {
	var petitions []data.Petition
	err := db.Find(&petitions).Error
	return petitions, err
}

func (r *PetitionRepo) CreatePetition(db *gorm.DB, petition *data.Petition) error {
	return db.Create(petition).Error
}

func (r *PetitionRepo) DeletePetitionByID(db *gorm.DB, id uint) error {
	return db.Delete(&data.Petition{}, id).Error
}

type CandidateRepo struct{}

func (r *CandidateRepo) FindAllCandidates(db *gorm.DB) ([]data.Candidate, error) {
	var candidates []data.Candidate
	err := db.Find(&candidates).Error
	return candidates, err
}

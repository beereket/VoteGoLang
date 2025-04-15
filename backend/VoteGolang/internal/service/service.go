package service

import (
	"VoteGolang/internal/data"
	"VoteGolang/internal/repository"
	"gorm.io/gorm"
)

// Service interfaces
type VoteService interface {
	GetAllVotes(db *gorm.DB) ([]data.Vote, error)
	DeleteVoteByID(db *gorm.DB, id uint) error
}

type UserService interface {
	GetAllUsers(db *gorm.DB) ([]data.User, error)
}

type PetitionService interface {
	GetAllPetitions(db *gorm.DB) ([]data.Petition, error)
	CreatePetition(db *gorm.DB, petition *data.Petition) error
	DeletePetitionByID(db *gorm.DB, id uint) error
}

type CandidateService interface {
	GetAllCandidates(db *gorm.DB) ([]data.Candidate, error)
}

// Service implementation
type voteService struct {
	repo repository.VoteRepository
}

func (s *voteService) GetAllVotes(db *gorm.DB) ([]data.Vote, error) {
	return s.repo.FindAllVotes(db)
}

func (s *voteService) DeleteVoteByID(db *gorm.DB, id uint) error {
	return s.repo.DeleteVoteByID(db, id)
}

type userService struct {
	repo repository.UserRepository
}

func (s *userService) GetAllUsers(db *gorm.DB) ([]data.User, error) {
	return s.repo.FindAllUsers(db)
}

type petitionService struct {
	repo repository.PetitionRepository
}

func (s *petitionService) GetAllPetitions(db *gorm.DB) ([]data.Petition, error) {
	return s.repo.FindAllPetitions(db)
}

func (s *petitionService) CreatePetition(db *gorm.DB, petition *data.Petition) error {
	return s.repo.CreatePetition(db, petition)
}

func (s *petitionService) DeletePetitionByID(db *gorm.DB, id uint) error {
	return s.repo.DeletePetitionByID(db, id)
}

type candidateService struct {
	repo repository.CandidateRepository
}

func (s *candidateService) GetAllCandidates(db *gorm.DB) ([]data.Candidate, error) {
	return s.repo.FindAllCandidates(db)
}

// Factory functions to instantiate services
func NewVoteService(repo repository.VoteRepository) VoteService {
	return &voteService{repo}
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{repo}
}

func NewPetitionService(repo repository.PetitionRepository) PetitionService {
	return &petitionService{repo}
}

func NewCandidateService(repo repository.CandidateRepository) CandidateService {
	return &candidateService{repo}
}

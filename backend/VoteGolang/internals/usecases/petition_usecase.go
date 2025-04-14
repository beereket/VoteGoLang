package usecases

import (
	"VoteGolang/internals/data"
	"VoteGolang/internals/repositories"
)

type PetitionUseCase interface {
	CreatePetition(p *data.Petition) error
	GetAllPetitions() ([]data.Petition, error)
	GetPetitionByID(id uint) (*data.Petition, error)
	Vote(userID uint, petitionID uint, voteType string) error
	DeletePetition(id uint) error
	HasUserVoted(userID uint, petitionID uint) (bool, error)
}

type petitionUseCase struct {
	petitionRepo     repositories.PetitionRepository
	petitionVoteRepo repositories.PetitionVoteRepository
}

func NewPetitionUseCase(pr repositories.PetitionRepository, pvr repositories.PetitionVoteRepository) PetitionUseCase {
	return &petitionUseCase{
		petitionRepo:     pr,
		petitionVoteRepo: pvr,
	}
}

func (uc *petitionUseCase) CreatePetition(p *data.Petition) error {
	return uc.petitionRepo.Create(p)
}

func (uc *petitionUseCase) GetAllPetitions() ([]data.Petition, error) {
	return uc.petitionRepo.GetAll()
}

func (uc *petitionUseCase) GetPetitionByID(id uint) (*data.Petition, error) {
	return uc.petitionRepo.GetByID(id)
}

func (uc *petitionUseCase) Vote(userID uint, petitionID uint, voteType string) error {
	voted, err := uc.petitionVoteRepo.HasUserVoted(userID, petitionID)
	if err != nil {
		return err
	}
	if voted {
		return nil // or custom error like: errors.New("user already voted")
	}

	vote := &data.PetitionVote{
		UserID:     userID,
		PetitionID: petitionID,
		VoteType:   voteType,
	}
	return uc.petitionVoteRepo.CreateVote(vote)
}

func (uc *petitionUseCase) DeletePetition(id uint) error {
	return uc.petitionRepo.Delete(id)
}

func (uc *petitionUseCase) HasUserVoted(userID uint, petitionID uint) (bool, error) {
	return uc.petitionVoteRepo.HasUserVoted(userID, petitionID)
}

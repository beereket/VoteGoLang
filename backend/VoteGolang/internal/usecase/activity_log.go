package usecase

import "VoteGolang/internal/domain"

type ActivityLogUseCase interface {
	List() ([]domain.ActivityLog, error)
	Create(userID int, action, description string) error
}

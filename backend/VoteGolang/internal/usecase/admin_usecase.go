package usecase

import "VoteGolang/internal/domain"

type AdminUseCase interface {
	GetDashboardData() (*domain.DashboardData, error)
	GetVotesPerDay() ([]domain.VoteDay, error)
	GetUserRegistrationsPerWeek() ([]domain.UserWeek, error)
	BanUser(id int) error
	UnbanUser(id int) error
	ListAllUsers() ([]any, error)
}

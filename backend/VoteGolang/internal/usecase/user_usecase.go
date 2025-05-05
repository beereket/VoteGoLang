package usecase

import "VoteGolang/internal/domain"

type UserUseCase interface {
	Register(user domain.User, isAdmin bool) (int, error)
	Login(input domain.LoginInput) (token string, role string, err error)
	ListAdmins() ([]domain.User, error)
}

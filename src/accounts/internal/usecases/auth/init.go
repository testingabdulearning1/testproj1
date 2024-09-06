package usecases

import (
	"school-management-app/src/accounts/domain/repositories"
	"school-management-app/src/accounts/domain/usecases"
)

type authUseCase struct {
	repo repositories.AuthRepository
}

func NewUserUseCase(repo repositories.AuthRepository) usecases.AuthUseCase {
	return &authUseCase{
		repo: repo,
	}
}

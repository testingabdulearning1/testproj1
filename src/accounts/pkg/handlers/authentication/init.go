package handlers

import (
	"school-management-app/src/accounts/domain/usecases"
)

type UserHandler struct {
	usecase usecases.AuthUseCase
}

func NewUserHandler(usecase usecases.AuthUseCase) *UserHandler {
	return &UserHandler{
		usecase: usecase,
	}
}

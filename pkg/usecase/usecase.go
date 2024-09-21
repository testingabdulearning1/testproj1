package usecase

import (
	interfaces "school-management-app/pkg/repository/interface"
	usecases "school-management-app/pkg/usecase/interface"
)

type Usecase struct{
	repo interfaces.Repository
}


func NewUsecase(repo interfaces.Repository) usecases.Usecases{
	return Usecase{ repo: repo}
}
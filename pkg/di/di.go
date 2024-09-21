package di

import (
	"school-management-app/pkg/api/handlers"
	repositories "school-management-app/pkg/repository"
	"school-management-app/pkg/usecase"

	"gorm.io/gorm"
)

//dependency injection

type Handlers struct {
	Handler handlers.Handler
}

func GetHandlers(db *gorm.DB) *Handlers {
	Repo := repositories.NewRepository(db)
	UseCase := usecase.NewUsecase(Repo)
	SuperAdminH := handlers.NewHandler(UseCase)

	return &Handlers{
		Handler: *SuperAdminH,
	}
}

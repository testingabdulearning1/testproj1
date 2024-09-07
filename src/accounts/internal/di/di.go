package di

import (
	handlers "school-management-app/src/accounts/internal/handlers/authentication"
	"school-management-app/src/accounts/internal/repositories"
	usecases "school-management-app/src/accounts/internal/usecases/auth"

	"gorm.io/gorm"
)

//dependency injection

type Handlers struct {
	AuthHandler *handlers.UserHandler
}

func GetHandlers(db *gorm.DB) *Handlers { //replace interface{} with db type

	authRepo := repositories.NewUserRepo(db)
	authUseCase := usecases.NewUserUseCase(authRepo)
	authHandler := handlers.NewUserHandler(authUseCase)

	return &Handlers{
		AuthHandler: authHandler,
	}
}

package di

import (
	handlers "school-management-app/src/accounts/internal/handlers/authentication"
	"school-management-app/src/accounts/internal/repositories"
	usecases "school-management-app/src/accounts/internal/usecases/auth"
)

//dependency injection

type Handlers struct {
	AuthHandler *handlers.UserHandler
	// ABC *handlers.UserHandler
}

func GetHandlers(db interface{}) *Handlers { //replace interface{} with db type

	authRepo := repositories.NewUserRepo(db)
	authUseCase := usecases.NewUserUseCase(authRepo)
	authHandler := handlers.NewUserHandler(authUseCase)

	// abcRepo := repositories.NewUserRepo(config.Db)
	// abcUseCase := usecases.NewUserUseCase(abcRepo)
	// abcHandler := handlers.NewUserHandler(abcUseCase)

	return &Handlers{
		AuthHandler: authHandler,
		// ABC: abcHandler,
	}
}

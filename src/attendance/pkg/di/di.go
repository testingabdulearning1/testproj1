package di

import "gorm.io/gorm"

//dependency injection

type Handlers struct {
	// ABC *handlers.UserHandler
}

func GetHandlers(db *gorm.DB) *Handlers { //replace interface{} with db type

	// abcRepo := repositories.NewUserRepo(config.Db)
	// abcUseCase := usecases.NewUserUseCase(abcRepo)
	// abcHandler := handlers.NewUserHandler(abcUseCase)

	return &Handlers{
		// ABC: abcHandler,
	}
}

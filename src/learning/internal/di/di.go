package di

import (
	handlers "school-management-app/src/learning/internal/handlers/curriculum"
	"school-management-app/src/learning/internal/repositories"
	usecases "school-management-app/src/learning/internal/usecases/curriculum"

	"gorm.io/gorm"
)

//dependency injection

type Handlers struct {
	CurriculumHandler handlers.CurriculumHandler
}

func GetHandlers(db *gorm.DB) *Handlers {
	curriculumRepo := repositories.NewCurriculumRepo(db)
	curriculumUseCase := usecases.NewCurriculumUseCase(curriculumRepo)
	curriculumHandler := handlers.NewCurriculumHandler(curriculumUseCase)

	return &Handlers{
		CurriculumHandler: curriculumHandler,
	}
}

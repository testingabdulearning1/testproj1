package usecases

import (
	"school-management-app/src/learning/domain/repositories"
	"school-management-app/src/learning/domain/usecases"
)

type CurriculumUseCase struct {
	CurriculumRepo repositories.CurriculumRepo
}

func NewCurriculumUseCase(curriculumRepo repositories.CurriculumRepo) usecases.CurriculumUC {
	return &CurriculumUseCase{
		CurriculumRepo: curriculumRepo,
	}
}

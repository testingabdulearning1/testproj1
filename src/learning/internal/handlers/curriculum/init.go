package handlers

import "school-management-app/src/learning/domain/usecases"

type CurriculumHandler struct {
}

func NewCurriculumHandler(usecases.CurriculumUC) CurriculumHandler {
	return CurriculumHandler{}
}

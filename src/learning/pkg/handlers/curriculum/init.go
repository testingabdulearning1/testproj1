package handlers

import (
	"school-management-app/common/utils/validation"
	"school-management-app/src/learning/domain/entities/request"
	"school-management-app/src/learning/domain/usecases"

	"github.com/gofiber/fiber/v2"
)

type CurriculumHandler struct {
	CurriculumUC usecases.CurriculumUC
}

func NewCurriculumHandler(usecase usecases.CurriculumUC) CurriculumHandler {
	return CurriculumHandler{
		CurriculumUC: usecase,
	}
}

func (h *CurriculumHandler) TestHandler(c *fiber.Ctx) error {
	var req request.SampleReq
	if ok, err := validation.BindAndValidateRequest(c, &req); !ok {
		return err
	}
	resp := h.CurriculumUC.Sample(req)
	return resp.WriteToJSON(c)
}

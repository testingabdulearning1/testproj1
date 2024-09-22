package handlers

import (
	"school-management-app/pkg/domain/request"
	"school-management-app/pkg/utils/validation"
	usecases "school-management-app/pkg/usecase/interface"

	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	uc usecases.Usecases
}

func NewHandler(uc usecases.Usecases) *Handler {
	return &Handler{uc: uc}
}

func (h *Handler) SuperAdminSignin(c *fiber.Ctx) error {
	req := new(request.SuperAdminSigninReq)
	if ok, err := validation.BindAndValidateRequest(c, req); !ok {
		return err
	}

	response := h.uc.SuperAdminSignin(c.Context(), req)
	return response.WriteToJSON(c)
}

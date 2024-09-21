package handlers

import (
	usecases "school-management-app/pkg/usecase/interface"

	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	uc usecases.Usecases
}

func NewHandler(uc usecases.Usecases) *Handler {
	return &Handler{uc: uc}
}


func (h *Handler)SuperAdminSignin(c *fiber.Ctx)error{
	return nil
}
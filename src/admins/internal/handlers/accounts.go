package handlers

import (
	"school-management-app/common/schemas"
	"school-management-app/common/views"

	"github.com/gofiber/fiber/v2"
)

type UserController struct{}

func (u *UserController) LoginUser(c *fiber.Ctx) error {
	req := new(schemas.UserLogin1)
	if err := c.BodyParser(req); err != nil {
		return views.InvalidParams(c)
	}

	if err := req.Validate(); err != nil {
		return views.InvalidParams(c)
	}

	return views.DataView(c, "Iam Login api")
}

func (u *UserController) VerifyOtp(c *fiber.Ctx) error {
	req := new(schemas.ValidateOtp)
	if err := c.BodyParser(req); err != nil {
		return views.InvalidParams(c)
	}

	if err := req.Validate(); err != nil {
		return views.InvalidParams(c)
	}

	return views.DataView(c, "Iam Verify otp api")
}

func (u *UserController) ChangePassword(c *fiber.Ctx) error {
	req := new(schemas.ChangePassword)
	if err := c.BodyParser(req); err != nil {
		return views.InvalidParams(c)
	}

	if err := req.Validate(); err != nil {
		return views.InvalidParams(c)
	}

	return views.DataView(c, "Iam Change password api")
}

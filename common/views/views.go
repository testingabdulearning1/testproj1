package views

import "github.com/gofiber/fiber/v2"

func DataView(c *fiber.Ctx, data any) error {
	return c.
		Status(200).
		JSON(fiber.Map{
			"msg":  "ok",
			"data": data,
		})
}

func InvalidParams(c *fiber.Ctx) error {
	return c.
		Status(400).
		JSON(fiber.Map{
			"msg": "invalid params",
		})
}

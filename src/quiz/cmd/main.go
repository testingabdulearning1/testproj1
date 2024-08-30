package cmd

import (
	"fmt"
	"school-management-app/common/config"

	"github.com/gofiber/fiber/v2"
)

func helathCheck(c *fiber.Ctx) error {
	return c.Status(200).JSON(fiber.Map{
		"msg": "ok",
	})
}

func RunQuizServer() {
	app := fiber.New(fiber.Config{
		AppName:       "Infozio quiz app",
		StrictRouting: true,
	})

	// health check
	app.Get("/helath", helathCheck)

	port := config.PORT
	err := app.Listen(fmt.Sprintf(":%s", port))
	if err != nil {
		panic(err)
	}
}

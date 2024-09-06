package cmd

import (
	"fmt"
	"school-management-app/src/exam/config"
	route "school-management-app/src/exam/internal/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func healthCheck(c *fiber.Ctx) error {
	return c.Status(200).JSON(fiber.Map{
		"msg": "ok",
	})
}

func RunexamServer() {
	app := fiber.New(fiber.Config{
		AppName:       "Infozio exam app",
		StrictRouting: true,
	})
	app.Use(logger.New())

	// health check
	app.Get("/health", healthCheck)
	route.MountRoutes(app)

	err := app.Listen(fmt.Sprintf(":%s", config.Env.Port))
	if err != nil {
		panic(err)
	}
}

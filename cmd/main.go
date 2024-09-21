package main

import (
	"fmt"
	route "school-management-app/pkg/api/routes"
	"school-management-app/pkg/config"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func healthCheck(c *fiber.Ctx) error {
	return c.Status(200).JSON(fiber.Map{
		"msg": "ok",
	})
}

func main() {

	app := fiber.New(fiber.Config{
		AppName:       "Infozio app",
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

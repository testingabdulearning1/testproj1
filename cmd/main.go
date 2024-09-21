package main

import (
	"fmt"
	"school-management-app/pkg/config"
	route "school-management-app/pkg/routes"

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

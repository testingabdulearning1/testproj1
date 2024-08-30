package config

import (
	"fmt"
	"log"
	"school-management-app/common/config"
	"school-management-app/src/accounts/internal/routes"

	"github.com/gofiber/fiber/v2"
)

func healthCheck(c *fiber.Ctx) error {
	return c.SendString("OK")
}

func ConnectAccounts() {
	app := fiber.New(fiber.Config{
		AppName:       "SM USER",
		StrictRouting: true,
	})

	app.Get("/health", healthCheck)

	route.MountRoutes(app)

	port := config.PORT
	log.Println(config.PORT)
	log.Fatal(app.Listen(fmt.Sprintf(":%s", port)))
}

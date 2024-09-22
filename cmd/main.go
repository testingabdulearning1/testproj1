package main

import (
	"fmt"
	route "school-management-app/pkg/api/routes"
	"school-management-app/pkg/config"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
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
	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:5000, http://192.168.190.53:5000", // Add your specific origins
		AllowHeaders: "Origin, Content-Type, Accept",
		AllowMethods: "GET, POST, PUT, DELETE",
		AllowCredentials: true, // Enable credentials for specific origins
	}))

	// health check
	app.Get("/health", healthCheck)

	route.MountRoutes(app)

	err := app.Listen(fmt.Sprintf(":%s", config.Env.Port))
	if err != nil {
		panic("Couldn't start the server. Error: " + err.Error())
	}
}

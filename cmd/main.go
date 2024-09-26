package main

import (
	"github.com/gofiber/fiber/v2"
	"log"
)

func main() {
	// Initialize Fiber app
	app := fiber.New()

	// Define a basic route
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, Fiber!")
	})

	// Start the server
	if err := app.Listen(":3000"); err != nil {
		log.Fatalf("Error starting sderver: %v", err)
	}
}
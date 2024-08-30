package route

import "github.com/gofiber/fiber/v2"

func MountRoutes(app *fiber.App) {
	apiGroup := app.Group("/api")

	MountUser(apiGroup.Group("/user"))
}

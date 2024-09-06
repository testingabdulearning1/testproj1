package route

import (
	"school-management-app/src/accounts/config"
	"school-management-app/src/accounts/internal/di"

	"github.com/gofiber/fiber/v2"
)

func MountRoutes(app *fiber.App) {
	handlers := di.GetHandlers(config.PublicDB)
	apiGroup := app.Group("/api")

	MountStudentAuth(apiGroup.Group("/student"), handlers)
}

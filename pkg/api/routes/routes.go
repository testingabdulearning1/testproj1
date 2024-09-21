package route

import (
	"school-management-app/pkg/db"
	"school-management-app/pkg/di"

	"github.com/gofiber/fiber/v2"
)

func MountRoutes(app *fiber.App) {
	handlers := di.GetHandlers(db.PublicDB)
	saGroup := app.Group("/superadmin")
	saGroup.Post("/login",handlers.Handler.SuperAdminSignin)


}

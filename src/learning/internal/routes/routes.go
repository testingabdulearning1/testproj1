package route

import (
	cfg "school-management-app/common/config"
	"school-management-app/src/learning/internal/di"

	"github.com/gofiber/fiber/v2"
)

func MountRoutes(app *fiber.App) {
	handlers := di.GetHandlers(cfg.PublicDB)
	apiGroup := app.Group("/api")

	apiGroup.Post("/test", handlers.CurriculumHandler.TestHandler)
	MountCurriculumRoutes(apiGroup.Group("/curriculum"), handlers)
}

func MountCurriculumRoutes(curriculumGroup fiber.Router, handler *di.Handlers) {

}

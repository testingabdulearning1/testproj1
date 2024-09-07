package route

import (
	"school-management-app/src/learning/internal/db"
	"school-management-app/src/learning/internal/di"

	"github.com/gofiber/fiber/v2"
)

func MountRoutes(app *fiber.App) {
	handlers := di.GetHandlers(db.PublicDB)
	apiGroup := app.Group("/api")

	MountCurriculumRoutes(apiGroup.Group("/curriculum"), handlers)
}

func MountCurriculumRoutes(curriculumGroup fiber.Router,handler *di.Handlers){
	
}
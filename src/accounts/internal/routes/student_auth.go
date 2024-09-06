package route

import (
	// "school-management-app/src/accounts/internal/handlers/authentication"

	"school-management-app/src/accounts/internal/di"

	"github.com/gofiber/fiber/v2"
)

func MountStudentAuth(app fiber.Router, handlers *di.Handlers) {

	// non auth apis
	// app.Post("/student/login", handlers.AuthHandler.Login)
}

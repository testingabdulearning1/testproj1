package route

import (
	// "school-management-app/src/accounts/pkg/handlers/authentication"

	"school-management-app/src/accounts/pkg/di"

	"github.com/gofiber/fiber/v2"
)

func MountStudentAuth(app fiber.Router, handlers *di.Handlers) {

	// non auth apis
	// app.Post("/student/login", handlers.AuthHandler.Login)
}

package route

import (
	"school-management-app/src/accounts/internal/handlers"

	"github.com/gofiber/fiber/v2"
)

func MountUser(app fiber.Router) {
	thisController := handlers.UserController{}

	// non auth apis
	app.Post("/login", thisController.LoginUser)
	app.Post("/verify_otp", thisController.VerifyOtp)
	app.Post("/change_password", thisController.ChangePassword)
}

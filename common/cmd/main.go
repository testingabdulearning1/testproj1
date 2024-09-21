package main

// import (
// 	"fmt"
// 	"school-management-app/src/learning/config"
// 	accountroute "school-management-app/src/accounts/internal/routes"
// 	attendanceroute "school-management-app/src/attendance/internal/routes"
// 	examroute "school-management-app/src/exam/internal/routes"
// 	learningroute "school-management-app/src/learning/internal/routes"

// 	"github.com/gofiber/fiber/v2"
// 	"github.com/gofiber/fiber/v2/middleware/logger"
// )

// func healthCheck(c *fiber.Ctx) error {
// 	return c.Status(200).JSON(fiber.Map{
// 		"msg": "ok",
// 	})
// }

// func main() {
// 	app := fiber.New(fiber.Config{
// 		AppName:       "Infozio app",
// 		StrictRouting: true,
// 	})
// 	app.Use(logger.New())

// 	// health check
// 	app.Get("/health", healthCheck)

// 	accountroute.MountRoutes(app)
// 	attendanceroute.MountRoutes(app)
// 	examroute.MountRoutes(app)
// 	learningroute.MountRoutes(app)

// 	err := app.Listen(fmt.Sprintf(":%s", config.Env.Port))
// 	if err != nil {
// 		panic(err)
// 	}
// }

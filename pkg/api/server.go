package api

// import (
// 	"school-management-app/pkg/api/handlers"
// 	"school-management-app/pkg/di"

// 	"github.com/gofiber/fiber/v2"
// 	"github.com/gofiber/fiber/v2/middleware/logger"
// )

// type Server struct {
// 	server *fiber.App
// }

// func NewHttpServer(handler handlers.UserHandler) {
// 	app := fiber.New(fiber.Config{
// 		AppName:       "Infozio app",
// 		StrictRouting: true,
// 	})
// 	app.Use(logger.New())
// }

// func MountCurriculumRoutes(curriculumGroup fiber.Router, handler *di.Handlers) {

// }

// func (srv *Server) Run() {
// 	srv.server.Listen()
// }

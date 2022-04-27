package routes

import (
	"github.com/Nebojsa1999/XMLProjekat/backend/user_service/controller"
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	app.Post("/api/register", controller.RegisterANewUser)
	app.Post("/api/login", controller.Login)
	app.Get("/api/user/my-info", controller.GetInfoForLoggedInUser)
	app.Post("/api/logout", controller.Logout)
}

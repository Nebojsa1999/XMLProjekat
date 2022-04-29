package routes

import (
	"github.com/Nebojsa1999/XMLProjekat/backend/posting_service/controller"
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	app.Post("/api/register", controller.CreatePost)
}

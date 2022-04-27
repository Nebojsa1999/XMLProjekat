package main

import (
	"github.com/Nebojsa1999/XMLProjekat/backend/user_service/database"
	"github.com/Nebojsa1999/XMLProjekat/backend/user_service/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	database.Connect()

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))

	routes.Setup(app)

	app.Listen(":8000")
}

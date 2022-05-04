package main

import (
	"fmt"
	"github.com/Nebojsa1999/XMLProjekat/backend/user_service/database"
	"github.com/Nebojsa1999/XMLProjekat/backend/user_service/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"log"
)

func main() {
	database.Connect()
	defer database.Disconnect()

	// https://dev.to/koddr/go-fiber-by-examples-delving-into-built-in-functions-1p3k#configuration
	config := fiber.Config{
		Prefork: true,
	}
	app := fiber.New(config)

	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))

	routes.Setup(app)

	err := app.Listen(":8000")
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Server se pokrenuo na portu 8000.")
	}
}

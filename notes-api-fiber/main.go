package main

import (
	"github.com/gofiber/fiber/v2"
	"notes-fiber/database"
	"notes-fiber/router"
)

func main() {
	app := fiber.New()
	database.ConnectDB()

	router.SetupRoutes(app)
	app.Listen(":8000")
}

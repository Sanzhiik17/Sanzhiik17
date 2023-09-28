package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	noteRoutes "notes-fiber/internal/routes/note"
)

func SetupRoutes(app *fiber.App) {
	//app = fiber.App()
	api := app.Group("/api", logger.New())

	noteRoutes.SetupNoteRoutes(api)
	//user := api.Group("user")
	//
	//user.Get("/", func(c *fiber.Ctx) {})
	//
	//user.Get("/:userId", func(c *fiber.Ctx) {})
	//
	//user.Put("/:userId", func(c *fiber.Ctx) {})
}

package main

import (
	"github.com/AutonuKro/go-examples/rest-api/handlers"
	"github.com/gofiber/fiber/v2"
)

func generateApp() *fiber.App {
	app := fiber.New()
	// heath check
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.SendString("OK")
	})
	//library group
	libraryGroup := app.Group("/library")
	libraryGroup.Get("/", handlers.TestHandlers)

	return app
}

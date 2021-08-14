package api

import "github.com/gofiber/fiber/v2"

func SetupRoutes(app *fiber.App) {
	r := app.Group("/api")
	r.Get("/health", Health)
	r.Post("/token", GetToken)
}

package api

import (
	"strings"

	"github.com/Mshivam2409/hls-streamer/internal/db"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/spf13/viper"
)

func HTTPListen() error {

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
	}))

	SetupRoutes(app)

	app.Use(func(c *fiber.Ctx) error {
		tok := c.Get("x-gostreamer-token")
		// Check if header is valid
		if len(tok) == 21 {
			// Try to decode
			uri := strings.Split(c.Path(), "/")
			path := uri[len(uri)-2]
			rid, err := db.GoStreamer.BadgerClient.Get(tok)
			if err != nil || (path != rid) {
				return c.SendStatus(fiber.StatusUnauthorized)
			}
			return c.Next()
		}
		// Authentication failed
		return c.SendStatus(fiber.StatusUnauthorized)

	})

	app.Static("/hls", viper.GetString("cache.static"))

	err := app.Listen(":5000")
	if err != nil {
		return err
	}

	return nil
}

package api

import (
	"fmt"
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

	app.Static("/get", "./example")

	SetupRoutes(app)

	app.Use(func(c *fiber.Ctx) error {
		tok := c.Get("x-gostreamer-token")
		if len(tok) == 21 {
			uri := strings.Split(c.Path(), "/")
			path := uri[len(uri)-2]
			rid, err := db.GoStreamer.BadgerClient.Get(tok)
			if err != nil || (path != rid) {
				return c.SendStatus(fiber.StatusUnauthorized)
			}
			return c.Next()
		}
		return c.SendStatus(fiber.StatusUnauthorized)

	})

	app.Static("/hls", viper.GetString("cache.static"))

	err := app.Listen(fmt.Sprintf(":%d", viper.GetInt("port")))
	if err != nil {
		return err
	}

	return nil
}

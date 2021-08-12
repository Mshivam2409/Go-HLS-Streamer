package api

import (
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

	app.Static("/hls", viper.GetString("cache.static"))

	app.Static("/audio", "./example")

	err := app.Listen(":5000")
	if err != nil {
		return err
	}

	return nil
}

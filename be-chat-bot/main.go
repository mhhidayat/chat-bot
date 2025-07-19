package main

import (
	"chat_bot/configs"
	"chat_bot/constants"
	"chat_bot/handler"
	"chat_bot/middleware"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	conf := configs.NewConfig()

	app := fiber.New(fiber.Config{
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			code := fiber.StatusInternalServerError
			if e, ok := err.(*fiber.Error); ok {
				code = e.Code
			}
			return c.Status(code).JSON(fiber.Map{
				"status":  false,
				"message": err.Error(),
				"data":    nil,
			})
		},
	})

	// Setup middleware
	middleware.Setup(app, conf.Server.AllowOrigins)

	// Health check endpoint
	app.Get(constants.HealthPath, func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status":  "ok",
			"message": "Server is running",
		})
	})

	// API routes
	api := app.Group(constants.APIVersion)
	handler.NewHandler(api, conf)

	log.Fatal(app.Listen(constants.DefaultPort))
}

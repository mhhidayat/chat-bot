package main

import (
	"chat_bot/configs"
	"chat_bot/handler"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	conf := configs.NewConfig()
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: conf.Server.AllowOrigins,
	}))

	handler.NewHandler(app, conf)
	app.Listen(":3000")
}

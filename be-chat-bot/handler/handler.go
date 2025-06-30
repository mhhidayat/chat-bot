package handler

import (
	"chat_bot/aiutils"
	"chat_bot/configs"
	"chat_bot/dto"
	"chat_bot/jsonutils"
	"context"
	"net/http"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
)

type handler struct {
	conf *configs.Config
}

func NewHandler(app *fiber.App, conf *configs.Config) {
	h := handler{
		conf: conf,
	}

	app.Post("send-message", h.sendMessage)
}

func (h *handler) sendMessage(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(c.Context(), 10*time.Second)
	defer cancel()

	req := new(dto.SendMessageRequest)

	if err := c.BodyParser(req); err != nil || strings.TrimSpace(req.Prompt) == "" {
		return c.Status(http.StatusBadRequest).JSON(
			jsonutils.ErrorResponse("Prompt is required"),
		)
	}

	channel := make(chan string)
	go aiutils.GenerateAIResponse(ctx, req.Prompt, h.conf, channel)
	defer close(channel)

	return c.Status(http.StatusOK).JSON(
		jsonutils.SuccessResponse(
			"Successfully retrieved AI response",
			fiber.Map{
				"response": <-channel,
			},
		),
	)
}

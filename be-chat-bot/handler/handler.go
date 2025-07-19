package handler

import (
	"chat_bot/ai"
	"chat_bot/configs"
	"chat_bot/constants"
	"chat_bot/dto"
	"chat_bot/utils"
	"context"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
)

type handler struct {
	conf *configs.Config
}

func NewHandler(api fiber.Router, conf *configs.Config) {
	h := handler{
		conf: conf,
	}

	// Chat routes
	chat := api.Group("/chat")
	chat.Post("/send-message", h.sendMessage)
}

func (h *handler) sendMessage(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(c.Context(), constants.RequestTimeout*time.Second)
	defer cancel()

	req := new(dto.SendMessageRequest)

	if err := c.BodyParser(req); err != nil {
		return c.Status(http.StatusBadRequest).JSON(
			utils.ErrorResponse("Invalid request body"),
		)
	}

	if err := req.Validate(); err != nil {
		return c.Status(http.StatusBadRequest).JSON(
			utils.ErrorResponse(err.Error()),
		)
	}

	channel := make(chan *dto.SendMessageResponse)
	if req.Provider == constants.ProviderGemini {
		go ai.GenerateGeminiResponse(ctx, req, h.conf, channel)
	} else if req.Provider == constants.ProviderOpenRouter {
		go ai.GenerateOpenRouterResponse(ctx, req, h.conf, channel)
	} else {
		return c.Status(http.StatusInternalServerError).JSON(
			utils.ErrorResponse(
				"Provider unknown",
			),
		)
	}
	defer close(channel)

	aiResponse := <-channel
	if !aiResponse.Status {
		return c.Status(http.StatusInternalServerError).JSON(
			utils.ErrorResponse(
				"Failed to generate AI response: " + aiResponse.Response,
			),
		)
	}

	return c.Status(http.StatusOK).JSON(
		utils.SuccessResponse(
			"Successfully retrieved AI response",
			fiber.Map{
				"response": aiResponse.Response,
			},
		),
	)
}

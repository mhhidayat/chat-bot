package aiutils

import (
	"chat_bot/configs"
	"chat_bot/dto"
	"context"

	"google.golang.org/genai"
)

func GenerateGeminiResponse(ctx context.Context, sendMessageRequest *dto.SendMessageRequest, conf *configs.Config, c chan<- *dto.SendMessageResponse) {
	client, err := genai.NewClient(ctx, &genai.ClientConfig{
		APIKey:  conf.AI.GeminiApiKey,
		Backend: genai.BackendGeminiAPI,
	})
	if err != nil {
		c <- &dto.SendMessageResponse{
			Status:   false,
			Response: "Failed to create AI client: " + err.Error(),
		}
		return
	}

	config := &genai.GenerateContentConfig{
		SystemInstruction: genai.NewContentFromText(
			"Gunakan bahasa yang sama dengan input.",
			genai.RoleUser,
		),
	}

	result, err := client.Models.GenerateContent(
		ctx,
		sendMessageRequest.Model,
		genai.Text(sendMessageRequest.Prompt),
		config,
	)
	if err != nil {
		c <- &dto.SendMessageResponse{
			Status:   false,
			Response: "Failed to create AI client: " + err.Error(),
		}
		return
	}
	if result == nil {
		c <- &dto.SendMessageResponse{
			Status:   false,
			Response: "AI response is empty",
		}
		return
	}

	c <- &dto.SendMessageResponse{
		Status:   true,
		Response: result.Text(),
	}
}

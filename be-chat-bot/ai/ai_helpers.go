package ai

import (
	"bytes"
	"chat_bot/configs"
	"chat_bot/dto"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

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

func GenerateOpenRouterResponse(ctx context.Context, sendMessageRequest *dto.SendMessageRequest, conf *configs.Config, c chan<- *dto.SendMessageResponse) {
	payload := map[string]any{
		"model": sendMessageRequest.Model,
		"messages": []map[string]string{
			{
				"role":    "system",
				"content": "Tolong balas secepat mungkin, kalo bisa di bawah 10 detik, Gunakan bahasa yang sama dengan input.",
			},
			{
				"role":    "user",
				"content": sendMessageRequest.Prompt,
			},
		},
	}

	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		c <- &dto.SendMessageResponse{
			Status:   false,
			Response: "Failed encode JSON:" + err.Error(),
		}
		return
	}

	req, err := http.NewRequestWithContext(ctx, "POST", conf.AI.OpenRouterApiUrl, bytes.NewBuffer(jsonPayload))
	if err != nil {
		c <- &dto.SendMessageResponse{
			Status:   false,
			Response: "Failed to create AI client:" + err.Error(),
		}
		return
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", conf.AI.OpenRouterApiKey))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		c <- &dto.SendMessageResponse{
			Status:   false,
			Response: "Failed sending request:" + err.Error(),
		}
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		c <- &dto.SendMessageResponse{
			Status:   false,
			Response: "Failed to read response:" + err.Error(),
		}
		return
	}

	openRouterResponse := new(dto.OpenRouterResponse)
	if err := json.Unmarshal(body, openRouterResponse); err != nil {
		c <- &dto.SendMessageResponse{
			Status:   false,
			Response: "Failed to unmarshal OpenRouter response: " + err.Error(),
		}
		return
	}

	if len(openRouterResponse.Choices) > 0 {
		c <- &dto.SendMessageResponse{
			Status:   true,
			Response: openRouterResponse.Choices[0].Message.Content,
		}
	} else {
		c <- &dto.SendMessageResponse{
			Status:   false,
			Response: "No response choices from AI",
		}
	}
}

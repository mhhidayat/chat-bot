package aiutils

import (
	"chat_bot/configs"
	"context"
	"log"

	"google.golang.org/genai"
)

func GenerateAIResponse(ctx context.Context, prompt string, conf *configs.Config, c chan<- string) {
	client, err := genai.NewClient(ctx, &genai.ClientConfig{
		APIKey:  conf.AI.ApiKey,
		Backend: genai.BackendGeminiAPI,
	})
	if err != nil {
		log.Fatal(err)
	}

	config := &genai.GenerateContentConfig{
		SystemInstruction: genai.NewContentFromText("Jawab menggunakan bahasa indonesia", genai.RoleUser),
	}

	result, err := client.Models.GenerateContent(
		ctx,
		conf.AI.Model,
		genai.Text(prompt),
		config,
	)
	if err != nil {
		log.Fatal(err)
	}

	c <- result.Text()
}

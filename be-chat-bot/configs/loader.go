package configs

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	AI     AI
	Server Server
}

func NewConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	return &Config{
		AI: AI{
			GeminiApiKey:     os.Getenv("GEMINI_API_KEY"),
			OpenRouterApiKey: os.Getenv("OPENROUTER_API_KEY"),
			OpenRouterApiUrl: os.Getenv("OPENROUTER_API_URL"),
		},
		Server: Server{
			AllowOrigins: os.Getenv("AllowOrigins"),
		},
	}
}

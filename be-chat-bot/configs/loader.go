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
			ApiKey: os.Getenv("API_KEY"),
		},
		Server: Server{
			AllowOrigins: os.Getenv("AllowOrigins"),
		},
	}
}

package dto

import (
	"chat_bot/constants"
	"errors"
	"strings"
)

func (r *SendMessageRequest) Validate() error {
	if strings.TrimSpace(r.Prompt) == "" {
		return errors.New("prompt is required")
	}

	if strings.TrimSpace(r.Model) == "" {
		return errors.New("model is required")
	}

	if strings.TrimSpace(r.Provider) == "" {
		return errors.New("provider is required")
	}

	if r.Provider != constants.ProviderGemini && r.Provider != constants.ProviderOpenRouter {
		return errors.New("provider must be either 'gemini' or 'or'")
	}

	return nil
}

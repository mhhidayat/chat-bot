package dto

type SendMessageRequest struct {
	Prompt   string `json:"prompt"`
	Model    string `json:"model"`
	Provider string `json:"provider"`
}

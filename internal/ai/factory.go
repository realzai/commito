package ai

import (
	"fmt"
	"github.com/realzai/commito/internal/config"
)

type AIClient interface {
	Ask(prompt string, diff string) (string, error)
	Suggest(diff string) (string, error)
}

func NewClientFromConfig(cfg config.Config) (AIClient, error) {
	switch cfg.Provider {
	//case "OpenAI":
	//	return NewOpenAIClient(cfg.ApiKey, cfg.Model), nil
	case "Groq":
		return NewGroqClient(cfg.ApiKey, cfg.Model), nil
	default:
		return nil, fmt.Errorf("unsupported provider: %s", cfg.Provider)
	}
}

package utils

import (
	"fmt"
	"github.com/realzai/commito/internal/config"
)

func EnsureConfigured() (*config.Config, error) {
	cfg, err := config.LoadConfig()
	if err != nil {
		return nil, fmt.Errorf("❌ failed to load configuration: %w", err)
	}

	if cfg.Provider == "" || cfg.ApiKey == "" {
		return nil, fmt.Errorf("❌ missing provider or API key. Please run `commito setup`.")
	}

	if cfg.Provider == "Groq" && cfg.Model == "" {
		return nil, fmt.Errorf("❌ Groq model not specified. Please run `commito setup` again.")
	}

	return &cfg, nil
}

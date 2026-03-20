package ai

import (
	"context"
	"fmt"
)

func NewProvider(ctx context.Context, cfg Config) (Provider, error) {

	switch cfg.Provider {

	case "gemini":
		return NewGemini(ctx, cfg.APIKey, cfg.Model)

	case "openai":
		return NewOpenAI(cfg.APIKey, cfg.Model), nil

	default:
		return nil, fmt.Errorf("unknown AI provider: %s", cfg.Provider)

	}

}

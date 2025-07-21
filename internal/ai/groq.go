package ai

import "fmt"

type GroqClient struct {
	APIKey string
	Model  string
}

func NewGroqClient(apiKey, model string) *GroqClient {
	return &GroqClient{
		APIKey: apiKey,
		Model:  model,
	}
}

func (c *GroqClient) Ask(prompt string) (string, error) {
	return fmt.Sprintf("[Groq:%s] Here's a mock answer to: %s", c.Model, prompt), nil
}

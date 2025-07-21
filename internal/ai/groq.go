package ai

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

type GroqClient struct {
	ApiKey string
	Model  string
}

func NewGroqClient(ApiKey, model string) *GroqClient {
	return &GroqClient{
		ApiKey: ApiKey,
		Model:  model,
	}
}

type ChatMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type ChatRequest struct {
	Model    string        `json:"model"`
	Messages []ChatMessage `json:"messages"`
}

type ChatResponse struct {
	Choices []struct {
		Message ChatMessage `json:"message"`
	} `json:"choices"`
}

func (c *GroqClient) Ask(question, diff string) (string, error) {
	systemMsg := ChatMessage{
		Role:    "system",
		Content: "You are a helpful assistant that provides concise and accurate answers or commit messages based on the provided context.",
	}

	userMsg := ChatMessage{
		Role:    "user",
		Content: fmt.Sprintf("Based on the following changes:\n\n%s\n\nPlease answer the question: %s", diff, question),
	}

	resp, err := c.CreateChatCompletion([]ChatMessage{systemMsg, userMsg})
	if err != nil {
		return "", fmt.Errorf("failed to ask AI: %w", err)
	}

	return resp, nil
}

func (c *GroqClient) Suggest(diff string) (string, error) {
	systemMsg := ChatMessage{
		Role:    "system",
		Content: "You are a helpful assistant that provides concise and accurate commit messages based on the provided context.",
	}

	userMsg := ChatMessage{
		Role:    "user",
		Content: fmt.Sprintf("Based on the following changes:\n\n%s\n\nPlease suggest a commit message.", diff),
	}

	resp, err := c.CreateChatCompletion([]ChatMessage{systemMsg, userMsg})
	if err != nil {
		return "", fmt.Errorf("failed to suggest commit message: %w", err)
	}

	return resp, nil
}

func (c *GroqClient) CreateChatCompletion(messages []ChatMessage) (string, error) {
	url := "https://api.groq.com/openai/v1/chat/completions"

	body := ChatRequest{
		Model:    c.Model,
		Messages: messages,
	}

	jsonData, err := json.Marshal(body)
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return "", err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.ApiKey))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		resBody, _ := io.ReadAll(res.Body)
		return "", fmt.Errorf("Groq API error: %s", resBody)
	}

	var chatRes ChatResponse
	if err := json.NewDecoder(res.Body).Decode(&chatRes); err != nil {
		return "", err
	}

	if len(chatRes.Choices) == 0 {
		return "", errors.New("no choices returned from Groq API")
	}

	return chatRes.Choices[0].Message.Content, nil
}

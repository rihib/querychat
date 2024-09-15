package llm

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"regexp"

	"github.com/joho/godotenv"
	"github.com/rihib/querychat/internal/domain/entity"
	openai "github.com/sashabaranov/go-openai"
)

type GPT4 struct {
	apiKey string
}

func NewGPT4() (*GPT4, error) {
	if err := godotenv.Load(); err != nil {
		slog.Error("error loading .env file", "msg", err.Error())
		return nil, fmt.Errorf("error loading .env file: %v", err)
	}
	apiKey := os.Getenv("API_KEY")
	return &GPT4{apiKey: apiKey}, nil
}

func (gpt4 *GPT4) Ask(optimized entity.OptimizedPrompt) (*entity.LLMOutput, error) {
	c := openai.NewClient(gpt4.apiKey)
	resp, err := c.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT4,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleSystem,
					Content: optimized.SystemPrompt(),
				},
				{
					Role:    openai.ChatMessageRoleUser,
					Content: optimized.UserPrompt(),
				},
			},
		},
	)
	if err != nil {
		slog.Warn("openai chat completion error", "msg", err.Error())
		return nil, fmt.Errorf("openai chat completion error: %v", err)
	}
	output := resp.Choices[0].Message.Content

	var query, data string
	if query, err = extractPattern(output, "sql"); err != nil {
		return nil, err
	}
	if data, err = extractPattern(output, "json"); err != nil {
		return nil, err
	}
	return entity.NewLLMOutput(query, data)
}

func extractPattern(output string, patternType string) (string, error) {
	pattern := regexp.MustCompile(fmt.Sprintf("(?s)```%s\n(.+?)\n```", patternType))
	matches := pattern.FindStringSubmatch(output)

	if len(matches) <= 1 {
		slog.Info(fmt.Sprintf("%s not found", patternType), "output", output)
		return "", fmt.Errorf("%s not found", patternType)
	}

	return matches[1], nil
}

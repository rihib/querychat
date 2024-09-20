package llm

import (
	"context"
	"fmt"

	"github.com/rihib/querychat/internal/domain/entity"
	"github.com/rihib/querychat/internal/pkg"
	openai "github.com/sashabaranov/go-openai"
)

type GPT4 struct {
	apiKey string
}

func NewGPT4(apiKey string) (*GPT4, error) {
	if apiKey == "" {
		return nil, fmt.Errorf("api key cannot be empty")
	}
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
		return nil, fmt.Errorf("openai chat completion error: %v", err)
	}
	output := resp.Choices[0].Message.Content

	var query, chart string
	if query, err = pkg.FindPattern(output, "(?s)```sql\n(.+?)\n```"); err != nil {
		return nil, err
	}
	if chart, err = pkg.FindPattern(output, "(?s)```json\n(.+?)\n```"); err != nil {
		return nil, err
	}
	return entity.NewLLMOutput(query, chart)
}

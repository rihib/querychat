//go:generate mockgen -source=chat.go -destination=chat_mock.go -package=usecase
package usecase

import (
	"log/slog"

	"github.com/rihib/querychat/internal/domain/entity"
)

type ChatUsecase struct {
	llm  LLM
	repo ChatRepository
}

type LLM interface {
	Ask(prompt entity.OptimizedPrompt) (*entity.LLMOutput, error)
}

type ChatRepository interface {
	Exec(output entity.LLMOutput) ([]map[string]interface{}, error)
}

func NewChatUsecase(llm LLM, repo ChatRepository) *ChatUsecase {
	return &ChatUsecase{llm: llm, repo: repo}
}

func (cu *ChatUsecase) AskLLM(prompt entity.OptimizedPrompt) (*entity.LLMOutput, error) {
	slog.Info("AskLLM")
	return cu.llm.Ask(prompt)
}

/*
ExecQuery is a function that executes the query and returns the result.

Example return value:

[{"UserName": "Alice", "TotalAmount": 100}, {"UserName": "Bob", "TotalAmount": 200}]
*/
func (cu *ChatUsecase) ExecQuery(output entity.LLMOutput) ([]map[string]interface{}, error) {
	slog.Info(
		"ExecQuery",
		slog.Group(
			"LLMOutput",
			"query", output.Query(),
			"chart", output.Chart(),
		),
	)
	return cu.repo.Exec(output)
}

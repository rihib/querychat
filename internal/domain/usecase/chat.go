package usecase

import (
	"database/sql"

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
	Exec(output entity.LLMOutput) (*sql.Rows, error)
}

func NewChatUsecase(llm LLM, repo ChatRepository) *ChatUsecase {
	return &ChatUsecase{llm: llm, repo: repo}
}

func (cu *ChatUsecase) Ask(prompt entity.OptimizedPrompt) (*entity.LLMOutput, error) {
	return cu.llm.Ask(prompt)
}

func (cu *ChatUsecase) Exec(output entity.LLMOutput) (*sql.Rows, error) {
	return cu.repo.Exec(output)
}

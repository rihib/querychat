package app

import (
	"github.com/rihib/querychat/internal/domain/entity"
	"github.com/rihib/querychat/internal/domain/usecase"
)

func Chat(qcc entity.QueryChatConfig, llm usecase.LLM, repo usecase.ChatRepository) (*entity.VisualizableData, error) {
	originalPrompt, err := entity.NewOriginalPrompt(qcc.Prompt())
	if err != nil {
		return nil, err
	}

	formatPrompt, err := entity.NewFormatPrompt(qcc.SystemPrompt(), qcc.UserPrompt())
	if err != nil {
		return nil, err
	}

	optimizedPrompt, err := entity.NewOptimizedPrompt(*originalPrompt, *formatPrompt, qcc.DBName(), qcc.Schema())
	if err != nil {
		return nil, err
	}

	cu := usecase.NewChatUsecase(llm, repo)
	output, err := cu.AskLLM(*optimizedPrompt)
	if err != nil {
		return nil, err
	}

	datas, err := cu.ExecQuery(*output)
	if err != nil {
		return nil, err
	}

	vd, err := entity.NewVisualizableData(datas, *output)
	if err != nil {
		return nil, err
	}

	return vd, nil
}

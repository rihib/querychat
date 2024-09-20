package app

import (
	"github.com/rihib/querychat/internal/domain/entity"
	"github.com/rihib/querychat/internal/domain/usecase"
)

func Chat(qcConfig entity.QueryChatConfig, llm usecase.LLM, repo usecase.ChatRepository) (*entity.VisualizableData, error) {
	originalPrompt, err := entity.NewOriginalPrompt(qcConfig.Prompt())
	if err != nil {
		return nil, err
	}

	templatePrompt, err := entity.NewTemplatePrompt(qcConfig.SystemPrompt(), qcConfig.UserPrompt())
	if err != nil {
		return nil, err
	}

	optimizedPrompt, err := entity.NewOptimizedPrompt(*originalPrompt, *templatePrompt, qcConfig.DBName(), qcConfig.Schema())
	if err != nil {
		return nil, err
	}

	cu := usecase.NewChatUsecase(llm, repo)
	output, err := cu.Ask(*optimizedPrompt)
	if err != nil {
		return nil, err
	}

	datas, err := cu.Exec(*output)
	if err != nil {
		return nil, err
	}

	vd, err := entity.NewVisualizableData(datas, *output)
	if err != nil {
		return nil, err
	}

	return vd, nil
}

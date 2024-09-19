package app

import (
	"github.com/rihib/querychat/internal/config"
	"github.com/rihib/querychat/internal/domain/entity"
	"github.com/rihib/querychat/internal/domain/usecase"
)

func Chat(prompt string, llm usecase.LLM, repo usecase.ChatRepository) (*entity.VisualizableData, error) {
	originalPrompt, err := entity.NewOriginalPrompt(prompt)
	if err != nil {
		return nil, err
	}

	templatePrompt, err := entity.NewTemplatePrompt(config.SYSTEM_PROMPT, config.USER_PROMPT)
	if err != nil {
		return nil, err
	}

	optimizedPrompt, err := entity.NewOptimizedPrompt(*originalPrompt, *templatePrompt, config.DB_NAME, config.SCHEMA_FILE_PATH)
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

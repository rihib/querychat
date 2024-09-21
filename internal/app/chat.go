package app

import (
	"log/slog"

	"github.com/rihib/querychat/internal/domain/entity"
	"github.com/rihib/querychat/internal/domain/usecase"
)

/*
Chat is a function that asks LLM, gets the query and executes it, and returns the visualizable data.
*/
func Chat(cc entity.ChatConfig, llm usecase.LLM, repo usecase.ChatRepository) (*entity.VisualizableData, error) {
	optimizedPrompt, err := entity.NewOptimizedPrompt(cc)
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

	vd, err := entity.NewVisualizableData(*output, datas)
	if err != nil {
		return nil, err
	}
	slog.Info("VisualizableData", "chart", vd.Chart(), "datas", vd.Datas(), "query", vd.ExecutedQuery())

	return vd, nil
}

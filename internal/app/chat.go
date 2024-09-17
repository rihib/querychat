package app

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/rihib/querychat/internal/config"
	"github.com/rihib/querychat/internal/domain/entity"
	"github.com/rihib/querychat/internal/domain/usecase"
	"github.com/rihib/querychat/internal/gateway/llm"
	"github.com/rihib/querychat/internal/gateway/rdb"
)

func Chat(prompt string) (*entity.VisualizableData, error) {
	originalPrompt, err := entity.NewOriginalPrompt(prompt)
	if err != nil {
		return nil, err
	}

	templatePrompt, err := entity.NewTemplatePrompt(config.SYSTEM_PROMPT, config.USER_PROMPT)
	if err != nil {
		return nil, err
	}

	info, err := entity.NewUserDBInfo(config.DB_NAME, config.DB_FILE_PATH, config.SCHEMA_FILE_PATH)
	if err != nil {
		return nil, err
	}

	optimizedPrompt := entity.NewOptimizedPrompt(*originalPrompt, *templatePrompt, *info)
	fmt.Println("System Prompt:\n" + optimizedPrompt.SystemPrompt())
	fmt.Println("User Prompt:\n" + optimizedPrompt.UserPrompt())

	if err := godotenv.Load("/Users/rihib/dev/querychat/internal/config/.env"); err != nil {
		return nil, fmt.Errorf("error loading .env file: %v", err)
	}
	apiKey := os.Getenv("API_KEY")
	llm, err := llm.NewGPT4(apiKey)
	if err != nil {
		return nil, err
	}

	repo, err := rdb.NewSQLite(info)
	if err != nil {
		return nil, err
	}

	cu := usecase.NewChatUsecase(llm, repo)
	output, err := cu.Ask(*optimizedPrompt)
	if err != nil {
		return nil, err
	}
	fmt.Println("LLM Output: Query:\n" + output.Query())
	fmt.Println("LLM Output: Data:\n" + output.Data())

	rows, err := cu.Exec(*output)
	if err != nil {
		return nil, err
	}

	vd, err := entity.NewVisualizableData(rows, *output)
	if err != nil {
		return nil, err
	}
	fmt.Printf("Visualizable Data: Datas:\n%v\n", vd.Datas())
	fmt.Printf("Visualizable Data: Chart:\n%v\n", vd.Chart())

	return vd, nil
}

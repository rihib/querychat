package app

import (
	"fmt"
	"log/slog"
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
	schema, err := readFile(config.SCHEMA_FILE_PATH)
	if err != nil {
		return nil, err
	}
	info, err := entity.NewUserDBInfo(config.DB_NAME, config.DB_FILE_PATH, schema)
	if err != nil {
		return nil, err
	}
	optimizedPrompt := entity.NewOptimizedPrompt(*originalPrompt, *templatePrompt, *info)
	if err := godotenv.Load("/Users/rihib/dev/querychat/internal/config/.env"); err != nil {
		slog.Error("error loading .env file", "msg", err.Error())
		return nil, fmt.Errorf("error loading .env file: %v", err)
	}
	fmt.Println("System Prompt: " + optimizedPrompt.SystemPrompt())
	fmt.Println("User Prompt: " + optimizedPrompt.UserPrompt())
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
	fmt.Println("LLM Output: " + output.Query())
	fmt.Println("LLM Output: " + output.Data())
	rows, err := cu.Exec(*output)
	if err != nil {
		return nil, err
	}
	vd, err := entity.NewVisualizableData(rows, *output)
	if err != nil {
		return nil, err
	}
	fmt.Printf("Visualizable Data Datas: %v", vd.Datas())
	fmt.Printf("Visualizable Data Chart: %v", vd.Chart())
	return vd, nil
}

func readFile(filePath string) (string, error) {
	content, err := os.ReadFile(filePath)
	if err != nil {
		slog.Error("failed to read file", "error", err.Error())
		return "", fmt.Errorf("failed to read file: %v", err)
	}

	return string(content), nil
}

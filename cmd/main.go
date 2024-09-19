package main

import (
	"log/slog"
	"os"

	"github.com/joho/godotenv"
	"github.com/rihib/querychat/internal/app"
	"github.com/rihib/querychat/internal/config"
	"github.com/rihib/querychat/internal/gateway/llm"
	"github.com/rihib/querychat/internal/gateway/rdb"
)

const (
	PROMPT = "What are the monthly sales for 2013?"
)

func init() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)
}

func main() {
	if err := godotenv.Load(config.ENV_FILE_PATH); err != nil {
		slog.Error("error loading .env file", "msg", err.Error())
	}
	apiKey := os.Getenv("API_KEY")
	llm, err := llm.NewGPT4(apiKey)
	if err != nil {
		slog.Error("failed to create llm", "msg", err.Error())
	}

	repo, err := rdb.NewSQLite3(config.DB_NAME, config.DB_FILE_PATH)
	if err != nil {
		slog.Error("failed to create repo", "msg", err.Error())
	}

	vd, err := app.Chat(PROMPT, llm, repo)
	if err != nil {
		slog.Error("failed to chat", "msg", err.Error())
		return
	}
	slog.Info("chat successful", "vd", vd)
}

package main

import (
	"log/slog"
	"os"

	"github.com/joho/godotenv"
	"github.com/rihib/querychat/internal/app"
	"github.com/rihib/querychat/internal/domain/entity"
	"github.com/rihib/querychat/internal/gateway/llm"
	"github.com/rihib/querychat/internal/gateway/rdb"
)

func init() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)
}

func main() {
	PROMPT := os.Getenv("PROMPT")
	ENV_FILE_PATH := os.Getenv("ENV_FILE_PATH")
	if err := godotenv.Load(ENV_FILE_PATH); err != nil {
		slog.Error("error loading .env file", "msg", err.Error())
	}
	API_KEY := os.Getenv("API_KEY")
	ENVIRONMENT := os.Getenv("ENVIRONMENT")
	DB_NAME := os.Getenv("DB_NAME")
	SCHEMA_FILE_PATH := os.Getenv("SCHEMA_FILE_PATH")
	DB_FILE_PATH := os.Getenv("DB_FILE_PATH")

	switch ENVIRONMENT {
	case "local":
		// Chat Config
		schema, err := os.ReadFile(SCHEMA_FILE_PATH)
		if err != nil {
			slog.Error("failed to read schema file", "msg", err.Error())
		}
		cc, err := entity.NewChatConfig(PROMPT, DB_NAME, string(schema))
		if err != nil {
			slog.Error("failed to create query chat config", "msg", err.Error())
		}

		// LLM
		llm, err := llm.NewGPT4(API_KEY)
		if err != nil {
			slog.Error("failed to create llm", "msg", err.Error())
		}

		// Repo
		repo, err := rdb.NewSQLite3(DB_FILE_PATH)
		if err != nil {
			slog.Error("failed to create repo", "msg", err.Error())
		}

		// Chat
		_, err = app.Chat(*cc, llm, repo)
		if err != nil {
			slog.Error("failed to chat", "msg", err.Error())
		}

	case "development":

	case "staging":

	case "production":
	}
}

package main

import (
	"context"
	"log/slog"
	"os"
	"time"

	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"github.com/rihib/querychat/internal/app"
	"github.com/rihib/querychat/internal/domain/entity"
	"github.com/rihib/querychat/internal/gateway/llm"
	"github.com/rihib/querychat/internal/gateway/rdb"
)

const (
	requestTimeout = 10 * time.Second
)

func main() {
	// context
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, requestTimeout)
	defer cancel()
	ctx = context.WithValue(ctx, "requestID", uuid.New().String())

	// logger
	level := new(slog.LevelVar)
	logger := slog.New(
		slog.NewJSONHandler(
			os.Stdout,
			&slog.HandlerOptions{
				AddSource: true,
			},
		),
	)
	slog.SetDefault(logger)

	// env
	envFilePath := os.Getenv("ENV_FILE_PATH")
	if err := godotenv.Load(envFilePath); err != nil {
		slog.Error("error loading .env file", "msg", err.Error())
	}
	apiKey := os.Getenv("API_KEY")
	dbFilePath := os.Getenv("DB_FILE_PATH")
	dbName := os.Getenv("DB_NAME")
	environment := os.Getenv("ENVIRONMENT")
	prompt := os.Getenv("LLM_PROMPT")
	schemaFilePath := os.Getenv("SCHEMA_FILE_PATH")

	switch environment {
	case "local":
		level.Set(slog.LevelDebug)

		// Chat Config
		schema, err := os.ReadFile(schemaFilePath)
		if err != nil {
			slog.Error(err.Error())
		}
		cc, err := entity.NewChatConfig(prompt, dbName, string(schema))
		if err != nil {
			slog.Error(err.Error())
		}

		// LLM
		llm, err := llm.NewGPT4(apiKey)
		if err != nil {
			slog.Error(err.Error())
		}

		// Repo
		repo, err := rdb.NewSQLite3(dbFilePath)
		if err != nil {
			slog.Error(err.Error())
		}

		// Chat
		_, err = app.Chat(*cc, llm, repo)
		if err != nil {
			slog.Error(err.Error())
		}

	case "development":

	case "staging":

	case "production":
	}
}

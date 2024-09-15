package main

import (
	"log/slog"
	"os"

	"github.com/rihib/querychat/internal/app"
)

const (
	PROMPT = "What are the monthly sales for 2013?"
)

func init() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)
}

func main() {
	vd, err := app.Chat(PROMPT)
	if err != nil {
		slog.Error("failed to chat", "msg", err.Error())
		return
	}
	slog.Info("chat successful", "vd", vd)
}

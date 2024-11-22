package main

import (
	"context"
	_ "fmt"
	"log/slog"
	"os"
	"os/signal"

	"github.com/26thavenue/FXQLParser/app"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	a:= app.New(logger)
	if err := a.Start(ctx); err != nil {
		logger.Error("failed to start server", slog.Any("error", err))
	}
}
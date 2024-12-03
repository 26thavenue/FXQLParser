package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"

	"github.com/26thavenue/FXQLParser/app"
	_ "github.com/26thavenue/FXQLParser/docs"
	"github.com/joho/godotenv"
)

// @title FXQL Parser API
// @version 1.0
// @description This is a simple API that parsers FXQL statements/currency transactionxs
// @termsOfService http://swagger.io/terms/

// @contact.name Oni Oluwatomiwa

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /

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
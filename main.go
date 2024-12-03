package main

import (
	"context"
	_ "fmt"
	"log/slog"
	"os"
	"os/signal"

	"github.com/26thavenue/FXQLParser/app"
	"github.com/26thavenue/FXQLParser/database"
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

	database.Connect()

	a:= app.New(logger,database.DBInstance.Instance)

	if err := a.Start(ctx); err != nil {
		logger.Error("failed to start server", slog.Any("error", err))
	}

}
package main

import (
	// "context"
	"fmt"
	// "log/slog"
	// "os"
	// "os/signal"

	// "github.com/26thavenue/FXQLParser/app"
	"github.com/26thavenue/FXQLParser/parser"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	// logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	// ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	// defer cancel()

	// a:= app.New(logger)
	// if err := a.Start(ctx); err != nil {
	// 	logger.Error("failed to start server", slog.Any("error", err))
	// }

	input := `USD-GBP{
			BUY 100
			SELL 200
			CAP 93800
			}`

	r, err := parser.Parse(input)
	if err !=nil{
		fmt.Printf("Err %s", err)
	}

	fmt.Printf("Output %v", r)
}
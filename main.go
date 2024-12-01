package main

import (
	_ "context"
	"fmt"
	_ "log/slog"
	_ "os"
	_ "os/signal"

	_ "github.com/26thavenue/FXQLParser/app"

	"github.com/joho/godotenv"

	_ "github.com/26thavenue/FXQLParser/docs"

	"github.com/26thavenue/FXQLParser/parser"
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

	// logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	// ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	// defer cancel()

	// a:= app.New(logger)

	// if err := a.Start(ctx); err != nil {
	// 	logger.Error("failed to start server", slog.Any("error", err))
	// }

	// input := []string{
	// 	`USD-GBP {
	// 					BUY 100

	// 					SELL 200

	// 					CAP 93800
	// 					}`,
	// }

	in := `USD-GBP {
						BUY 100
						SELL 200
						CAP 93800
						}`

	vr,dd := parser.Dummy(in)

	// vr, err := parser.ProcessBlock(input)

	// if err != nil{
	// 	fmt.Printf(" Error %s", err)
	// }

	fmt.Println("%i",vr)
	fmt.Println("%i", dd)

}
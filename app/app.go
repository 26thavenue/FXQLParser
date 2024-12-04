package app

import (
	"context"
	"errors"
	"log/slog"
	"net/http"
	"time"

	_ "github.com/26thavenue/FXQLParser/docs"
	middlewares "github.com/26thavenue/FXQLParser/middleware"
	"gorm.io/gorm"
)

type App struct {
	logger *slog.Logger
	router *http.ServeMux
	db  *gorm.DB
}

func New(logger *slog.Logger, db *gorm.DB) *App{
	router := http.NewServeMux()

	app := &App{
		logger: logger,
		router: router,
		db:db,
	}

	return app
}

func (a *App)Start (ctx context.Context) error{
	
	server := http.Server{
		Addr:    ":8080",
		Handler: middlewares.Logger(a.logger ,a.router ),
	}

	a.loadRoutes()

	done := make(chan struct{})
	go func() {
		err := server.ListenAndServe()
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			a.logger.Error("failed to listen and serve", slog.Any("error", err))
		}
		close(done)
	}()

	a.logger.Info("Server listening", slog.String("addr", ":8080"))
	select {
	case <-done:
		break
	case <-ctx.Done():
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
		server.Shutdown(ctx)
		cancel()
	}

	return nil
}
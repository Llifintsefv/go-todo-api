package main

import (
	"context"
	"fmt"
	"go-todo-api/internal/config"
	"go-todo-api/internal/handler"
	"go-todo-api/internal/repository/postgres"
	"go-todo-api/internal/router"
	"go-todo-api/internal/service"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stderr, nil))

	slog.SetDefault(logger)

	cfg, err := config.NewConfig()

	if err != nil {
		slog.Error("failed to load config", "error", err)
		os.Exit(1)
	}

	db, err := postgres.NewDB(cfg.DBConnStr)

	if err != nil {
		slog.Error("failed to connect to database", "error", err)
		os.Exit(1)
	}

	repository := postgres.NewTaskRepository(db, logger)
	service := service.NewTaskService(repository, logger)
	handler := handler.NewTaskHandler(service, logger)

	app := router.SetupRouter(handler)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		if err := app.Listen(cfg.Port); err != nil && err != http.ErrServerClosed {
			slog.Error("failed to start server", "error", err)
			os.Exit(1)
		}
	}()

	slog.Info(fmt.Sprintf("Server started on port %s", cfg.Port))
	<-quit
	slog.Info("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	defer db.Close(ctx)

	if err := app.ShutdownWithContext(ctx); err != nil {
		slog.Error("Server forced to shutdown: ", "error", err)
		os.Exit(1)
	}

	slog.Info("Server exiting")
}

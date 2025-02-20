package handler

import (
	"go-todo-api/internal/service"
	"log/slog"
)

type TaskHandler interface {
}

type taskHandler struct {
	service service.TaskService
	logger  *slog.Logger
}

func NewTaskHandler(service service.TaskService, logger *slog.Logger) TaskHandler {
	return &taskHandler{
		service: service,
		logger:  logger,
	}
}

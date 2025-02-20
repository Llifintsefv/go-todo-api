package service

import (
	"go-todo-api/internal/repository"
	"log/slog"
)

type TaskService interface {
}

type taskService struct {
	repo   repository.TaskRepository
	logger *slog.Logger
}

func NewTaskService(repo repository.TaskRepository, logger *slog.Logger) TaskService {
	return &taskService{
		repo:   repo,
		logger: logger,
	}
}


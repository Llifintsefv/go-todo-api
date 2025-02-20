package service

import (
	"context"
	"go-todo-api/internal/models"
	"go-todo-api/internal/repository"
	"log/slog"
)

type TaskService interface {
	CreateTask(ctx context.Context, CreateRequest models.Task) (models.Task, error)
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

func (s *taskService) CreateTask(ctx context.Context, CreateRequest models.Task) (models.Task, error) {

	CreateRequest.Status = "new"

	task, err := s.repo.CreateTask(ctx, CreateRequest)
	if err != nil {
		s.logger.Error("Error creating task", slog.Any("error", err))
		return models.Task{}, err
	}

	return task, nil
}

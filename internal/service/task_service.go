package service

import (
	"context"
	"go-todo-api/internal/models"
	"go-todo-api/internal/repository"
	"log/slog"
)

type TaskService interface {
	CreateTask(ctx context.Context, CreateRequest models.Task) (models.Task, error)
	GetTasks(ctx context.Context) ([]models.Task, error)
	UpdateTask(ctx context.Context, updateRequest models.Task) (models.Task, error)
	DeleteTask(ctx context.Context, taskID int) error
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

	CreateRequest.Status = models.StatusNew

	task, err := s.repo.CreateTask(ctx, CreateRequest)
	if err != nil {
		s.logger.Error("Error creating task", slog.Any("error", err))
		return models.Task{}, err
	}

	return task, nil
}

func (s *taskService) GetTasks(ctx context.Context) ([]models.Task, error) {
	return s.repo.GetTasks(ctx)
}

func (s *taskService) UpdateTask(ctx context.Context, updateRequest models.Task) (models.Task, error) {
	return s.repo.UpdateTask(ctx, updateRequest)
}

func (s *taskService) DeleteTask(ctx context.Context, taskID int) error {
	return s.repo.DeleteTask(ctx, taskID)
}

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
	GetTask(ctx context.Context, taskId int) (models.Task, error)
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

func (s *taskService) GetTask(ctx context.Context, taskId int) (models.Task, error) {
	return s.repo.GetTask(ctx, taskId)
}

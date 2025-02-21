package repository

import (
	"context"
	"go-todo-api/internal/models"
)

type TaskRepository interface {
	CreateTask(ctx context.Context, CreateTask models.Task) (models.Task, error)
	GetTasks(ctx context.Context) ([]models.Task, error)
	UpdateTask(ctx context.Context, updateRequest models.Task) (models.Task, error)
	DeleteTask(ctx context.Context, taskID int) error
}

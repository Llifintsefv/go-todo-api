package repository

import (
	"context"
	"go-todo-api/internal/models"
)

type TaskRepository interface {
	CreateTask(ctx context.Context, CreateTask models.Task) (models.Task, error)
	GetTasks(ctx context.Context) ([]models.Task, error)
	GetTask(ctx context.Context, taskId int) (models.Task, error)
}

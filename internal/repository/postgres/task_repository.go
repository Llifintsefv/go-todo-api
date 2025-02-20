package postgres

import (
	"context"
	"go-todo-api/internal/models"
	"go-todo-api/internal/repository"
	"log/slog"

	"github.com/jackc/pgx/v5"
)

type taskRepository struct {
	db     *pgx.Conn
	logger *slog.Logger
}

func NewTaskRepository(db *pgx.Conn, logger *slog.Logger) repository.TaskRepository {
	return &taskRepository{
		db:     db,
		logger: logger,
	}
}

func (r *taskRepository) CreateTask(ctx context.Context, CreateRequest models.Task) (models.Task, error) {
	const query = `INSERT INTO tasks (title, description,status) VALUES ($1, $2,$3) RETURNING id, title, description, status, created_at, updated_at;`
	var task models.Task
	err := r.db.QueryRow(ctx, query, CreateRequest.Title, CreateRequest.Description, CreateRequest.Status).Scan(&task.ID, &task.Title, &task.Description, &task.Status, &task.CreatedAt, &task.UpdatedAt)
	if err != nil {
		slog.Error("Error create task", slog.Any("error", err))
		return models.Task{}, err
	}

	return task, nil
}

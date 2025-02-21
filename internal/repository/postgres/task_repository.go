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

func (r *taskRepository) GetTasks(ctx context.Context) ([]models.Task, error) {
	const query = `SELECT * FROM tasks`
	var tasks []models.Task
	rows, err := r.db.Query(ctx, query)
	if err != nil {
		slog.Error("Error get tasks", slog.Any("error", err))
		return []models.Task{}, err
	}

	for rows.Next() {
		var task models.Task
		if err := rows.Scan(&task.ID, &task.Description, &task.Title, &task.Status, &task.CreatedAt, &task.UpdatedAt); err != nil {
			slog.Error("Error get task", "error", err)
			return []models.Task{}, err
		}
		tasks = append(tasks, task)
	}

	return tasks, nil
}

func (r *taskRepository) UpdateTask(ctx context.Context, updateRequest models.Task) (models.Task, error) {
	const query = "UPDATE tasks SET title = $1, description = $2, status = $3, updated_at = CURRENT_TIMESTAMP WHERE id = $4 RETURNING id, title, description, status, created_at, updated_at;"
	var task models.Task
	err := r.db.QueryRow(ctx, query, updateRequest.Title, updateRequest.Description, updateRequest.Status, updateRequest.ID).Scan(&task.ID, &task.Title, &task.Description, &task.Status, &task.CreatedAt, &task.UpdatedAt)
	if err == pgx.ErrNoRows {
		return models.Task{}, models.ErrTaskNotFound
	} else if err != nil {
		r.logger.Error("Error updating task", slog.Any("error", err))
		return models.Task{}, err
	}
	return task, nil
}

func (r *taskRepository) DeleteTask(ctx context.Context, taskID int) error {
	const query = "DELETE FROM tasks WHERE id = $1;"
	result, err := r.db.Exec(ctx, query, taskID)
	if err != nil {
		r.logger.Error("Error deleting task", slog.Any("error", err))
		return err
	}
	if result.RowsAffected() == 0 {
		return models.ErrTaskNotFound
	}
	return nil
}

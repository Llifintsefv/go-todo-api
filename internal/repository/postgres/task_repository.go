package postgres

import (
	"database/sql"
	"go-todo-api/internal/repository"
	"log/slog"

	_ "github.com/lib/pq"
)

type taskRepository struct {
	db     *sql.DB
	logger *slog.Logger
}

func NewTaskRepository(db *sql.DB, logger *slog.Logger) repository.TaskRepository {
	return &taskRepository{
		db:     db,
		logger: logger,
	}
}

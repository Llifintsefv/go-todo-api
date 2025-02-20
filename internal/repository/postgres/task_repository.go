package postgres

import (
	"go-todo-api/internal/repository"
	"log/slog"

	"github.com/jackc/pgx/v5"
	_ "github.com/lib/pq"
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

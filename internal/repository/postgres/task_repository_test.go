package postgres

import (
	"context"
	"testing"
	"time"

	"go-todo-api/internal/models"
	"log/slog"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/stretchr/testify/assert"
)

type testDB struct {
	conn *pgx.Conn
}

func setupTestDB(t *testing.T) *testDB {
	connStr := "postgres://postgres:password@localhost:5433/test_db?sslmode=disable"

	conn, err := pgx.Connect(context.Background(), connStr)
	if err != nil {
		t.Fatalf("failed to connect to test database: %v", err)
	}
	_, err = conn.Exec(context.Background(), `
		CREATE TABLE IF NOT EXISTS tasks (
			id SERIAL PRIMARY KEY,
			title VARCHAR(255) NOT NULL,
			description TEXT,
			status VARCHAR(50) NOT NULL,
			created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
			updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
		)
	`)
	if err != nil {
		t.Fatalf("failed to create table: %v", err)
	}

	return &testDB{conn: conn}
}

func (tdb *testDB) cleanup(t *testing.T) {
	_, err := tdb.conn.Exec(context.Background(), "TRUNCATE TABLE tasks RESTART IDENTITY")
	if err != nil {
		t.Fatalf("failed to cleanup database: %v", err)
	}
	tdb.conn.Close(context.Background())
}

func TestTaskRepository_CreateTask(t *testing.T) {
	testDB := setupTestDB(t)
	defer testDB.cleanup(t)

	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	repo := NewTaskRepository(testDB.conn, logger)

	tests := []struct {
		name         string
		inputTask    models.Task
		wantErr      bool
		validateTask func(t *testing.T, task models.Task)
	}{
		{
			name: "Successful task creation",
			inputTask: models.Task{
				Title:       "Test Task",
				Description: "Test Description",
				Status:      "TODO",
			},
			wantErr: false,
			validateTask: func(t *testing.T, task models.Task) {
				assert.NotZero(t, task.ID, "task ID should be set")
				assert.Equal(t, "Test Task", task.Title)
				assert.Equal(t, "Test Description", task.Description)
				assert.Equal(t, "TODO", task.Status)
				assert.WithinDuration(t, time.Now(), task.CreatedAt, time.Second)
				assert.WithinDuration(t, time.Now(), task.UpdatedAt, time.Second)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()

			gotTask, err := repo.CreateTask(ctx, tt.inputTask)

			if tt.wantErr {
				assert.Error(t, err)
				return
			}
			assert.NoError(t, err)

			tt.validateTask(t, gotTask)

			var count int
			err = testDB.conn.QueryRow(ctx, "SELECT COUNT(*) FROM tasks WHERE title = $1", tt.inputTask.Title).Scan(&count)
			assert.NoError(t, err)
			assert.Equal(t, 1, count)
		})
	}
}

package postgres

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

func NewDB(connStr string) (*pgx.Conn, error) {
	pool, err := pgx.Connect(context.Background(), connStr)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}
	if err := pool.Ping(context.Background()); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}
	return pool, nil
}

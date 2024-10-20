package postgres

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
)

// ConnectDB connects to the database and returns a connection pool
func ConnectDB(connStr string) (*pgxpool.Pool, error) {
	pool, err := pgxpool.New(context.Background(), connStr)
	if err != nil {
		return nil, err
	}
	return pool, nil
}

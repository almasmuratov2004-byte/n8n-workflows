package db

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
)

func Connect(url string) (*pgxpool.Pool, error) {
	pool, err := pgxpool.New(context.Background(), url)
	return pool, err
}

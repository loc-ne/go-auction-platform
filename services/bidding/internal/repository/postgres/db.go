package postgres

import (
    "context"
    "os"
    "github.com/jackc/pgx/v5/pgxpool" 
)

type PostgresDB struct {
    Pool *pgxpool.Pool
}

func NewPostgresDB(ctx context.Context) (*PostgresDB, error) {
    connStr := os.Getenv("DATABASE_URL")
    
    config, err := pgxpool.ParseConfig(connStr)
    if err != nil {
        return nil, err
    }

    config.MaxConns = 25 
    
    pool, err := pgxpool.NewWithConfig(ctx, config)
    if err != nil {
        return nil, err
    }

    if err := pool.Ping(ctx); err != nil {
        return nil, err
    }

    return &PostgresDB{Pool: pool}, nil
}
package db

import (
	"context"
	"fmt"
	"github.com/ellioht/eth-gas/config"
	"github.com/jackc/pgx/v5/pgxpool"
)

func ConnectPostgres(cfg config.Database) (*pgxpool.Pool, error) {
	ctx := context.Background()

	connString := fmt.Sprintf("postgres://%s:%s@%s:%d/%s", cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.Database)
	connConfig, err := pgxpool.ParseConfig(connString)
	if err != nil {
		return nil, fmt.Errorf("unable to parse database config: %w", err)
	}

	dbPool, err := pgxpool.NewWithConfig(ctx, connConfig)
	if err != nil {
		return nil, fmt.Errorf("unable to connect to database: %w", err)
	}

	conn, err := dbPool.Acquire(context.Background())
	if err != nil {
		return nil, fmt.Errorf("failed to acquire connection: %w", err)
	}
	defer conn.Release()

	if err := conn.Ping(context.Background()); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	return dbPool, nil
}

package db

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/tern/v2/migrate"
	"os"
)

func RunDBMigrations(ctx context.Context, pool *pgxpool.Pool) error {
	conn, err := pool.Acquire(ctx)
	if err != nil {
		return fmt.Errorf("error acquiring connection: %w", err)
	}
	defer conn.Release()

	ternMigrator, err := migrate.NewMigrator(ctx, conn.Conn(), "tern")
	if err != nil {
		return fmt.Errorf("error creating migrator: %w", err)
	}

	path := os.DirFS("./database/migrations")

	if err := ternMigrator.LoadMigrations(path); err != nil {
		return fmt.Errorf("error loading migrations: %w", err)
	}

	if err := ternMigrator.Migrate(ctx); err != nil {
		return fmt.Errorf("error running migrations: %w", err)
	}

	return nil
}

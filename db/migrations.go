package db

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/tern/v2/migrate"
	"log"
	"os"
)

func RunDBMigrations(ctx context.Context, pool *pgxpool.Pool) error {
	conn, err := pool.Acquire(ctx)
	if err != nil {
		return fmt.Errorf("error acquiring connection: %w", err)
	}
	defer conn.Release()

	ternMigrator, err := migrate.NewMigrator(ctx, conn.Conn(), "schema_version")
	if err != nil {
		return fmt.Errorf("error creating migrator: %w", err)
	}

	ternMigrator.OnStart = func(seq int32, name string, directionName string, sql string) {
		log.Println(fmt.Sprintf("Starting migration %d/%d: %s", seq, len(ternMigrator.Migrations), name))
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

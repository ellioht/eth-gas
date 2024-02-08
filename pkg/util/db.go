package util

import (
	"context"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
)

func TxBegin(ctx context.Context, pool *pgxpool.Pool) (pgx.Tx, error) {
	tx, err := pool.Begin(ctx)
	if err != nil {
		return nil, err
	}
	return tx, nil
}

func TxCommit(ctx context.Context, tx pgx.Tx) error {
	err := tx.Commit(ctx)
	if err != nil {
		return err
	}
	return nil
}

func TxRollback(tx pgx.Tx, ctx context.Context) {
	err := tx.Rollback(ctx)
	if err != nil {
		if err != pgx.ErrTxClosed {
			log.Fatalf("Transaction already closed: %v", err)
		}
	}
}

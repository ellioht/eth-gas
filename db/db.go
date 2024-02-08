package db

import (
	"context"
	"fmt"
	"github.com/ellioht/eth-gas/pkg/util"
	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/jackc/pgx/v5/pgxpool"
	"time"
)

type Database struct {
	pool *pgxpool.Pool
}

type GasPriceRecord struct {
	Timestamp time.Time
	GasPrice  int64 // Gas price in Wei
}

func NewDatabase(pool *pgxpool.Pool) *Database {
	return &Database{pool}
}

func (db *Database) SaveGasPrice(ctx context.Context, record GasPriceRecord) error {
	query := `INSERT INTO gas_prices (timestamp, gas_price) VALUES ($1, $2)`

	tx, err := util.TxBegin(ctx, db.pool)
	if err != nil {
		return fmt.Errorf("error beginning transaction: %w", err)
	}
	defer util.TxRollback(tx, ctx)

	_, err = tx.Exec(ctx, query, record.Timestamp, record.GasPrice)
	if err != nil {
		return fmt.Errorf("error inserting gas price record: %w", err)
	}

	if err = util.TxCommit(ctx, tx); err != nil {
		return fmt.Errorf("error committing transaction: %w", err)
	}

	return nil
}

func (db *Database) RetrieveGasPrices(ctx context.Context, start, end time.Time) ([]GasPriceRecord, error) {
	query := `SELECT timestamp, gas_price FROM gas_prices WHERE timestamp BETWEEN $1 AND $2`

	tx, err := util.TxBegin(ctx, db.pool)
	if err != nil {
		return nil, fmt.Errorf("error beginning transaction: %w", err)
	}
	defer util.TxRollback(tx, ctx)

	var records []GasPriceRecord
	err = pgxscan.Select(ctx, tx, &records, query, start, end)
	if err != nil {
		return nil, fmt.Errorf("error retrieving gas price records: %w", err)
	}

	if err = util.TxCommit(ctx, tx); err != nil {
		return nil, fmt.Errorf("error committing transaction: %w", err)
	}

	return records, nil
}

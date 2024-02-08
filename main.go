package main

import (
	"context"
	"github.com/ellioht/eth-gas/api"
	"github.com/ellioht/eth-gas/config"
	database "github.com/ellioht/eth-gas/db"
	"github.com/ellioht/eth-gas/ethclient"
	"log"
)

func main() {
	ctx := context.Background()

	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Error loading configuration: %v", err)
	}

	pool, err := database.ConnectPostgres(cfg.Database)
	if err != nil {
		log.Fatalf("Error initializing database: %v", err)
	}
	defer pool.Close()

	if err := database.RunDBMigrations(ctx, pool); err != nil {
		log.Fatalf("Error running database migrations: %v", err)
	}

	db := database.NewDatabase(pool)

	ethClient, err := ethclient.NewClient(cfg.Ethereum)
	if err != nil {
		log.Fatalf("Error initializing Ethereum client: %v", err)
	}

	go ethClient.StartMonitoringGasPrices(db)

	server := api.NewServer(cfg.API, db)
	if err := server.Start("80"); err != nil {
		log.Fatalf("Error starting API server: %v", err)
	}
}

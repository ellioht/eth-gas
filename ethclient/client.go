package ethclient

import (
	"context"
	"fmt"
	"github.com/ellioht/eth-gas/config"
	database "github.com/ellioht/eth-gas/db"
	"log"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
)

type Client struct {
	ethClient *ethclient.Client
	config    config.Ethereum
}

func NewClient(cfg config.Ethereum) (*Client, error) {
	url := fmt.Sprintf("%s/%s", cfg.NodeURL, cfg.InfuraKey)
	rpcClient, err := rpc.Dial(url)
	if err != nil {
		return nil, err
	}

	ethClient := ethclient.NewClient(rpcClient)

	return &Client{
		ethClient: ethClient,
		config:    cfg,
	}, nil
}

func (c *Client) StartMonitoringGasPrices(db *database.Database) {
	ticker := time.NewTicker(c.config.PollingInterval)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			price, err := c.fetchCurrentGasPrice()
			if err != nil {
				log.Printf("Error fetching gas price: %s", err)
				continue
			}

			record := database.GasPriceRecord{
				Timestamp: time.Now(),
				GasPrice:  price.Int64(),
			}

			if err := db.SaveGasPrice(context.Background(), record); err != nil {
				log.Printf("Error saving gas price record: %s", err)
			}
		}
	}
}

func (c *Client) fetchCurrentGasPrice() (*big.Int, error) {
	return c.ethClient.SuggestGasPrice(context.Background())
}

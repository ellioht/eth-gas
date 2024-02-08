package config

import "time"

type Config struct {
	Ethereum Ethereum
	Database Database
	API      API
}

type Ethereum struct {
	NodeURL           string        // URL of the Ethereum node
	PollingInterval   time.Duration // Interval for polling gas prices
	HistoricalDataCap int           // Max number of historical records to keep
}

type Database struct {
	PoolSize int
	User     string
	Password string
	Host     string
	Port     int
	Database string
}

type API struct {
	Port string
}

package config

import "time"

type Config struct {
	Ethereum Ethereum
	Database Database
	API      API
}

type Ethereum struct {
	NodeURL           string
	InfuraKey         string
	PollingInterval   time.Duration
	HistoricalDataCap int
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

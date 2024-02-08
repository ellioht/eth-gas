package config

import (
	"github.com/spf13/viper"
	"time"
)

func init() {
	viper.SetDefault("ethereum.nodeurl", "https://mainnet.infura.io/v3/{API_KEY}")
	viper.SetDefault("ethereum.pollinginterval", 30*time.Second)
	viper.SetDefault("ethereum.historicaldatacap", 1000)

	viper.SetDefault("database.poolsize", 5)
	viper.SetDefault("database.host", "localhost")
	viper.SetDefault("database.user", "ethgas")
	viper.SetDefault("database.password", "ethgas_pass")
	viper.SetDefault("database.database", "ethgas")
	viper.SetDefault("database.port", 5434)

	viper.SetDefault("api.port", "80")
}

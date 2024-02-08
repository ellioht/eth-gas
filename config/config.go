package config

import (
	"errors"
	"github.com/spf13/viper"
	"log"
	"strings"
)

func LoadConfig() (*Config, error) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("/etc/appname")
	viper.AddConfigPath("$HOME/.appname")
	viper.AddConfigPath(".")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()
	viper.AllowEmptyEnv(true)

	if err := viper.ReadInConfig(); err != nil {
		var configFileNotFoundError viper.ConfigFileNotFoundError
		if errors.As(err, &configFileNotFoundError) {
			log.Printf("Config file not found; using default values")
		}
	}

	var config Config
	err := viper.Unmarshal(&config)
	return &config, err
}

package config

import (
	"strings"

	"github.com/spf13/viper"
)

const (
	defaultHttpPort = "80"
)

type (
	Config struct {
		HTTP  HTTPConfig
		MySQL MySQLConfig
	}

	HTTPConfig struct {
		Port string
	}

	MySQLConfig struct {
		URI          string
		DatabaseName string
		User         string
		Password     string
	}
)

func Init(path string) (*Config, error) {
	populateDefaults()

	if err := parseConfigFile(path); err != nil {
		return nil, err
	}

	var cfg Config
	if err := unmarshal(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}

func unmarshal(cfg *Config) error {
	if err := viper.UnmarshalKey("http", &cfg.HTTP); err != nil {
		return err
	}

	if err := viper.UnmarshalKey("mysql", &cfg.MySQL); err != nil {
		return err
	}

	return nil
}

func parseConfigFile(filepath string) error {
	path := strings.Split(filepath, "/")
	viper.AddConfigPath(path[0])
	viper.SetConfigName(path[1])

	return viper.ReadInConfig()
}

func populateDefaults() {
	viper.SetDefault("http.port", defaultHttpPort)
}

package config

import (
	"log"
	"os"

	"github.com/spf13/viper"
)

type Config struct {
	App AppConfig
}

type AppConfig struct {
	Port        string
	Timeout     int
	IdleTimeout int
	WaitTimeout int
}

func NewConfig(name string) (*Config, error) {
	var cfg Config

	// Read the configuration directory from the environment variable, fallback to the default value
	cfgDir := os.Getenv("CONFIG_DIR")
	if cfgDir == "" {
		cfgDir = "./config"
	}

	// Set the configuration file name and read in the environment variables
	viper.AddConfigPath(cfgDir)
	viper.SetConfigName(name)
	viper.SetConfigType("yml")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Println("Config file not found...")
		return nil, err
	}

	if err := viper.Unmarshal(&cfg); err != nil {
		log.Println("Failed to unmarshal config...")
		return nil, err
	}

	return &cfg, nil
}

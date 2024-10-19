package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	App AppConfig
}

type AppConfig struct {
	Name    string
	Port    int
	Timeout int
	Idle    int
	Wait    int
}

func NewConfig(name string) (*Config, error) {
	var cfg Config

	viper.AddConfigPath("./config")
	viper.SetConfigName(name)
	viper.SetConfigType("yml")

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

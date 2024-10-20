package config

import (
	"log"
	"os"

	"github.com/spf13/viper"
)

// Config holds the application configuration.
type Config struct {
	App   AppConfig
	DB    DBConfig
	Kafka KafkaConfig
}

// AppConfig holds the application configuration.
type AppConfig struct {
	Port        string `mapstructure:"port"`
	Timeout     int    `mapstructure:"timeout"`
	IdleTimeout int    `mapstructure:"idle_timeout"`
	WaitTimeout int    `mapstructure:"wait_timeout"`
}

// DBConfig holds the database configuration.
type DBConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
	SSLMode  string
}

// KafkaConfig holds the Kafka configuration.
type KafkaConfig struct {
	Brokers      string
	ChanSize     int
	FlushTimeout int
	Topic        string
}

func (db *DBConfig) GetDSN() string {
	return "host=" + db.Host + " port=" + db.Port + " user=" + db.User + " password=" + db.Password + " dbname=" +
		db.Name + " sslmode=" + db.SSLMode
}

// NewConfig reads the configuration file and environment variables.
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
	// Ensure that the environment variables are mapped to the configuration fields
	viper.AutomaticEnv()

	// Bind environment variables
	BindEnv("db.host", "POSTGRES_HOST")
	BindEnv("db.port", "POSTGRES_PORT")
	BindEnv("db.user", "POSTGRES_USER")
	BindEnv("db.password", "POSTGRES_PASSWORD")
	BindEnv("db.name", "POSTGRES_DBNAME")
	BindEnv("db.sslmode", "POSTGRES_SSLMODE")

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

func BindEnv(key string, val string) {
	if err := viper.BindEnv(key, val); err != nil {
		log.Println("Failed to bind environment variable: ", key)
	}
}

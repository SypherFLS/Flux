package config

//Егор: добавил конфиг, написал функцию для загрузки конфигурации из файла

import (
	"log"
	"os"
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	Env   string      `yaml:"env"`
	HTTP  HTTPConfig  `yaml:"http"`
	DB    DBConfig    `yaml:"db"`
	Kafka KafkaConfig `yaml:"kafka"`
	JWT   JWTConfig   `yaml:"jwt"`
}

type HTTPConfig struct {
	Address     string        `yaml:"address"`
	Timeout     time.Duration `yaml:"timeout"`
	IdleTimeout time.Duration `yaml:"idle_timeout"`
}

type DBConfig struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	DBName   string `yaml:"dbname"`
	SSLMode  string `yaml:"sslmode"`
}

type KafkaConfig struct {
	Broker string `yaml:"broker"`
	Topic  string `yaml:"topic"`
}

type JWTConfig struct {
	SecretKey string        `yaml:"secret_key"`
	AccessTTL time.Duration `yaml:"access_ttl"`
}

func MustLoad() *Config {
	configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
		log.Fatal("CONFIG_PATH is not set")
	}

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("config file does not exist: %s", configPath)
	}

	viper.SetConfigFile(configPath)
	viper.SetConfigType("yaml")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("error reading config file: %s", err)
	}

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		log.Fatalf("error unmarshaling config: %s", err)
	}

	applyDefaults(&cfg)
	return &cfg
}

// Егор: добавил функцию для применения значений по умолчанию, если они не указаны в конфиге
func applyDefaults(cfg *Config) {
	if cfg.Env == "" {
		cfg.Env = "local"
	}
	if cfg.HTTP.Address == "" {
		cfg.HTTP.Address = "localhost:8080"
	}
	if cfg.HTTP.Timeout == 0 {
		cfg.HTTP.Timeout = 30 * time.Second
	}
	if cfg.HTTP.IdleTimeout == 0 {
		cfg.HTTP.IdleTimeout = 60 * time.Second
	}
	if cfg.DB.SSLMode == "" {
		cfg.DB.SSLMode = "disable"
	}
	if cfg.DB.Port == 0 {
		cfg.DB.Port = 5432
	}
	if cfg.Kafka.Topic == "" {
		cfg.Kafka.Topic = "logs"
	}
	if cfg.JWT.AccessTTL == 0 {
		cfg.JWT.AccessTTL = 24 * time.Hour
	}
}

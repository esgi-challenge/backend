package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port    string `env:"API_PORT"`
	BaseUrl string `env:"BASE_URL"`

	Postgres PostgresConfig
}

type PostgresConfig struct {
	Host     string `env:"PG_HOST"`
	Port     string `env:"PG_PORT"`
	User     string `env:"PG_USER"`
	Password string `env:"PG_PASSWORD"`
	Dbname   string `env:"PG_DBNAME"`
}

func LoadConfig(filePath string, env string) (*Config, error) {
	if env == "LOCAL" {
		if _, err := os.Stat(filePath); err != nil {
			return nil, err
		}

		err := godotenv.Load(filePath)
		if err != nil {
			return nil, err
		}
	}

	return &Config{
		Port:    os.Getenv("API_PORT"),
		BaseUrl: os.Getenv("BASE_URL"),
		Postgres: PostgresConfig{
			Host:     os.Getenv("PG_HOST"),
			Port:     os.Getenv("PG_PORT"),
			User:     os.Getenv("PG_USER"),
			Password: os.Getenv("PG_PASSWORD"),
			Dbname:   os.Getenv("PG_DBNAME"),
		},
	}, nil
}

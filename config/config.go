package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
  Port  string `env:"API_PORT"`
  BaseUrl string `env:"BASE_URL"`
}

func LoadConfig(filePath string, env string) (*Config, error) {
  if env == "DEV" {
    if _, err := os.Stat(filePath); err != nil {
      return nil, err
    }

    err := godotenv.Load(filePath)
    if err != nil {
      return nil, err
    }
  }

  var c Config

  c.Port = os.Getenv("API_PORT")
  c.BaseUrl = os.Getenv("BASE_URL")

  return &c, nil
}

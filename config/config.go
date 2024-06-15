package config

import (
	"errors"
	"os"
	"reflect"

	"github.com/joho/godotenv"
)

type Config struct {
	Port    string `env:"API_PORT"`
	BaseUrl string `env:"BASE_URL"`

	JwtSecret string `env:"JWT_SECRET"`
	Postgres  PostgresConfig
	Smtp      SMTPConfig
}

type PostgresConfig struct {
	Host     string `env:"PG_HOST"`
	Port     string `env:"PG_PORT"`
	User     string `env:"PG_USER"`
	Password string `env:"PG_PASSWORD"`
	Dbname   string `env:"PG_DBNAME"`
}

type SMTPConfig struct {
	Username string `env:"SMTP_USERNAME"`
	Password string `env:"SMTP_PASSWORD"`
	Host     string `env:"SMTP_HOST"`
}

func hasEmptyFields(v interface{}) bool {
	val := reflect.ValueOf(v).Elem()

	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)

		if isEmptyField(field) {
			return true
		}
	}

	return false
}

func isEmptyField(v reflect.Value) bool {
	switch v.Kind() {
	case reflect.String:
		return v.String() == ""
	case reflect.Struct:
		for i := 0; i < v.NumField(); i++ {
			if isEmptyField(v.Field(i)) {
				return true
			}
		}
	}
	return false
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

	config := &Config{
		Port:      os.Getenv("API_PORT"),
		BaseUrl:   os.Getenv("BASE_URL"),
		JwtSecret: os.Getenv("JWT_SECRET"),
		Postgres: PostgresConfig{
			Host:     os.Getenv("PG_HOST"),
			Port:     os.Getenv("PG_PORT"),
			User:     os.Getenv("PG_USER"),
			Password: os.Getenv("PG_PASSWORD"),
			Dbname:   os.Getenv("PG_DBNAME"),
		},
		Smtp: SMTPConfig{
			Username: os.Getenv("SMTP_USERNAME"),
			Password: os.Getenv("SMTP_PASSWORD"),
			Host:     os.Getenv("SMTP_HOST"),
		},
	}

	if hasEmptyFields(config) {
		return nil, errors.New("Some env variables are not set, API won't work without them.")
	}

	return config, nil
}

package config

import (
	"fmt"
	"os"
)

type Config struct {
	HTTPAddr    string
	DatabaseURL string
}

func Load() (Config, error) {
	databaseURL, err := databaseURL()
	if err != nil {
		return Config{}, err
	}

	return Config{
		HTTPAddr:    envOr("HTTP_ADDR", ":8080"),
		DatabaseURL: databaseURL,
	}, nil
}

func databaseURL() (string, error) {
	if url := os.Getenv("DATABASE_URL"); url != "" {
		return url, nil
	}

	user, err := envRequired("POSTGRES_USER")
	if err != nil {
		return "", err
	}
	password, err := envRequired("POSTGRES_PASSWORD")
	if err != nil {
		return "", err
	}
	dbName, err := envRequired("POSTGRES_DB")
	if err != nil {
		return "", err
	}

	host := envOr("POSTGRES_HOST", "localhost")
	port := envOr("POSTGRES_PORT", "5432")

	return fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		user, password, host, port, dbName,
	), nil
}

func envRequired(key string) (string, error) {
	value := os.Getenv(key)
	if value == "" {
		return "", fmt.Errorf("required environment variable %s is not set", key)
	}
	return value, nil
}

func envOr(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}

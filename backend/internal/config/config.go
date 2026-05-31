package config

import (
	"fmt"
	"os"
)

type Config struct {
	HTTPAddr    string
	DatabaseURL string
}

func Load() Config {
	return Config{
		HTTPAddr:    envOr("HTTP_ADDR", ":8080"),
		DatabaseURL: databaseURL(),
	}
}

func databaseURL() string {
	if url := os.Getenv("DATABASE_URL"); url != "" {
		return url
	}

	user := envOr("POSTGRES_USER", "record")
	password := envOr("POSTGRES_PASSWORD", "pass")
	host := envOr("POSTGRES_HOST", "localhost")
	port := envOr("POSTGRES_PORT", "5432")
	dbName := envOr("POSTGRES_DB", "record")

	return fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		user, password, host, port, dbName,
	)
}

func envOr(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}

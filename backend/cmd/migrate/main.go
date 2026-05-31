package main

import (
	"log"

	"github.com/eric-kuo/database_limit_test/backend/internal/config"
	"github.com/eric-kuo/database_limit_test/backend/internal/db"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("config: %v", err)
	}

	if err := db.Migrate(cfg.DatabaseURL); err != nil {
		log.Fatalf("migrate: %v", err)
	}

	log.Println("migrations applied")
}

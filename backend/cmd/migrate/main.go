package main

import (
	"log"

	"github.com/eric-kuo/database_limit_test/backend/internal/config"
	"github.com/eric-kuo/database_limit_test/backend/internal/db"
)

func main() {
	cfg := config.Load()

	if err := db.Migrate(cfg.DatabaseURL); err != nil {
		log.Fatalf("migrate: %v", err)
	}

	log.Println("migrations applied")
}

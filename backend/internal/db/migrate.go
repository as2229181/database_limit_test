package db

import (
	"database/sql"
	"fmt"

	"github.com/eric-kuo/database_limit_test/backend/migrations"
	"github.com/pressly/goose/v3"
	_ "github.com/jackc/pgx/v5/stdlib"
)

func Migrate(databaseURL string) error {
	sqlDB, err := sql.Open("pgx", databaseURL)
	if err != nil {
		return fmt.Errorf("open database: %w", err)
	}
	defer sqlDB.Close()

	goose.SetBaseFS(migrations.FS)

	if err := goose.SetDialect("postgres"); err != nil {
		return fmt.Errorf("set dialect: %w", err)
	}

	if err := goose.Up(sqlDB, "."); err != nil {
		return fmt.Errorf("run migrations: %w", err)
	}

	return nil
}

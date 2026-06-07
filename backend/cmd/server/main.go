package main

import (
	"log"
	"github.com/eric-kuo/database_limit_test/backend/internal/app"
)


func main() {
	APP, err := app.New()
	if err != nil {
		log.Fatalf("create app: %v", err)
	}

	APP.Start()
}
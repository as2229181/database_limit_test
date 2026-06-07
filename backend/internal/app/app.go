package app

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"fmt"
	"github.com/eric-kuo/database_limit_test/backend/internal/config"
	"github.com/eric-kuo/database_limit_test/backend/internal/db"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)


type App struct {
	cfg config.Config
	pool *pgxpool.Pool
    server *http.Server
}

func New() (*App, error){
	cfg, err := config.Load()
	if err != nil {
		return nil, fmt.Errorf("config: %w", err)
	}

	ctx := context.Background()

	pool, err := db.NewPool(ctx, cfg.DatabaseURL)
	if err != nil {
		log.Fatalf("connect database: %v", err)
	}

	router := gin.Default()

	app := &App {
		cfg: cfg,
		pool: pool,
	}

	app.registerRoutes(router)

	app.server = &http.Server{
		Addr:    cfg.HTTPAddr,
		Handler: router,
	}

	return app, nil
}

func (app *App) Start() {
	go func() {
		log.Printf("listening on %s", app.cfg.HTTPAddr)
        if err := app.server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("listen: %v", err)
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
	<-stop

	log.Println("server stopped")

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer cancel()

	if err := app.server.Shutdown(shutdownCtx); err != nil {
		log.Fatalf("shutdown: %v", err)
	}
	app.pool.Close()
	log.Println("server exited properly")
}
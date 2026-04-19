package main

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/codingbart/todoapp/task-api/internal/config"
	"github.com/codingbart/todoapp/task-api/internal/health"
	"github.com/codingbart/todoapp/task-api/internal/logger"
)

type Application interface {
	Mount() http.Handler
	Run(h http.Handler) error
}

type app struct {
	config config.Config
	logger logger.Logger
	db     *sql.DB
}

func NewApplication(config config.Config, logger logger.Logger, db *sql.DB) Application {
	return &app{
		config: config,
		logger: logger,
		db:     db,
	}
}

func (app *app) Mount() http.Handler {
	mux := http.NewServeMux()

	healthService := health.NewService()
	healthHandler := health.NewHandler(healthService)
	mux.HandleFunc("GET /api/health", healthHandler.GetHealthStatus)

	return mux
}

func (app *app) Run(h http.Handler) error {
	address := fmt.Sprintf("%s:%d", app.config.Host, app.config.Port)

	server := &http.Server{
		Addr:    address,
		Handler: h,
	}

	app.logger.Info("server started", "addr", address)

	return server.ListenAndServe()
}

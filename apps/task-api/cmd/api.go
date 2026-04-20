package main

import (
	"fmt"
	"net/http"

	"github.com/codingbart/todoapp/task-api/internal/config"
	db "github.com/codingbart/todoapp/task-api/internal/db/postgresql"
	"github.com/codingbart/todoapp/task-api/internal/health"
	"github.com/codingbart/todoapp/task-api/internal/logger"
)

type Application interface {
	Mount() http.Handler
	Run(h http.Handler) error
}

type app struct {
	config  config.Config
	logger  logger.Logger
	queries *db.Queries
}

func NewApplication(config config.Config, logger logger.Logger, queries *db.Queries) Application {
	return &app{
		config:  config,
		logger:  logger,
		queries: queries,
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

// @title           Task API
// @version         1.0
// @description     REST API dla aplikacji todo

package main

import (
	"fmt"
	"net/http"

	_ "github.com/codingbart/todoapp/task-api/docs"
	"github.com/codingbart/todoapp/task-api/internal/config"
	db "github.com/codingbart/todoapp/task-api/internal/db/postgresql"
	"github.com/codingbart/todoapp/task-api/internal/health"
	"github.com/codingbart/todoapp/task-api/internal/logger"
	httpSwagger "github.com/swaggo/http-swagger"
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
	mux.HandleFunc("GET /swagger/", httpSwagger.WrapHandler)

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

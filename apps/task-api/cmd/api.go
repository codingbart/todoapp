package main

import (
	"log"
	"net/http"

	"github.com/codingbart/todoapp/task-api/internal/health"
)

type Application interface {
	Mount() http.Handler
	Run(h http.Handler) error
}

type app struct {
	address string
}

func NewApplication(address string) Application {
	return &app{
		address: address,
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
	server := &http.Server{
		Addr:    app.address,
		Handler: h,
	}

	log.Printf("server has started at addr %s", app.address)

	return server.ListenAndServe()
}

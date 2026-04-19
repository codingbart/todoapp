package main

import (
	"os"

	"github.com/codingbart/todoapp/task-api/internal/config"
	"github.com/codingbart/todoapp/task-api/internal/db"
	"github.com/codingbart/todoapp/task-api/internal/logger"
)

func main() {
	log := logger.NewSlog()
	cfg := config.NewConfig(log)

	db, err := db.NewPostgresqlStorage(cfg)
	if err != nil {
		log.Error("failed to connect to database", "err", err)
		os.Exit(1)
	}
	log.Info("connected to database")

	app := NewApplication(cfg, log, db)

	if err := app.Run(app.Mount()); err != nil {
		log.Error("server error", "err", err)
		os.Exit(1)
	}
}

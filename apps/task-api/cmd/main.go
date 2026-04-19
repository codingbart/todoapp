package main

import (
	"os"

	"github.com/codingbart/todoapp/task-api/internal/config"
	"github.com/codingbart/todoapp/task-api/internal/logger"
)

func main() {
	log := logger.NewSlog()
	cfg := config.NewConfig(log)
	app := NewApplication(cfg, log)

	if err := app.Run(app.Mount()); err != nil {
		log.Error("server error", "err", err)
		os.Exit(1)
	}
}

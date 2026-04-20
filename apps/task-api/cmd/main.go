// @title           Task API
// @version         1.0
// @description     REST API dla aplikacji todo
//
// @securityDefinitions.oauth2.accessCode Keycloak
// @authorizationUrl KEYCLOAK_AUTH_URL
// @tokenUrl KEYCLOAK_TOKEN_URL

package main

import (
	"os"

	"github.com/codingbart/todoapp/task-api/internal/config"
	db "github.com/codingbart/todoapp/task-api/internal/db/postgresql"
	"github.com/codingbart/todoapp/task-api/internal/logger"
)

func main() {
	log := logger.NewSlog()
	cfg := config.NewConfig(log)

	queries, err := db.NewPostgresqlQueries(cfg)
	if err != nil {
		log.Error("failed to connect to database", "err", err)
		os.Exit(1)
	}
	log.Info("connected to database")

	app := NewApplication(cfg, log, queries)

	if err := app.Run(app.Mount()); err != nil {
		log.Error("server error", "err", err)
		os.Exit(1)
	}
}

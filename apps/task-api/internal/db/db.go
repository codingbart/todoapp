package db

import (
	"database/sql"

	"github.com/codingbart/todoapp/task-api/internal/config"
	_ "github.com/jackc/pgx/v5/stdlib"
)

func NewPostgresqlStorage(cfg config.Config) (*sql.DB, error) {
	db, err := sql.Open("pgx", cfg.DBUrl)
	if err != nil {
		return nil, err
	}

	return db, nil
}

package postgresql

import (
	"database/sql"

	"github.com/codingbart/todoapp/task-api/internal/config"
	_ "github.com/jackc/pgx/v5/stdlib"
)

func NewPostgresqlQueries(cfg config.Config) (*Queries, error) {
	db, err := sql.Open("pgx", cfg.DBUrl)
	if err != nil {
		return nil, err
	}

	return New(db), nil
}

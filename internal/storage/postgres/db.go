package postgres

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/s4bb4t/url-shortener/internal/config"
)

type Storage struct {
	db *sql.DB
}

func New(cfg config.Postgres) (*Storage, error) {
	db, err := sql.Open("postgres", cfg.DSN())
	if err != nil {
		return nil, fmt.Errorf("could not connect to postgres: %w", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("could not ping postgres: %w", err)
	}

	return &Storage{db: db}, nil
}

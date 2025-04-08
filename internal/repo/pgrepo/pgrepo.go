package pgrepo

import (
	"github.com/siavoid/shortener/pkg/logger"
	"github.com/siavoid/shortener/pkg/postgres"
)

type PostgresRepo struct {
	db     *postgres.Postgres
	logger *logger.Logger
}

// New -.
func New(pg *postgres.Postgres, l *logger.Logger) *PostgresRepo {
	return &PostgresRepo{
		db:     pg,
		logger: l,
	}
}

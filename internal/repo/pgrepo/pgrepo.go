package pgrepo

import (
	"github.com/siavoid/shortener/pkg/logger"
	"github.com/siavoid/shortener/pkg/postgres"
)

// Задел на будущее
type Interface interface {
}

type PostgresRepo struct {
	db     *postgres.Postgres
	logger *logger.Logger
}

var _ Interface = (*PostgresRepo)(nil)

// New -.
func New(pg *postgres.Postgres, l *logger.Logger) *PostgresRepo {
	return &PostgresRepo{
		db:     pg,
		logger: l,
	}
}

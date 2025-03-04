// Package usecase implements application business logic. Each logic group in own file.
package usecase

import (
	"context"

	"github.com/siavoid/shortener/config"
	"github.com/siavoid/shortener/internal/repo/pgrepo"
	"github.com/siavoid/shortener/internal/repo/urlstore"
	"github.com/siavoid/shortener/pkg/logger"
)

type (
	Interface interface {
		shortener
	}

	// Shortener - .
	shortener interface {
		GetShortenURL(context.Context, string) (string, error)
		GetOriginalURL(context.Context, string) (string, error)
	}
)

type UseCase struct {
	db       pgrepo.Interface
	l        logger.Interface
	baseURL  string
	urlStore *urlstore.URLStore
}

var _ Interface = (*UseCase)(nil)

func New(cfg *config.Config, l logger.Interface, db pgrepo.Interface) *UseCase {
	return &UseCase{
		db:       db,
		l:        l,
		baseURL:  cfg.Shortener.BaseURL,
		urlStore: urlstore.NewURLStore(),
	}
}

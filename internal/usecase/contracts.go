// Package usecase implements application business logic. Each logic group in own file.
package usecase

import (
	"context"

	"github.com/siavoid/shortener/config"
	"github.com/siavoid/shortener/internal/repo/pgrepo"
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

	URLStoreInterface interface {
		GetLongURL(shortURL string) (string, bool)
		GetShortURL(url string) (string, bool)
		Put(url string, shortURL string) error
	}
)

type UseCase struct {
	db       pgrepo.Interface
	l        logger.Interface
	baseURL  string
	urlStore URLStoreInterface
}

var _ Interface = (*UseCase)(nil)

func New(cfg *config.Config, l logger.Interface, db pgrepo.Interface, repo URLStoreInterface) *UseCase {
	return &UseCase{
		db:       db,
		l:        l,
		baseURL:  cfg.Shortener.BaseURL,
		urlStore: repo,
	}
}

// Package usecase implements application business logic. Each logic group in own file.
package usecase

import (
	"context"

	"github.com/siavoid/shortener/config"
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
		RepoPing(context.Context) error
	}

	URLStoreInterface interface {
		GetLongURL(shortURL string) (string, bool)
		GetShortURL(url string) (string, bool)
		Put(url string, shortURL string) error
		Ping(ctx context.Context) error
	}
)

type UseCase struct {
	l        logger.Interface
	baseURL  string
	urlStore URLStoreInterface
}

var _ Interface = (*UseCase)(nil)

func New(cfg *config.Config, l logger.Interface, repo URLStoreInterface) *UseCase {
	return &UseCase{
		l:        l,
		baseURL:  cfg.Shortener.BaseURL,
		urlStore: repo,
	}
}

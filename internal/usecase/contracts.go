// Package usecase implements application business logic. Each logic group in own file.
package usecase

import (
	"context"
	"fmt"
	"sync"

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
	url      string
	urlStore *urlstore.URLStore

	// Поскольку получение по короткой ссылке оригинальной url намного быстрее,
	// чем наоброт, может возникнуть ситуация, когда короткая ссылка ещё не сохранилась,
	// но её уже запрашивают.
	mu sync.Mutex
}

var _ Interface = (*UseCase)(nil)

func New(cfg *config.Config, l logger.Interface, db pgrepo.Interface) *UseCase {
	url := fmt.Sprintf("%s:%s", cfg.HTTP.Host, cfg.HTTP.Port)
	return &UseCase{
		db:       db,
		l:        l,
		url:      "http://" + url,
		urlStore: urlstore.NewURLStore(),
	}
}

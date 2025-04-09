package v1

import (
	"context"
	"net/http"

	"github.com/siavoid/shortener/config"
	"github.com/siavoid/shortener/internal/entity"
	"github.com/siavoid/shortener/pkg/logger"

	"github.com/gorilla/mux"

	_ "github.com/siavoid/shortener/docs"
)

// @title URL Shortener API
// @version         2.0
// @description API for shortening URLs.

// @host localhost:8080
// @BasePath /

// Shortener - .
type UserCaseShortener interface {
	GetShortenURL(context.Context, string) (string, error)
	GetOriginalURL(context.Context, string) (string, error)
	RepoPing(context.Context) error
	ShortenBatch(ctx context.Context, origBatch []entity.BatchOriginalURL) ([]entity.BatchShortURL, error)
}

type LoggerInterface logger.Interface

// Server - http server
type Server struct {
	url        string
	router     *mux.Router
	httpServer *http.Server
	u          UserCaseShortener
	logger     LoggerInterface
}

func New(cfg *config.Config, u UserCaseShortener, l LoggerInterface) *Server {
	router := mux.NewRouter()
	url := cfg.HTTP.ServerAddress
	s := Server{
		u:      u,
		router: router,
		url:    url,
		httpServer: &http.Server{
			Addr:    url,
			Handler: router,
		},
		logger: l,
	}
	s.routeRegistration() // регистрация маршрутов
	return &s
}

// Run - .
func (s *Server) Run() error {
	s.logger.Info("Сервер запущен : %s", s.url)
	return s.httpServer.ListenAndServe()
}

// Stop - .
func (s *Server) Stop(ctx context.Context) error {
	// Останавливаем сервер
	if err := s.httpServer.Shutdown(ctx); err != nil {
		return err
	}
	return nil
}

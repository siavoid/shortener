package v1

import (
	"context"
	"net/http"
	"time"

	"github.com/siavoid/shortener/config"
	"github.com/siavoid/shortener/internal/usecase"
	"github.com/siavoid/shortener/pkg/logger"

	"github.com/gorilla/mux"

	_ "github.com/siavoid/shortener/docs"
)

// @title URL Shortener API
// @version         2.0
// @description API for shortening URLs.

// @host localhost:8080
// @BasePath /

type UserCaseInterface usecase.Interface

type LoggerInterface logger.Interface

// Server - http server
type Server struct {
	url        string
	router     *mux.Router
	httpServer *http.Server
	u          UserCaseInterface
	logger     LoggerInterface
}

func New(cfg *config.Config, u UserCaseInterface, l LoggerInterface) *Server {
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
	// Устанавливаем таймаут для завершения работы сервера
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	// Останавливаем сервер
	if err := s.httpServer.Shutdown(ctx); err != nil {
		return err
	}
	return nil
}

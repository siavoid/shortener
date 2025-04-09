package v1

import (
	"net/http"

	"github.com/siavoid/shortener/internal/controllers/http/v1/middleware"
	httpSwagger "github.com/swaggo/http-swagger"
)

func (s *Server) routeRegistration() {
	// swagger
	s.router.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)
	//s.router.HandleFunc("/ping", s.repoPing).Methods(http.MethodGet)

	shortenRouter := s.router.PathPrefix("").Subrouter()

	// shortenRouter.Use(middleware.CORSMiddleware) // включение CORS заголовков
	shortenRouter.Use(middleware.GzipMiddleware)
	shortenRouter.Use(middleware.LoggingMiddleware(s.logger))
	shortenRouter.HandleFunc("/ping", s.repoPing).Methods(http.MethodGet)

	shortenRouter.HandleFunc("/", s.shortenURLHandler).Methods(http.MethodPost, http.MethodOptions)
	shortenRouter.HandleFunc("/{id}", s.getOriginalURLHandler).Methods(http.MethodGet)

	shortenRouter.HandleFunc("/api/shorten", s.shortenURLInJSONHandler).Methods(http.MethodPost, http.MethodOptions)

	shortenRouter.HandleFunc("/api/shorten/batch", s.shortenBatchHandler).Methods(http.MethodPost, http.MethodOptions)
}

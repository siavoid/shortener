package v1

import (
	"net/http"

	httpSwagger "github.com/swaggo/http-swagger"
)

func (s *Server) routeRegistration() {
	// swagger
	s.router.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	shortenRouter := s.router.PathPrefix("").Subrouter()

	// shortenRouter.Use(middleware.CORSMiddleware) // включение CORS заголовков

	shortenRouter.HandleFunc("/", s.shortenURLHandler).Methods(http.MethodPost, http.MethodOptions)
	shortenRouter.HandleFunc("/{id}", s.getOriginalURLHandler).Methods(http.MethodGet)
}

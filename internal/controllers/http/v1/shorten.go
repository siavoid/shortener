package v1

import (
	"fmt"
	"io"
	"net/http"

	"github.com/gorilla/mux"
)

// @Summary Shorten URL
// @Description Receives a URL and returns a shortened version.
// Tags shortener
// @Accept text/plain
// @Produce text/plain
// @Param url body string true "URL to be shortened"
// @Success 201 {string} string "Shortened URL"
// @Failure 400 {string} string "Invalid request"
// @Failure 500 {string} string "Error shortening URL"
// @Router / [post]
func (s *Server) shortenURLHandler(w http.ResponseWriter, r *http.Request) {
	// Чтение URL из тела запроса
	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	originalURL := string(body)

	// Сокращение URL
	shortenedURL, err := s.u.GetShortenURL(r.Context(), originalURL)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Формирование ответа
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusCreated)

	// Как лучше и почему?
	fmt.Fprint(w, shortenedURL) // w.Write([]byte(shortenedURL))
}

// @Summary Get Original URL
// @Description Redirects to the original URL based on the shortened URL ID.
// @Produce plain
// @Param id path string true "Shortened URL ID"
// @Success 307 {string} string "Temporary Redirect"
// @Failure 400 {string} string "Invalid request"
// @Router /{id} [get]
func (s *Server) getOriginalURLHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	// Получение оригинального URL по идентификатору
	originalURL, err := s.u.GetOriginalURL(r.Context(), id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.Header().Set("Location", originalURL)
	w.WriteHeader(http.StatusTemporaryRedirect)
}

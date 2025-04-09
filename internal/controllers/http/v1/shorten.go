package v1

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/siavoid/shortener/internal/controllers/http/v1/dto"
	"github.com/siavoid/shortener/internal/internalerror"
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
	if err != nil && !errors.Is(err, internalerror.ErrConflict) {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Формирование ответа
	w.Header().Set("Content-Type", "text/plain")
	if errors.Is(err, internalerror.ErrConflict) {
		w.WriteHeader(http.StatusConflict)
	} else {
		w.WriteHeader(http.StatusCreated)
	}
	fmt.Fprint(w, shortenedURL)
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

// @Summary Shorten URL
// @Description Receives a URL in JSON and returns a shortened version in JSON.
// Tags shortener
// @Accept application/json
// @Produce application/json
// @Param url body dto.ShortenURLRequest true "URL to be shortened"
// @Success 200 {object} dto.ShortenURLResponse "Shortened URL"
// @Failure 400 {string} string "Invalid request"
// @Failure 500 {string} string "Error shortening URL"
// @Router /api/shorten [post]
func (s *Server) shortenURLInJSONHandler(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil || len(body) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "read body")
		return
	}

	var req dto.ShortenURLRequest
	err = json.Unmarshal(body, &req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "unmarshal body", err.Error())
		return
	}

	if !req.Validate() {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "url is empty")
		return
	}

	// Сокращение URL
	shortenedURL, err := s.u.GetShortenURL(r.Context(), req.URL)
	if err != nil && !errors.Is(err, internalerror.ErrConflict) {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	res := dto.ShortenURLResponse{
		URL: shortenedURL,
	}

	// Формирование ответа
	w.Header().Set("Content-Type", "application/json")
	if errors.Is(err, internalerror.ErrConflict) {
		w.WriteHeader(http.StatusConflict)
	} else {
		w.WriteHeader(http.StatusCreated)
	}
	json.NewEncoder(w).Encode(res)
}

package v1

import (
	"encoding/json"
	"net/http"

	"github.com/siavoid/shortener/internal/entity"
)

func (s *Server) shortenBatchHandler(w http.ResponseWriter, r *http.Request) {
	var requests []entity.BatchOriginalURL

	// Декодируем тело запроса
	if err := json.NewDecoder(r.Body).Decode(&requests); err != nil {
		http.Error(w, "Invalid request format", http.StatusBadRequest)
		return
	}

	res, err := s.u.ShortenBatch(r.Context(), requests)
	if err != nil {
		s.logger.Error("shorten batch: %s", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}

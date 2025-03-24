package middleware

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/siavoid/shortener/pkg/logger"
)

// TODO: Сведения о запросах должны содержать URI, метод запроса и время, затраченное на его выполнение.
// TODO: Сведения об ответах должны содержать код статуса и размер содержимого ответа.

type LogData struct {
	URI          string        `json:"uri"`
	Method       string        `json:"method"`
	Status       int           `json:"status"`
	ResponseSize int           `json:"response_size"`
	Duration     time.Duration `json:"duration"`
}

// LoggingMiddleware logs details about incoming requests and outgoing responses.
func LoggingMiddleware(log logger.Interface) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()
			rw := &responseWriter{ResponseWriter: w, statusCode: http.StatusOK}
			next.ServeHTTP(rw, r)
			duration := time.Since(start)
			logData := LogData{
				URI:          r.RequestURI,
				Method:       r.Method,
				Status:       rw.statusCode,
				ResponseSize: rw.size,
				Duration:     duration,
			}
			logDataByte, err := json.Marshal(logData)
			if err != nil {
				log.Warn("request stat error : %s", err.Error())
			}
			log.Info("request stat: %s", string(logDataByte))

		})
	}
}

// responseWriter is a custom writer to capture the status code and response size
type responseWriter struct {
	http.ResponseWriter
	statusCode int
	size       int
}

// WriteHeader captures the status code
func (rw *responseWriter) WriteHeader(code int) {
	rw.statusCode = code
	rw.ResponseWriter.WriteHeader(code)
}

// Write captures the size of the response
func (rw *responseWriter) Write(b []byte) (int, error) {
	size, err := rw.ResponseWriter.Write(b)
	rw.size += size
	return size, err
}

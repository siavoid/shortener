package middleware

import (
	"compress/gzip"
	"io"
	"net/http"
	"strings"
)

// Middleware для распаковки gzip-запросов
func GzipRequestMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Content-Encoding") == "gzip" {
			gzipReader, err := gzip.NewReader(r.Body)
			if err != nil {
				http.Error(w, "Failed to create gzip reader", http.StatusBadRequest)
				return
			}
			defer gzipReader.Close()
			r.Body = gzipReader
			w.Header().Set("Request-Encoding", "gzip")
		}
		next.ServeHTTP(w, r)
	})
}

// Middleware для упаковки gzip-ответов
func GzipResponseMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.Contains(r.Header.Get("Accept-Encoding"), "gzip") {
			gzipWriter := gzip.NewWriter(w)
			defer gzipWriter.Close()

			w.Header().Set("Content-Encoding", "gzip")
			next.ServeHTTP(&gzipResponseWriter{ResponseWriter: w, Writer: gzipWriter}, r)
		} else {
			next.ServeHTTP(w, r)
		}
	})
}

type gzipResponseWriter struct {
	http.ResponseWriter
	Writer io.Writer
}

func (w *gzipResponseWriter) Write(b []byte) (int, error) {
	var gzipAllowed bool
	contentType := w.Header().Get("Content-Type")
	if contentType == "text/html" || contentType == "application/json" {
		gzipAllowed = true
	}
	if w.Header().Get("Request-Encoding") == "gzip" && gzipAllowed {
		w.Writer.Write(b)
	}
	w.Header().Del("Request-Encoding")
	return w.ResponseWriter.Write(b)
}

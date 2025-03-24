package middleware

import (
	"compress/gzip"
	"io"
	"net/http"
	"strings"
)

// GzipMiddleware обрабатывает как распаковку запросов, так и упаковку ответов
func GzipMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Распаковка gzip-запросов
		if r.Header.Get("Content-Encoding") == "gzip" {
			gzipReader, err := gzip.NewReader(r.Body)
			if err != nil {
				http.Error(w, "Failed to create gzip reader", http.StatusBadRequest)
				return
			}
			defer gzipReader.Close()
			r.Body = gzipReader
		}

		// Упаковка gzip-ответов
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
	// var gzipAllowed bool
	// contentType := w.Header().Get("Content-Type")
	// if contentType == "text/html" || contentType == "application/json" {
	// 	gzipAllowed = true
	// }
	// if gzipAllowed {
	// 	w.Writer.Write(b)
	// }
	// return w.ResponseWriter.Write(b)
	return w.Writer.Write(b)
}

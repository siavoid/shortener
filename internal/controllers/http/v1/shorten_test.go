package v1

import (
	"bytes"
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gorilla/mux"
	"github.com/siavoid/shortener/config"
	"github.com/siavoid/shortener/internal/usecase"
	"github.com/stretchr/testify/assert"
)

func Test_ShortenURLHandler(t *testing.T) {
	cfg := config.Config{
		HTTP: config.HTTP{
			Host: "localhost",
			Port: "8080",
		},
	}
	useCase := usecase.New(&cfg, nil, nil)

	server := &Server{u: useCase}

	tests := []struct {
		name           string
		input          string
		expectedStatus int
	}{
		{"valid URL", "https://example.com", http.StatusCreated},
		{"empty URL", "", http.StatusBadRequest},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBufferString(tt.input))
			rec := httptest.NewRecorder()

			server.shortenURLHandler(rec, req)

			res := rec.Result()
			defer res.Body.Close()

			assert.Equal(t, tt.expectedStatus, res.StatusCode)
		})
	}
}

func Test_GetOriginalURLHandler(t *testing.T) {
	cfg := config.Config{
		HTTP: config.HTTP{
			Host: "localhost",
			Port: "8080",
		},
	}
	localURL := fmt.Sprintf("http://%s:%s/", cfg.HTTP.Host, cfg.HTTP.Port)
	useCase := usecase.New(&cfg, nil, nil)
	server := &Server{u: useCase}

	type want struct {
		statusCode int
		url        string
	}

	type testcase struct {
		name    string
		request string
		want    want
	}
	tests := []testcase{
		{
			name:    "get original URL without creating short URL",
			request: "/test",
			want: want{
				statusCode: http.StatusBadRequest,
				url:        "",
			},
		},
	}

	originalURL := "https://example.com"
	shortURL, _ := useCase.GetShortenURL(context.Background(), originalURL)
	shortURL = strings.TrimPrefix(shortURL, localURL)
	tests = append(tests, testcase{
		name:    "create and get original URL",
		request: "/" + shortURL,
		want: want{
			statusCode: http.StatusTemporaryRedirect,
			url:        originalURL,
		},
	})

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			req := httptest.NewRequest(http.MethodGet, tt.request, nil)
			rec := httptest.NewRecorder()

			// Use mux router to simulate path variable
			router := mux.NewRouter()
			router.HandleFunc("/{id}", server.getOriginalURLHandler)
			router.ServeHTTP(rec, req)

			res := rec.Result()
			defer res.Body.Close()

			assert.Equal(t, tt.want.statusCode, res.StatusCode)
			if res.StatusCode == http.StatusTemporaryRedirect {
				url := res.Header.Get("Location")
				assert.Equal(t, tt.want.url, url)
			}
		})
	}
}

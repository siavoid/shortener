package usecase

import (
	"context"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/siavoid/shortener/internal/repo/urlstore"
)

func Test_shortenURL(t *testing.T) {
	u := &UseCase{}
	tests := []struct {
		name string
		url  string
	}{
		{"simple URL", "https://example.com"},
		{"URL with query", "https://example.com?query=123"},
		{"URL with path", "https://example.com/path"},
		{"URL with special chars", "https://example.com/@#$%"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			shortURL := u.shortenURL(tt.url)
			assert.GreaterOrEqual(t, len(shortURL), 4, "shortened URL should have at least 4 characters")
		})
	}
}

func Test_cleanURL(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"no special chars", "abcdef", "abcdef"},
		{"with special chars", "abc$%def", "abcaadef"},
		{"long input", "abcdefghijklm", "abcdefgh"},
		{"only special chars", "$%#@!*", "aaaaaa"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cleaned := cleanURL(tt.input)
			require.GreaterOrEqual(t, 8, len(cleaned), "cleaned URL should not exceed 8 characters")
			assert.Equal(t, tt.expected, cleaned)
		})
	}
}

func TestUseCase_GetShortenURL(t *testing.T) {
	storeFile := "test.json"
	defer os.Remove(storeFile)
	urlStore, err := urlstore.NewURLStore(storeFile)
	require.NoError(t, err)
	u := &UseCase{
		baseURL:  "http://localhost",
		urlStore: urlStore,
	}

	tests := []struct {
		name        string
		originalURL string
	}{
		{"simple URL", "https://example.com/test1"},
		{"URL with query", "https://example.com/test2?query=123"},
		{"URL with path", "https://example.com/test3/path"},
		{"URL with special chars", "https://example.com/test4/@#$%"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			shortURL, err := u.GetShortenURL(context.Background(), tt.originalURL)
			require.NoError(t, err, "GetShortenURL should not return an error")

			shortURL = strings.TrimPrefix(shortURL, u.baseURL+"/")

			retrievedURL, err := u.GetOriginalURL(context.Background(), shortURL)
			require.NoError(t, err, "GetOriginalURL should not return an error")
			assert.Equal(t, tt.originalURL, retrievedURL, "retrieved URL should match the original URL")
		})
	}
}

func TestUseCase_GetOriginalURL(t *testing.T) {
	storeFile := "test.json"
	defer os.Remove(storeFile)
	urlStore, err := urlstore.NewURLStore(storeFile)
	require.NoError(t, err)
	u := &UseCase{
		urlStore: urlStore,
	}

	tests := []struct {
		name        string
		shortURL    string
		originalURL string
		wantErr     bool
	}{
		{"existing short URL", "short1", "https://example.com/test1", false},
		{"non-existent short URL", "nonexistent", "", true},
		{"empty short URL", "", "", true},
		{"special chars", "short@", "", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if !tt.wantErr {
				u.urlStore.Put(tt.originalURL, tt.shortURL)
			}

			retrievedURL, err := u.GetOriginalURL(context.Background(), tt.shortURL)
			if tt.wantErr {
				assert.Error(t, err, "expected an error for non-existent short URL")
			} else {
				require.NoError(t, err, "GetOriginalURL should not return an error")
				assert.Equal(t, tt.originalURL, retrievedURL, "retrieved URL should match the original URL")
			}
		})
	}
}

func TestUseCase_CreateShortURLAndGetOriginalURL(t *testing.T) {
	storeFile := "test.json"
	defer os.Remove(storeFile)
	urlStore, err := urlstore.NewURLStore(storeFile)
	require.NoError(t, err)
	u := &UseCase{
		urlStore: urlStore,
	}

	tests := []struct {
		name        string
		originalURL string
		wantErr     bool
	}{
		{"new short URL", "https://example.com/test1", false},
		{"existing short URL", "https://example.com/test1", false},
		{"existing short URL", "https://example.com/test1?id=30&text=qwerqwerqwrq", false},
		{"empty short URL", "", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			shortURL, err := u.GetShortenURL(context.Background(), tt.originalURL)
			if tt.wantErr {
				assert.Error(t, err)
			}
			shortURL = strings.TrimPrefix(shortURL, u.baseURL+"/")
			retrievedURL, err := u.GetOriginalURL(context.Background(), shortURL)
			if tt.wantErr {
				assert.Error(t, err, "expected an error for non-existent short URL")
			} else {
				require.NoError(t, err, "GetOriginalURL should not return an error")
				assert.Equal(t, tt.originalURL, retrievedURL, "retrieved URL should match the original URL")
			}
		})
	}
}

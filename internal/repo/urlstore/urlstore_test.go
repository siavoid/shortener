package urlstore

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_URLStore(t *testing.T) {
	type urlPair struct {
		url      string
		shortURL string
	}

	tests := []struct {
		name      string
		storeInit []urlPair
		checks    []urlPair
	}{
		{
			name: "add and retrieve URLs",
			storeInit: []urlPair{
				{"https://example.com/1", "ex1"},
				{"https://example.com/2", "ex2"},
				{"https://example.com/3", "ex3"},
			},
			checks: []urlPair{
				{"https://example.com/1", "ex1"},
				{"https://example.com/2", "ex2"},
				{"https://example.com/3", "ex3"},
			},
		},
		{
			name:      "retrieve non-existing URLs",
			storeInit: []urlPair{},
			checks: []urlPair{
				{"https://example.com/nonexistent", ""},
				{"", "nonexistent"},
			},
		},
	}

	testStorePath := "test.json"
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer os.Remove(testStorePath)
			store, err := NewURLStore(testStorePath) // TODO: затем удалить этот файл
			require.NoError(t, err)

			// Initialize store with URL pairs
			for _, pair := range tt.storeInit {
				store.Put(pair.url, pair.shortURL)
			}

			// Check stored URLs
			for _, pair := range tt.checks {
				if pair.url != "" {
					shortURL, exists := store.GetShortURL(pair.url)
					require.Equal(t, pair.shortURL != "", exists)
					assert.Equal(t, pair.shortURL, shortURL)
				}

				if pair.shortURL != "" {
					longURL, exists := store.GetLongURL(pair.shortURL)
					require.Equal(t, pair.url != "", exists)
					assert.Equal(t, pair.url, longURL)
				}
			}
		})
	}
}

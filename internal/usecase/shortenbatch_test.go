package usecase

import (
	"context"
	"os"
	"testing"

	"github.com/siavoid/shortener/internal/entity"
	"github.com/siavoid/shortener/internal/repo/urlstore"
	"github.com/siavoid/shortener/pkg/logger"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestShortenBatch(t *testing.T) {
	ctx := context.Background()
	file := "test.json"
	defer os.Remove(file)
	repo, err := urlstore.NewURLStore(file)
	require.NoError(t, err)
	u := &UseCase{
		logger.New("debug"),
		"localhost:8080",
		repo,
	}

	// Входные данные
	batchOrig := []entity.BatchOriginalURL{
		{CorrelationID: "1", OriginalURL: "https://example.com"},
		{CorrelationID: "2", OriginalURL: "https://example.org"},
		{CorrelationID: "3", OriginalURL: "https://example.net"},
	}

	// Ожидаемые выходные данные
	// Здесь предполагается, что логика `GetShortenURL` возвращает "short" + CorrelationID
	expectedBatchShort := []entity.BatchShortURL{
		{CorrelationID: "1", ShortURL: "short1"},
		{CorrelationID: "2", ShortURL: "short2"},
		{CorrelationID: "3", ShortURL: "short3"},
	}

	for i := range len(expectedBatchShort) {
		expectedBatchShort[i].ShortURL, err = u.GetShortenURL(context.Background(), batchOrig[i].OriginalURL)
		require.NoError(t, err)
	}
	// Вызов тестируемой функции
	batchShort, err := u.ShortenBatch(ctx, batchOrig)

	// Проверка результатов
	require.NoError(t, err)
	assert.Equal(t, expectedBatchShort, batchShort)

	for i := range len(expectedBatchShort) {
		expectedBatchShort[i].ShortURL, err = u.GetShortenURL(context.Background(), batchOrig[i].OriginalURL)
		require.NoError(t, err)
	}
	assert.Equal(t, expectedBatchShort, batchShort)
}

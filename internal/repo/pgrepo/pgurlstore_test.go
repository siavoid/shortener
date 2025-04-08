package pgrepo

import (
	"os"
	"testing"

	"github.com/siavoid/shortener/internal/entity"
	"github.com/siavoid/shortener/pkg/logger"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func setupTestDB(dsn string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	err = db.AutoMigrate(&entity.URL{})
	return db, err
}

func TestPostgresRepo_PutAndGet(t *testing.T) {
	dsn := os.Getenv("DATABASE_DSN")
	if dsn == "" {
		return
	}
	db, err := setupTestDB(dsn)
	require.NoError(t, err)
	defer func() {
		sqlDB, _ := db.DB()
		sqlDB.Close()
	}()

	logger := logger.New("debug") // Use your actual logger initialization
	repo, err := NewPostgresRepo(dsn, logger)
	require.NoError(t, err)

	originalURL := "https://example.com"
	shortURL := "exmpl"

	// Test Put
	err = repo.Put(originalURL, shortURL)
	defer repo.DeleteByOriginalURL(originalURL)
	require.NoError(t, err)

	// Test GetLongURL
	longURL, found := repo.GetLongURL(shortURL)
	assert.True(t, found)
	assert.Equal(t, originalURL, longURL)

	// Test GetShortURL
	retrievedShortURL, found := repo.GetShortURL(originalURL)
	assert.True(t, found)
	assert.Equal(t, shortURL, retrievedShortURL)
}

func TestPostgresRepo_GetNonExistent(t *testing.T) {
	dsn := os.Getenv("DATABASE_DSN")
	if dsn == "" {
		return
	}
	db, err := setupTestDB(dsn)
	require.NoError(t, err)
	defer func() {
		sqlDB, _ := db.DB()
		sqlDB.Close()
	}()

	logger := logger.New("debug") // Use your actual logger initialization
	repo, err := NewPostgresRepo(dsn, logger)
	require.NoError(t, err)

	// Test GetLongURL for non-existent URL
	_, found := repo.GetLongURL("nonexistent")
	assert.False(t, found)

	// Test GetShortURL for non-existent URL
	_, found = repo.GetShortURL("https://nonexistent.com")
	assert.False(t, found)
}

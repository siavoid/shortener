package pgrepo

import (
	"github.com/siavoid/shortener/internal/entity"
	"github.com/siavoid/shortener/pkg/logger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PostgresRepo struct {
	DB     *gorm.DB
	logger logger.Interface
}

func NewPostgresRepo(dsn string, logger *logger.Logger) (*PostgresRepo, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		logger.Error("failed to connect to database")
		return nil, err
	}

	pg := &PostgresRepo{
		DB:     db,
		logger: logger,
	}
	if err = pg.DB.AutoMigrate(&entity.URL{}); err != nil {
		pg.logger.Fatal("migration failed : %s", err.Error())
		return pg, err
	}
	return pg, nil
}

func (p *PostgresRepo) AutoMigrate(models ...interface{}) error {
	err := p.DB.AutoMigrate(models...)
	if err != nil {
		p.logger.Error("failed to apply migrations")
		return err
	}

	p.logger.Info("database migrations applied successfully")
	return nil
}

func (p *PostgresRepo) DeleteByShortURL(shortURL string) error {
	result := p.DB.Delete(&entity.URL{}, "short_url = ?", shortURL)
	if result.Error != nil {
		p.logger.Error("DeleteByShortURL error: %v", result.Error)
		return result.Error
	}
	return nil
}

func (p *PostgresRepo) DeleteByOriginalURL(originalURL string) error {
	result := p.DB.Delete(&entity.URL{}, "original_url = ?", originalURL)
	if result.Error != nil {
		p.logger.Error("DeleteByOriginalURL error: %v", result.Error)
		return result.Error
	}
	return nil
}

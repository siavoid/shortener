package pgrepo

import "github.com/siavoid/shortener/internal/entity"

func (p *PostgresRepo) GetLongURL(shortURL string) (string, bool) {
	var url entity.URL
	result := p.DB.First(&url, "short_url = ?", shortURL)
	if result.Error != nil {
		p.logger.Error("GetLongURL error: %v", result.Error)
		return "", false
	}
	return url.OriginalURL, true
}

func (p *PostgresRepo) GetShortURL(originalURL string) (string, bool) {
	var url entity.URL
	result := p.DB.First(&url, "original_url = ?", originalURL)
	if result.Error != nil {
		p.logger.Error("GetShortURL error: %v", result.Error)
		return "", false
	}
	return url.ShortURL, true
}

func (p *PostgresRepo) Put(originalURL string, shortURL string) error {
	url := entity.URL{ShortURL: shortURL, OriginalURL: originalURL}
	result := p.DB.Create(&url)
	if result.Error != nil {
		p.logger.Error("Put error: %v", result.Error)
		return result.Error
	}
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

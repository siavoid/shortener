package usecase

import (
	"context"
	"fmt"

	"github.com/siavoid/shortener/internal/entity"
)

func (u *UseCase) ShortenBatch(ctx context.Context, batchOrig []entity.BatchOriginalURL) ([]entity.BatchShortURL, error) {
	batchShort := make([]entity.BatchShortURL, 0, len(batchOrig))
	for _, v := range batchOrig {
		shortURL, err := u.GetShortenURL(ctx, v.OriginalURL)
		if err != nil {
			return nil, fmt.Errorf("get short url: %w", err)
		}
		batchShort = append(
			batchShort, entity.BatchShortURL{
				CorrelationID: v.CorrelationID,
				ShortURL:      shortURL,
			},
		)
	}

	return batchShort, nil
}

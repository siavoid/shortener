package usecase

import "context"

func (u *UseCase) GetShortenURL(ctx context.Context, url string) (string, error) {
	return "", nil
}

func (u *UseCase) GetOriginalURL(ctx context.Context, shortURL string) (string, error) {
	return "", nil
}

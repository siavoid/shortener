package usecase

import (
	"context"
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"fmt"
	"math/rand"
	"strings"
)

// ShortenURL создает сокращённый URL на основе хеша и случайного числа
func (u *UseCase) shortenURL(url string) string {
	// Генерация хеша от URL
	hash := sha256.Sum256([]byte(url))

	// Получение случайного числа
	randomNumber := rand.Intn(9999)

	// Комбинирование первых 4 байт хеша с случайным числом для уникальности
	shortURL := base64.URLEncoding.EncodeToString(hash[:4]) + base64.URLEncoding.EncodeToString([]byte{byte(randomNumber)})

	// Удаление символов, которые не подходят для URL
	shortURL = cleanURL(shortURL)

	return shortURL
}

// cleanURL удаляет нежелательные символы из URL
func cleanURL(url string) string {
	result := ""
	for _, c := range url {
		if (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') || (c >= '0' && c <= '9') {
			result += string(c)
		}
	}
	return result[:8] // Ограничение длины до 8 символов
}

func (u *UseCase) GetShortenURL(ctx context.Context, url string) (string, error) {
	url = strings.TrimSpace(url)
	shortURL, err := u.createOrGetShortenURL(ctx, url)
	shortURL = fmt.Sprintf("%s/%s", u.url, shortURL)
	return shortURL, err
}

func (u *UseCase) createOrGetShortenURL(ctx context.Context, url string) (string, error) {
	// есть ли уже сокращенная ссылка
	shortURL, ok := u.urlStore.GetShortUrl(url)
	if ok {
		return shortURL, nil
	}

	// создание короткой ссылки
	for i := 0; i < 100; i++ { // маловроятно, но вдруг ...
		shortURL = u.shortenURL(url)
		// проверим, что короткая ссылка не занята
		if _, ok := u.urlStore.GetLongUrl(shortURL); !ok {
			break
		}
	}
	shortURL = u.shortenURL(url)
	u.urlStore.Put(url, shortURL)

	return shortURL, nil
}

func (u *UseCase) GetOriginalURL(ctx context.Context, shortURL string) (string, error) {
	shortURL = strings.TrimSpace(shortURL)
	if url, ok := u.urlStore.GetLongUrl(shortURL); ok {
		return url, nil
	}

	return "", errors.New("no data")
}

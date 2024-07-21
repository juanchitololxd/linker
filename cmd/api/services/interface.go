package services

import "url-shortener/cmd/api/domain"

type URLShortenerService interface {
	ShortenURL(originalURL string) domain.URLMapping
	GetOriginalURL(shortURL string) (string, bool)
	GetHistory() []domain.URLMapping
	GetPing() string
}

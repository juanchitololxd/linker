package services

import "url-shortener/cmd/api/domain"

type (
	URLShortenerService interface {
		ShortenURL(originalURL string) (domain.URLMapping, error)
		GetOriginalURL(shortURL string) (string, error)
		GetHistory() ([]domain.URLMapping, error)
		GetPing() string
	}

	URLRepository interface {
		Save(url domain.URLMapping) error
		FindByOriginalURL(originalURL string) (domain.URLMapping, error)
		FindByShortURL(shortURL string) (domain.URLMapping, error)
		FindAll() ([]domain.URLMapping, error)
	}
)

package services

import (
	"math/rand"
	"sync"
	"url-shortener/cmd/api/domain"
)

type urlShortenerService struct {
	urlMap map[string]string
	mutex  sync.Mutex
}

func NewURLShortenerService() URLShortenerService {
	return &urlShortenerService{
		urlMap: make(map[string]string),
	}
}

func (s *urlShortenerService) ShortenURL(originalURL string) domain.URLMapping {
	shortURL := generateShortURL()

	s.mutex.Lock()
	s.urlMap[shortURL] = originalURL
	s.mutex.Unlock()

	return domain.URLMapping{
		OriginalURL: originalURL,
		ShortURL:    "http://1.unli.ink/s/" + shortURL,
	}
}

func (s *urlShortenerService) GetOriginalURL(shortURL string) (string, bool) {
	s.mutex.Lock()

	originalURL, ok := s.urlMap[shortURL]

	s.mutex.Unlock()

	return originalURL, ok
}

func generateShortURL() string {
	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	b := make([]byte, 6)

	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}

	return string(b)
}

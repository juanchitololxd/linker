package services

import (
	"math/rand"
	"url-shortener/cmd/api/domain"
)

type urlShortenerService struct {
	URLRepository URLRepository
}

func NewURLShortenerService() URLShortenerService {
	return &urlShortenerService{}
}

func (s *urlShortenerService) ShortenURL(originalURL string) (domain.URLMapping, error) {
	var urlMap domain.URLMapping

	savedURL, err := s.URLRepository.FindByOriginalURL(originalURL)
	if err != nil {
		return urlMap, err
	}

	if savedURL.OriginalURL != "" && savedURL.ShortURL != "" {
		return savedURL, nil
	}

	shortURL := generateShortURL()
	urlMap.OriginalURL = originalURL
	urlMap.ShortURL = "http://1.unli.ink/s/" + shortURL

	err = s.URLRepository.Save(urlMap)
	if err != nil {
		return domain.URLMapping{}, err
	}

	return urlMap, nil
}

func (s *urlShortenerService) GetOriginalURL(shortURL string) (string, error) {
	url, err := s.URLRepository.FindByOriginalURL(shortURL)
	if err != nil {
		return "", err
	}

	return url.OriginalURL, nil
}

func (s *urlShortenerService) GetHistory() ([]domain.URLMapping, error) {
	urls, err := s.URLRepository.FindAll()
	if err != nil {
		return nil, err
	}

	return urls, nil
}

func generateShortURL() string {
	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	b := make([]byte, 6)

	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}

	return string(b)
}

func (s *urlShortenerService) GetPing() string {
	return "pong"
}

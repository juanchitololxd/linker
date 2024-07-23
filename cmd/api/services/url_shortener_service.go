package services

import (
	"math/rand"
	"os"
	"url-shortener/cmd/api/domain"
)

type urlShortenerService struct {
	URLRepository URLRepository
}

func NewURLShortenerService(repository URLRepository) URLShortenerService {
	return &urlShortenerService{
		URLRepository: repository,
	}
}

func (s *urlShortenerService) ShortenURL(originalURL string) (domain.URLMapping, error) {
	var urlMap domain.URLMapping

	savedURL, err := s.URLRepository.FindByOriginalURL(originalURL)
	if err != nil {
		return urlMap, err
	}

	baseURL := os.Getenv("BASE_URL")

	if savedURL.OriginalURL != "" && savedURL.ShortURL != "" {
		saved := baseURL + "/s/" + savedURL.ShortURL
		savedURL.ShortURL = saved
		return savedURL, nil
	}

	shortURL := generateShortURL()
	urlMap.OriginalURL = originalURL
	urlMap.ShortURL = shortURL

	err = s.URLRepository.Save(urlMap)
	if err != nil {
		return domain.URLMapping{}, err
	}

	urlMap.ShortURL = baseURL + "/s/" + shortURL

	return urlMap, nil
}

func (s *urlShortenerService) GetOriginalURL(shortURL string) (string, error) {
	url, err := s.URLRepository.FindByShortURL(shortURL)

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

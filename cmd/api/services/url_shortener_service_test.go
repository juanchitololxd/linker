package services

import (
	"testing"
	_ "url-shortener/cmd/api/domain"
)

func TestShortenURL(t *testing.T) {
	service := NewURLShortenerService()

	originalURL := "http://example.com"
	urlMapping := service.ShortenURL(originalURL)

	if urlMapping.OriginalURL != originalURL {
		t.Errorf("expected %s, got %s", originalURL, urlMapping.OriginalURL)
	}

	if len(urlMapping.ShortURL) == 0 {
		t.Errorf("expected a short URL, got an empty string")
	}

	if urlMapping.ShortURL[:20] != "http://1.unli.ink/s/" {
		t.Errorf("expected short URL to start with http://1.unli.ink/s/, got %s", urlMapping.ShortURL[:21])
	}
}

func TestGetOriginalURL(t *testing.T) {
	service := NewURLShortenerService()

	originalURL := "http://example.com"
	urlMapping := service.ShortenURL(originalURL)

	shortURL := urlMapping.ShortURL[20:] // extract the short part of the URL
	retrievedURL, ok := service.GetOriginalURL(shortURL)

	if !ok {
		t.Errorf("expected to retrieve original URL, but got false")
	}

	if retrievedURL != originalURL {
		t.Errorf("expected %s, got %s", originalURL, retrievedURL)
	}

	// Test retrieving a non-existent URL
	_, ok = service.GetOriginalURL("nonexistent")
	if ok {
		t.Errorf("expected to not retrieve a URL, but got true")
	}
}

func TestGenerateShortURL(t *testing.T) {
	shortURL := generateShortURL()

	if len(shortURL) != 6 {
		t.Errorf("expected short URL of length 6, got %d", len(shortURL))
	}

	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	for _, char := range shortURL {
		if !contains(letters, char) {
			t.Errorf("expected character in short URL to be one of %s, but got %c", letters, char)
		}
	}
}

func contains(str string, char rune) bool {
	for _, c := range str {
		if c == char {
			return true
		}
	}
	return false
}

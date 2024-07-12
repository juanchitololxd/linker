package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"url-shortener/cmd/api/domain"
	_ "url-shortener/cmd/api/services"
)

type mockURLShortenerService struct {
	urlMap map[string]string
}

func (m *mockURLShortenerService) ShortenURL(originalURL string) domain.URLMapping {
	shortURL := "abc123"
	m.urlMap[shortURL] = originalURL
	return domain.URLMapping{
		OriginalURL: originalURL,
		ShortURL:    "http://1.unli.ink/s/" + shortURL,
	}
}

func (m *mockURLShortenerService) GetOriginalURL(shortURL string) (string, bool) {
	originalURL, ok := m.urlMap[shortURL]
	return originalURL, ok
}

func TestShortenURLHandler(t *testing.T) {
	mockService := &mockURLShortenerService{urlMap: make(map[string]string)}
	handler := NewURLHandler(mockService)

	t.Run("Valid request", func(t *testing.T) {
		urlMapping := domain.URLMapping{OriginalURL: "http://example.com"}
		body, _ := json.Marshal(urlMapping)
		req, err := http.NewRequest(http.MethodPost, "/shorten", bytes.NewBuffer(body))
		if err != nil {
			t.Fatal(err)
		}
		rr := httptest.NewRecorder()
		http.HandlerFunc(handler.ShortenURLHandler).ServeHTTP(rr, req)

		if status := rr.Code; status != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
		}

		var got domain.URLMapping
		if err := json.NewDecoder(rr.Body).Decode(&got); err != nil {
			t.Errorf("could not decode response: %v", err)
		}

		expected := domain.URLMapping{
			OriginalURL: "http://example.com",
			ShortURL:    "http://1.unli.ink/s/abc123",
		}

		if got != expected {
			t.Errorf("handler returned unexpected body: got %v want %v", got, expected)
		}
	})

	t.Run("Invalid request method", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodGet, "/shorten", nil)
		if err != nil {
			t.Fatal(err)
		}
		rr := httptest.NewRecorder()
		http.HandlerFunc(handler.ShortenURLHandler).ServeHTTP(rr, req)

		if status := rr.Code; status != http.StatusMethodNotAllowed {
			t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusMethodNotAllowed)
		}

		expected := "Invalid request method\n"
		if rr.Body.String() != expected {
			t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
		}
	})

	t.Run("Invalid request body", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodPost, "/shorten", bytes.NewBuffer([]byte("invalid body")))
		if err != nil {
			t.Fatal(err)
		}
		rr := httptest.NewRecorder()
		http.HandlerFunc(handler.ShortenURLHandler).ServeHTTP(rr, req)

		if status := rr.Code; status != http.StatusBadRequest {
			t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusBadRequest)
		}

		expected := "Invalid request body\n"
		if rr.Body.String() != expected {
			t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
		}
	})
}

func TestRedirectHandler(t *testing.T) {
	mockService := &mockURLShortenerService{urlMap: make(map[string]string)}
	handler := NewURLHandler(mockService)
	mockService.ShortenURL("http://example.com")

	t.Run("Valid short URL", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodGet, "/s/abc123", nil)
		if err != nil {
			t.Fatal(err)
		}
		rr := httptest.NewRecorder()
		http.HandlerFunc(handler.RedirectHandler).ServeHTTP(rr, req)

		if status := rr.Code; status != http.StatusFound {
			t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusFound)
		}

		expected := "http://example.com"
		if location := rr.Header().Get("Location"); location != expected {
			t.Errorf("handler returned wrong location header: got %v want %v", location, expected)
		}
	})

	t.Run("Invalid short URL", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodGet, "/s/nonexistent", nil)
		if err != nil {
			t.Fatal(err)
		}
		rr := httptest.NewRecorder()
		http.HandlerFunc(handler.RedirectHandler).ServeHTTP(rr, req)

		if status := rr.Code; status != http.StatusNotFound {
			t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusNotFound)
		}
	})
}

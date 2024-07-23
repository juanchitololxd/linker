package handlers

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"url-shortener/cmd/api/domain"
	_ "url-shortener/cmd/api/services"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Mocking the URLShortenerService
type MockURLShortenerService struct {
	mock.Mock
}

func (m *MockURLShortenerService) ShortenURL(originalURL string) (domain.URLMapping, error) {
	args := m.Called(originalURL)
	return args.Get(0).(domain.URLMapping), args.Error(1)
}

func (m *MockURLShortenerService) GetOriginalURL(shortURL string) (string, error) {
	args := m.Called(shortURL)
	return args.String(0), args.Error(1)
}

func (m *MockURLShortenerService) GetHistory() ([]domain.URLMapping, error) {
	args := m.Called()
	return args.Get(0).([]domain.URLMapping), args.Error(1)
}

func (m *MockURLShortenerService) GetPing() string {
	args := m.Called()
	return args.String(0)
}

func TestShortenURLHandler(t *testing.T) {
	mockService := new(MockURLShortenerService)
	handler := NewURLHandler(mockService)

	urlMapping := domain.URLMapping{OriginalURL: "http://example.com"}
	shortenedURL := domain.URLMapping{OriginalURL: "http://example.com", ShortURL: "abcdef"}
	mockService.On("ShortenURL", urlMapping.OriginalURL).Return(shortenedURL, nil)

	body, _ := json.Marshal(urlMapping)
	req := httptest.NewRequest(http.MethodPost, "/shorten", bytes.NewBuffer(body))
	w := httptest.NewRecorder()

	handler.ShortenURLHandler(w, req)

	resp := w.Result()
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	var result domain.URLMapping
	json.NewDecoder(resp.Body).Decode(&result)
	assert.Equal(t, shortenedURL, result)
}

func TestShortenURLHandler_InvalidMethod(t *testing.T) {
	mockService := new(MockURLShortenerService)
	handler := NewURLHandler(mockService)

	req := httptest.NewRequest(http.MethodGet, "/shorten", nil)
	w := httptest.NewRecorder()

	handler.ShortenURLHandler(w, req)

	resp := w.Result()
	assert.Equal(t, http.StatusMethodNotAllowed, resp.StatusCode)
}

func TestShortenURLHandler_InvalidBody(t *testing.T) {
	mockService := new(MockURLShortenerService)
	handler := NewURLHandler(mockService)

	req := httptest.NewRequest(http.MethodPost, "/shorten", bytes.NewBuffer([]byte("invalid body")))
	w := httptest.NewRecorder()

	handler.ShortenURLHandler(w, req)

	resp := w.Result()
	assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
}

func TestShortenURLHandler_Error(t *testing.T) {
	mockService := new(MockURLShortenerService)
	handler := NewURLHandler(mockService)

	urlMapping := domain.URLMapping{OriginalURL: "http://example.com"}
	mockService.On("ShortenURL", urlMapping.OriginalURL).Return(domain.URLMapping{}, errors.New("service error"))

	body, _ := json.Marshal(urlMapping)
	req := httptest.NewRequest(http.MethodPost, "/shorten", bytes.NewBuffer(body))
	w := httptest.NewRecorder()

	handler.ShortenURLHandler(w, req)

	resp := w.Result()
	assert.Equal(t, http.StatusInternalServerError, resp.StatusCode)
}

func TestRedirectHandler(t *testing.T) {
	mockService := new(MockURLShortenerService)
	handler := NewURLHandler(mockService)

	shortURL := "abcdef"
	originalURL := "http://example.com"
	mockService.On("GetOriginalURL", shortURL).Return(originalURL, nil)

	req := httptest.NewRequest(http.MethodGet, "/s/"+shortURL, nil)
	w := httptest.NewRecorder()

	handler.RedirectHandler(w, req)

	resp := w.Result()
	assert.Equal(t, http.StatusFound, resp.StatusCode)
	assert.Equal(t, originalURL, resp.Header.Get("Location"))
}

func TestRedirectHandler_Error(t *testing.T) {
	mockService := new(MockURLShortenerService)
	handler := NewURLHandler(mockService)

	shortURL := "abcdef"
	mockService.On("GetOriginalURL", shortURL).Return("", errors.New("service error"))

	req := httptest.NewRequest(http.MethodGet, "/s/"+shortURL, nil)
	w := httptest.NewRecorder()

	handler.RedirectHandler(w, req)

	resp := w.Result()
	assert.Equal(t, http.StatusNotFound, resp.StatusCode)
}

func TestGetHistory(t *testing.T) {
	mockService := new(MockURLShortenerService)
	handler := NewURLHandler(mockService)

	history := []domain.URLMapping{
		{OriginalURL: "http://example1.com", ShortURL: "abc123"},
		{OriginalURL: "http://example2.com", ShortURL: "def456"},
	}
	mockService.On("GetHistory").Return(history, nil)

	req := httptest.NewRequest(http.MethodGet, "/history", nil)
	w := httptest.NewRecorder()

	handler.GetHistory(w, req)

	resp := w.Result()
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	var result []domain.URLMapping
	json.NewDecoder(resp.Body).Decode(&result)
	assert.Equal(t, history, result)
}

func TestGetHistory_InvalidMethod(t *testing.T) {
	mockService := new(MockURLShortenerService)
	handler := NewURLHandler(mockService)

	req := httptest.NewRequest(http.MethodPost, "/history", nil)
	w := httptest.NewRecorder()

	handler.GetHistory(w, req)

	resp := w.Result()
	assert.Equal(t, http.StatusMethodNotAllowed, resp.StatusCode)
}

func TestGetHistory_Error(t *testing.T) {
	mockService := new(MockURLShortenerService)
	handler := NewURLHandler(mockService)

	// Devuelve un slice vacío en lugar de nil
	mockService.On("GetHistory").Return([]domain.URLMapping{}, errors.New("service error"))

	req := httptest.NewRequest(http.MethodGet, "/history", nil)
	w := httptest.NewRecorder()

	handler.GetHistory(w, req)

	resp := w.Result()
	assert.Equal(t, http.StatusInternalServerError, resp.StatusCode)
}

func TestGetPing(t *testing.T) {
	mockService := new(MockURLShortenerService)
	handler := NewURLHandler(mockService)

	mockService.On("GetPing").Return("pong")

	req := httptest.NewRequest(http.MethodGet, "/ping", nil)
	w := httptest.NewRecorder()

	handler.GetPing(w, req)

	resp := w.Result()
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	var result string
	json.NewDecoder(resp.Body).Decode(&result)
	assert.Equal(t, "pong", result)
}

func TestGetPing_InvalidMethod(t *testing.T) {
	mockService := new(MockURLShortenerService)
	handler := NewURLHandler(mockService)

	req := httptest.NewRequest(http.MethodPost, "/ping", nil)
	w := httptest.NewRecorder()

	handler.GetPing(w, req)

	resp := w.Result()
	assert.Equal(t, http.StatusMethodNotAllowed, resp.StatusCode)
}

package services

import (
	"errors"
	"os"
	"testing"
	"url-shortener/cmd/api/domain"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Mocking the URLRepository
type MockURLRepository struct {
	mock.Mock
}

func (m *MockURLRepository) Save(url domain.URLMapping) error {
	args := m.Called(url)
	return args.Error(0)
}

func (m *MockURLRepository) FindByOriginalURL(originalURL string) (domain.URLMapping, error) {
	args := m.Called(originalURL)
	return args.Get(0).(domain.URLMapping), args.Error(1)
}

func (m *MockURLRepository) FindByShortURL(shortURL string) (domain.URLMapping, error) {
	args := m.Called(shortURL)
	return args.Get(0).(domain.URLMapping), args.Error(1)
}

func (m *MockURLRepository) FindAll() ([]domain.URLMapping, error) {
	args := m.Called()
	return args.Get(0).([]domain.URLMapping), args.Error(1)
}

func TestShortenURL(t *testing.T) {
	mockRepo := new(MockURLRepository)
	service := NewURLShortenerService(mockRepo)

	originalURL := "http://example.com"
	//shortURL := "abcdef"

	mockRepo.On("FindByOriginalURL", originalURL).Return(domain.URLMapping{}, nil)
	mockRepo.On("Save", mock.Anything).Return(nil)

	os.Setenv("BASE_URL", "http://short.url")
	result, err := service.ShortenURL(originalURL)
	assert.NoError(t, err)
	assert.Equal(t, result.ShortURL, result.ShortURL)

	mockRepo.AssertExpectations(t)
}

func TestGetOriginalURL(t *testing.T) {
	mockRepo := new(MockURLRepository)
	service := NewURLShortenerService(mockRepo)

	shortURL := "abcdef"
	originalURL := "http://example.com"

	mockRepo.On("FindByShortURL", shortURL).Return(domain.URLMapping{OriginalURL: originalURL}, nil)

	result, err := service.GetOriginalURL(shortURL)
	assert.NoError(t, err)
	assert.Equal(t, originalURL, result)

	mockRepo.AssertExpectations(t)
}

func TestGetHistory(t *testing.T) {
	mockRepo := new(MockURLRepository)
	service := NewURLShortenerService(mockRepo)

	history := []domain.URLMapping{
		{OriginalURL: "http://example1.com", ShortURL: "abc123"},
		{OriginalURL: "http://example2.com", ShortURL: "def456"},
	}

	mockRepo.On("FindAll").Return(history, nil)

	result, err := service.GetHistory()
	assert.NoError(t, err)
	assert.Equal(t, history, result)

	mockRepo.AssertExpectations(t)
}

func TestGetPing(t *testing.T) {
	service := &urlShortenerService{}
	result := service.GetPing()
	assert.Equal(t, "pong", result)
}

func TestShortenURL_ExistingURL(t *testing.T) {
	mockRepo := new(MockURLRepository)
	service := NewURLShortenerService(mockRepo)

	originalURL := "http://example.com"
	shortURL := "abcdef"
	baseURL := "http://short.url"

	mockRepo.On("FindByOriginalURL", originalURL).Return(domain.URLMapping{OriginalURL: originalURL, ShortURL: shortURL}, nil)

	os.Setenv("BASE_URL", baseURL)
	result, err := service.ShortenURL(originalURL)
	assert.NoError(t, err)
	assert.Equal(t, baseURL+"/s/"+shortURL, result.ShortURL)

	mockRepo.AssertExpectations(t)
}

func TestShortenURL_SaveError(t *testing.T) {
	mockRepo := new(MockURLRepository)
	service := NewURLShortenerService(mockRepo)

	originalURL := "http://example.com"
	//shortURL := "abcdef"

	mockRepo.On("FindByOriginalURL", originalURL).Return(domain.URLMapping{}, nil)
	mockRepo.On("Save", mock.Anything).Return(errors.New("save error"))

	_, err := service.ShortenURL(originalURL)
	assert.Error(t, err)
	assert.Equal(t, "save error", err.Error())

	mockRepo.AssertExpectations(t)
}

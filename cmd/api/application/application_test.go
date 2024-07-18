package application

import (
	_ "embed"
	"encoding/json"
	"testing"
	"url-shortener/cmd/api/domain"
	"url-shortener/cmd/api/handlers"
	"url-shortener/cmd/api/services"
)

//go:embed test_data/data.json
var testData []byte

func TestInitialize(t *testing.T) {
	Initialize()

	if URLService == nil {
		t.Errorf("URLService is nil; expected non-nil")
	}

	if URLHandler == nil {
		t.Errorf("URLHandler is nil; expected non-nil")
	}

	// Additional checks to ensure the correct types are assigned
	if _, ok := URLService.(services.URLShortenerService); !ok {
		t.Errorf("URLService is of the wrong type")
	}

	if _, ok := URLHandler.(handlers.URLHandler); !ok {
		t.Errorf("URLHandler is of the wrong type")
	}
}

func TestURLServiceInitialization(t *testing.T) {
	Initialize()

	if URLService == nil {
		t.Fatal("expected URLService to be initialized, got nil")
	}

	var urlMapping domain.URLMapping
	if err := json.Unmarshal(testData, &urlMapping); err != nil {
		t.Fatalf("failed to unmarshal test data: %v", err)
	}

	result := URLService.ShortenURL(urlMapping.OriginalURL)
	if result.OriginalURL != urlMapping.OriginalURL {
		t.Errorf("expected OriginalURL to be '%s', got '%s'", urlMapping.OriginalURL, result.OriginalURL)
	}

	if len(result.ShortURL) == 0 {
		t.Errorf("expected ShortURL to be non-empty")
	}
}

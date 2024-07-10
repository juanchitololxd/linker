package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

// resetURLMap is a helper function to reset the global urlMap and mutex for testing purposes
func resetURLMap() {
	mutex.Lock()
	urlMap = make(map[string]string)
	mutex.Unlock()
}

// TestShortenURLHandler tests the shortenURLHandler function
func TestShortenURLHandler(t *testing.T) {
	resetURLMap()
	originalURL := "http://example.com"
	body, _ := json.Marshal(URLMapping{OriginalURL: originalURL})
	req, err := http.NewRequest(http.MethodPost, "/shorten", bytes.NewBuffer(body))
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(shortenURLHandler)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	var response URLMapping
	err = json.NewDecoder(rr.Body).Decode(&response)
	if err != nil {
		t.Fatal(err)
	}

	if response.OriginalURL != originalURL {
		t.Errorf("handler returned unexpected body: got %v want %v", response.OriginalURL, originalURL)
	}

	if response.ShortURL == "" {
		t.Error("handler returned empty short URL")
	}

	mutex.Lock()
	if urlMap[response.ShortURL[len("http://1.unli.ink/s/"):]] != originalURL {
		t.Error("short URL not mapped correctly to original URL")
	}
	mutex.Unlock()
}

// TestShortenURLHandlerInvalidMethod tests the shortenURLHandler function with an invalid method
func TestShortenURLHandlerInvalidMethod(t *testing.T) {
	req, err := http.NewRequest(http.MethodGet, "/shorten", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(shortenURLHandler)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusMethodNotAllowed {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusMethodNotAllowed)
	}
}

// TestShortenURLHandlerInvalidBody tests the shortenURLHandler function with an invalid request body
func TestShortenURLHandlerInvalidBody(t *testing.T) {
	req, err := http.NewRequest(http.MethodPost, "/shorten", bytes.NewBufferString("invalid body"))
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(shortenURLHandler)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusBadRequest)
	}
}

// TestRedirectHandler tests the redirectHandler function
func TestRedirectHandler(t *testing.T) {
	resetURLMap()
	shortURL := "abc123"
	originalURL := "http://example.com"
	mutex.Lock()
	urlMap[shortURL] = originalURL
	mutex.Unlock()

	req, err := http.NewRequest(http.MethodGet, "/s/"+shortURL, nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(redirectHandler)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusFound {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusFound)
	}

	if location := rr.Header().Get("Location"); location != originalURL {
		t.Errorf("handler returned wrong location header: got %v want %v", location, originalURL)
	}
}

// TestRedirectHandlerNotFound tests the redirectHandler function when the short URL does not exist
func TestRedirectHandlerNotFound(t *testing.T) {
	resetURLMap()
	req, err := http.NewRequest(http.MethodGet, "/s/nonexistent", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(redirectHandler)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusNotFound {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusNotFound)
	}
}

// TestGenerateShortURL tests the generateShortURL function
func TestGenerateShortURL(t *testing.T) {
	shortURL := generateShortURL()
	if len(shortURL) != 6 {
		t.Errorf("generateShortURL returned wrong length: got %v want %v", len(shortURL), 6)
	}
	for _, char := range shortURL {
		if (char < 'a' || char > 'z') && (char < 'A' || char > 'Z') && (char < '0' || char > '9') {
			t.Errorf("generateShortURL returned invalid character: %v", char)
		}
	}
}

package handlers

import (
	"encoding/json"
	"net/http"
	"url-shortener/cmd/api/domain"
	"url-shortener/cmd/api/services"
)

type shortenURLHandler struct {
	service services.URLShortenerService
}

func NewURLHandler(service services.URLShortenerService) URLHandler {
	return &shortenURLHandler{service: service}
}

func (h *shortenURLHandler) ShortenURLHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var urlMapping domain.URLMapping

	if err := json.NewDecoder(r.Body).Decode(&urlMapping); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	urlMapping = h.service.ShortenURL(urlMapping.OriginalURL)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(urlMapping)
}

func (h *shortenURLHandler) RedirectHandler(w http.ResponseWriter, r *http.Request) {
	shortURL := r.URL.Path[len("/s/"):]
	originalURL, ok := h.service.GetOriginalURL(shortURL)

	if !ok {
		http.NotFound(w, r)
		return
	}

	http.Redirect(w, r, originalURL, http.StatusFound)
}

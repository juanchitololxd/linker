package handlers

import (
	"encoding/json"
	"log"
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
		log.Println("ERROR: Invalid request method")
		return
	}

	var urlMapping domain.URLMapping

	if err := json.NewDecoder(r.Body).Decode(&urlMapping); err != nil {
		log.Println("ERROR: Invalid request body")
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	urlMapping, err := h.service.ShortenURL(urlMapping.OriginalURL)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	log.Print("INFO: Shorten URL generated")

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(urlMapping)
}

func (h *shortenURLHandler) RedirectHandler(w http.ResponseWriter, r *http.Request) {
	shortURL := r.URL.Path[len("/s/"):]
	originalURL, err := h.service.GetOriginalURL(shortURL)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	log.Println("INFO: Redirecting to", originalURL)
	http.Redirect(w, r, originalURL, http.StatusFound)
}

func (h *shortenURLHandler) GetHistory(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		log.Println("ERROR: Invalid request method")
		return
	}

	urlMapping, err := h.service.GetHistory()
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	log.Print("INFO: History generated")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(urlMapping)
}

func (h *shortenURLHandler) GetPing(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		log.Println("ERROR: Invalid request method")
		return
	}

	urlMapping := h.service.GetPing()

	log.Print("INFO: ping call")

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(urlMapping)
}

package handlers

import (
	"net/http"
)

type URLHandler interface {
	ShortenURLHandler(w http.ResponseWriter, r *http.Request)
	RedirectHandler(w http.ResponseWriter, r *http.Request)
}

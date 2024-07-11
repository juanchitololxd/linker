package application

import (
	"url-shortener/cmd/api/handlers"
	"url-shortener/cmd/api/services"
)

var (
	URLService services.URLShortenerService
	URLHandler handlers.URLHandler
)

func Initialize() {
	// Init Services
	URLService = services.NewURLShortenerService()

	// Init Handlers
	URLHandler = handlers.NewURLHandler(URLService)
}

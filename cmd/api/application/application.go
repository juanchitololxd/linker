package application

import (
	"url-shortener/cmd/api/handlers"
	"url-shortener/cmd/api/services"
	"github.com/joho/godotenv"
	"log"
)

var (
	URLService services.URLShortenerService
	URLHandler handlers.URLHandler
)

func Initialize() {
	// Init Services
	godotenv.Load()
	log.Println("Variables de entorno cargadas")
	URLService = services.NewURLShortenerService()

	// Init Handlers
	URLHandler = handlers.NewURLHandler(URLService)
}

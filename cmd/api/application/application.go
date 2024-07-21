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
    if err := godotenv.Load(".env"); err != nil {
        log.Fatalf("Error loading .env file")
    }
	log.Print("Variables de entorno cargadas")
	// Init Service
	URLService = services.NewURLShortenerService()

	// Init Handler
	URLHandler = handlers.NewURLHandler(URLService)
}

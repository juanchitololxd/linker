package application

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"log"
	"os"
	"url-shortener/cmd/api/handlers"
	"url-shortener/cmd/api/services"
)

var (
	URLService services.URLShortenerService
	URLHandler handlers.URLHandler
	loadEnv    = godotenv.Load
	sqlOpen    = sql.Open
)

func Initialize() {
	if err := loadEnv(".env"); err != nil {
		log.Fatalf("Error loading .env file")
	}
	log.Print("Variables de entorno cargadas")

	// Init Database
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	dbUser := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWD")

	db, err := sqlOpen("mysql", dbUser+":"+dbPassword+"@tcp("+dbHost+":"+dbPort+")/"+dbName)
	if err != nil {
		log.Fatal(err)
	}

	// Init Repository
	URLRepository := services.NewURLRepository(db)

	// Init Service
	URLService = services.NewURLShortenerService(URLRepository)

	// Init Handler
	URLHandler = handlers.NewURLHandler(URLService)
}

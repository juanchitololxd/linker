package main

import (
	"log"
	"net/http"
	"url-shortener/cmd/api/application"
)

func main() {
	application.Initialize()

	http.Handle("/", http.FileServer(http.Dir("./cmd/api/static")))
	http.HandleFunc("/shorten", application.URLHandler.ShortenURLHandler)
	http.HandleFunc("/s/", application.URLHandler.RedirectHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

package main

import (
    "encoding/json"
    "log"
    "math/rand"
    "net/http"
    "sync"
    "time"
)

type URLMapping struct {
    OriginalURL string `json:"original_url"`
    ShortURL    string `json:"short_url"`
}

var (
    urlMap = make(map[string]string)
    mutex  sync.Mutex
)

func main() {
    rand.Seed(time.Now().UnixNano())
    http.Handle("/", http.FileServer(http.Dir("./static")))
    http.HandleFunc("/shorten", shortenURLHandler)
    http.HandleFunc("/s/", redirectHandler)
    log.Fatal(http.ListenAndServe(":8080", nil))
}

func shortenURLHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
        return
    }
    var urlMapping URLMapping
    if err := json.NewDecoder(r.Body).Decode(&urlMapping); err != nil {
        http.Error(w, "Invalid request body", http.StatusBadRequest)
        return
    }
    shortURL := generateShortURL()
    mutex.Lock()
    urlMap[shortURL] = urlMapping.OriginalURL
    mutex.Unlock()
    urlMapping.ShortURL = "http://1.unli.ink/s/" + shortURL
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(urlMapping)
}

func redirectHandler(w http.ResponseWriter, r *http.Request) {
    shortURL := r.URL.Path[len("/s/"):]
    mutex.Lock()
    originalURL, ok := urlMap[shortURL]
    mutex.Unlock()
    if !ok {
        http.NotFound(w, r)
        return
    }
    http.Redirect(w, r, originalURL, http.StatusFound)
}

func generateShortURL() string {
    const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
    b := make([]byte, 6)
    for i := range b {
        b[i] = letters[rand.Intn(len(letters))]
    }
    
    return string(b)
}

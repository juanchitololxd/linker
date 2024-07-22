package services

import (
	"database/sql"
	"fmt"
	"url-shortener/cmd/api/domain"
)

type URLRepo struct {
	URLRepository *sql.DB
}

func NewURLRepository(urlRepository *sql.DB) *URLRepo {
	return &URLRepo{URLRepository: urlRepository}
}

func (r *URLRepo) Save(url domain.URLMapping) error {
	_, err := r.URLRepository.Exec("INSERT INTO urls (original_url, shorten_url) VALUES (?, ?)", url.OriginalURL, url.ShortURL)
	if err != nil {
		return fmt.Errorf("ERROR: Error al insertar la URL '%v' - '%v'\n", url.OriginalURL, url.ShortURL)
	}

	return nil
}

func (r *URLRepo) FindByOriginalURL(originalURL string) (domain.URLMapping, error) {
	var url domain.URLMapping

	err := r.URLRepository.QueryRow("SELECT urls.original_url, urls.shorten_url FROM urls WHERE original_url = ?", originalURL).Scan(&url.OriginalURL, &url.ShortURL)
	if err != nil {
		return domain.URLMapping{}, fmt.Errorf("ERROR: Error al consultar la URL '%v'\n", originalURL)
	}

	return url, nil
}

func (r *URLRepo) FindByShortURL(shortURL string) (domain.URLMapping, error) {
	var url domain.URLMapping

	err := r.URLRepository.QueryRow("SELECT urls.original_url, urls.shorten_url FROM urls WHERE shorten_url = ?", shortURL).Scan(&url.OriginalURL, &url.ShortURL)
	if err != nil {
		return domain.URLMapping{}, fmt.Errorf("ERROR: Error al consultar la URL '%v'\n", shortURL)
	}

	return url, nil
}

func (r *URLRepo) FindAll() ([]domain.URLMapping, error) {
	rows, err := r.URLRepository.Query("SELECT original_url, shorten_url FROM urls")
	if err != nil {
		return nil, fmt.Errorf("ERROR: Error al consultar la URL '%v'\n", err)
	}
	defer rows.Close()

	var urls []domain.URLMapping

	for rows.Next() {
		var url domain.URLMapping
		var originalURL string
		var shortenURL string

		err = rows.Scan(&originalURL, &shortenURL)
		if err != nil {
			return nil, fmt.Errorf("ERROR: Error al consultar la URL '%v'\n", err)
		}

		url.OriginalURL = originalURL
		url.ShortURL = shortenURL

		urls = append(urls, url)
	}

	return urls, nil
}

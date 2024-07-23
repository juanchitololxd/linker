package services

import (
	"database/sql"
	"errors"
	"testing"
	"url-shortener/cmd/api/domain"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestURLRepo_Save(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := NewURLRepository(db)

	url := domain.URLMapping{
		OriginalURL: "http://example.com",
		ShortURL:    "abcdef",
	}

	mock.ExpectExec("INSERT INTO urls").
		WithArgs(url.OriginalURL, url.ShortURL).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err = repo.Save(url)
	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestURLRepo_Save_Error(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := NewURLRepository(db)

	url := domain.URLMapping{
		OriginalURL: "http://example.com",
		ShortURL:    "abcdef",
	}

	mock.ExpectExec("INSERT INTO urls").
		WithArgs(url.OriginalURL, url.ShortURL).
		WillReturnError(errors.New("insert error"))

	err = repo.Save(url)
	assert.Error(t, err)
	assert.Equal(t, "ERROR: Error al insertar la URL 'http://example.com' - 'abcdef'\n", err.Error())
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestURLRepo_FindByOriginalURL(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := NewURLRepository(db)

	originalURL := "http://example.com"
	shortURL := "abcdef"

	rows := sqlmock.NewRows([]string{"original_url", "shorten_url"}).
		AddRow(originalURL, shortURL)

	mock.ExpectQuery("SELECT urls.original_url, urls.shorten_url FROM urls WHERE original_url = ?").
		WithArgs(originalURL).
		WillReturnRows(rows)

	result, err := repo.FindByOriginalURL(originalURL)
	assert.NoError(t, err)
	assert.Equal(t, originalURL, result.OriginalURL)
	assert.Equal(t, shortURL, result.ShortURL)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestURLRepo_FindByOriginalURL_NoRows(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := NewURLRepository(db)

	originalURL := "http://example.com"

	mock.ExpectQuery("SELECT urls.original_url, urls.shorten_url FROM urls WHERE original_url = ?").
		WithArgs(originalURL).
		WillReturnError(sql.ErrNoRows)

	result, err := repo.FindByOriginalURL(originalURL)
	assert.NoError(t, err)
	assert.Equal(t, domain.URLMapping{}, result)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestURLRepo_FindByOriginalURL_Error(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := NewURLRepository(db)

	originalURL := "http://example.com"

	mock.ExpectQuery("SELECT urls.original_url, urls.shorten_url FROM urls WHERE original_url = ?").
		WithArgs(originalURL).
		WillReturnError(errors.New("query error"))

	_, err = repo.FindByOriginalURL(originalURL)
	assert.Error(t, err)
	assert.Equal(t, "ERROR: Error al consultar la URL 'http://example.com'\n", err.Error())
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestURLRepo_FindByShortURL(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := NewURLRepository(db)

	originalURL := "http://example.com"
	shortURL := "abcdef"

	rows := sqlmock.NewRows([]string{"original_url", "shorten_url"}).
		AddRow(originalURL, shortURL)

	mock.ExpectQuery("SELECT urls.original_url, urls.shorten_url FROM urls WHERE shorten_url = ?").
		WithArgs(shortURL).
		WillReturnRows(rows)

	result, err := repo.FindByShortURL(shortURL)
	assert.NoError(t, err)
	assert.Equal(t, originalURL, result.OriginalURL)
	assert.Equal(t, shortURL, result.ShortURL)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestURLRepo_FindByShortURL_Error(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := NewURLRepository(db)

	shortURL := "abcdef"

	mock.ExpectQuery("SELECT urls.original_url, urls.shorten_url FROM urls WHERE shorten_url = ?").
		WithArgs(shortURL).
		WillReturnError(errors.New("query error"))

	_, err = repo.FindByShortURL(shortURL)
	assert.Error(t, err)
	assert.Equal(t, "ERROR: Error al consultar la URL 'abcdef'\n", err.Error())
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestURLRepo_FindAll(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := NewURLRepository(db)

	rows := sqlmock.NewRows([]string{"original_url", "shorten_url"}).
		AddRow("http://example1.com", "abc123").
		AddRow("http://example2.com", "def456")

	mock.ExpectQuery("SELECT original_url, shorten_url FROM urls").
		WillReturnRows(rows)

	result, err := repo.FindAll()
	assert.NoError(t, err)
	assert.Len(t, result, 2)
	assert.Equal(t, "http://example1.com", result[0].OriginalURL)
	assert.Equal(t, "abc123", result[0].ShortURL)
	assert.Equal(t, "http://example2.com", result[1].OriginalURL)
	assert.Equal(t, "def456", result[1].ShortURL)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestURLRepo_FindAll_Error(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := NewURLRepository(db)

	mock.ExpectQuery("SELECT original_url, shorten_url FROM urls").
		WillReturnError(errors.New("query error"))

	_, err = repo.FindAll()
	assert.Error(t, err)
	assert.Equal(t, "ERROR: Error al consultar la URL 'query error'\n", err.Error())
	assert.NoError(t, mock.ExpectationsWereMet())
}

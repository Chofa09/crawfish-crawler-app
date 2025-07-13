package models

import (
	"database/sql"
	"encoding/json"
	"time"
)

type URL struct {
	ID           int             `json:"id"`
	Address      string          `json:"address"`
	PageTitle    string          `json:"page_title"`
	HTMLVersion  string          `json:"html_version"`
	Headings     json.RawMessage `json:"headings"`
	HasLoginForm bool            `json:"has_login_form"`
	CrawlStatus  string          `json:"crawl_status"`
	CreatedAt    time.Time       `json:"created_at"`
}

// Returns all urls from the db
func GetAllURLs(db *sql.DB) ([]URL, error) {
	rows, err := db.Query(`
		SELECT id, address, page_title, html_version, headings, has_login_form, crawl_status, created_at
		FROM urls
	`)
	if err != nil {
		return nil, err
	}

	// delays execution of Close until this function exits - cleans up resource
	defer rows.Close()

	var urls []URL

	for rows.Next() {
		var url URL
		if err := rows.Scan(
			&url.ID,
			&url.Address,
			&url.PageTitle,
			&url.HTMLVersion,
			&url.Headings,
			&url.HasLoginForm,
			&url.CrawlStatus,
			&url.CreatedAt,
		); err != nil {
			return nil, err
		}
		urls = append(urls, url)
	}

	return urls, nil
}

func GetURLById(db *sql.DB, id int) (*URL, error) {
	var url URL

	err := db.QueryRow(`
		SELECT id, address, page_title, html_version, headings, has_login_form, crawl_status, created_at
		FROM urls WHERE id = ?
	`, id).
		Scan(
			&url.ID,
			&url.Address,
			&url.PageTitle,
			&url.HTMLVersion,
			&url.Headings,
			&url.HasLoginForm,
			&url.CrawlStatus,
			&url.CreatedAt,
		)

	if err != nil {
		return nil, err
	}
	return &url, nil
}

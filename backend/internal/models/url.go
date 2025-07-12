package models

import (
	"database/sql"
	"time"
)

type URL struct {
	ID 					int		 		`json:"id"`
	URL 				string		`json:"url"`
	CreatedAt		time.Time	`json:"created_at"`
}

//Returns all urls from the db
func GetAllURLs(db *sql.DB) ([]URL, error) {
	rows, err := db.Query("SELECT id, url, created_at FROM urls")
	if err != nil {
		return nil, err
	}

	// delays execution of Close until this function exits - cleans up resource
	defer rows.Close()

	var urls []URL

	for rows.Next() {
		var url URL
		if err := rows.Scan(&url.ID, &url.URL, &url.CreatedAt); err!= nil {
			return nil, err
		}
		urls = append(urls, url)
	}

	return urls, nil
}

func GetURLById(db *sql.DB, id int) (*URL, error) {
	var url URL

	err := db.QueryRow("SELECT id, url, created_at FROM urls WHERE id = ?", id).
		Scan(&url.ID, &url.URL, &url.CreatedAt)

	if err != nil {
		return nil, err
	}
	return &url, nil
}

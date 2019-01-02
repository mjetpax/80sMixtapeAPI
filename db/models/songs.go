package models

import (
	"log"

	"github.com/mjetpax/80sMixtapeAPI/config"
)

// Song is a struct for housing song data.
type Song struct {
	ID            int    `db:"id" json:"id"`
	Title         string `db:"title" json:"title"`
	Artist        string `db:"artist" json:"artist"`
	Year          int    `db:"year" json:"year"`
	Video         string `db:"video" json:"video"`
	DurationLabel string `db:"duration_label" json:"play_time"`
	Duration      int    `db:"duration" json:"-"`
}

// QueryParams struct for use with select queries
type QueryParams struct {
	LastValue int `db:"last_value"`
	Limit     int `db:"limit"`
}

// FetchSongs retrieves all songs from the database.
func FetchSongs(lastValue int, limit int) *[]Song {
	songs := []Song{}
	db := config.Env.DB

	// Paginate with the seek method
	queryParams := QueryParams{
		LastValue: lastValue,
		Limit:     limit,
	}

	query := `SELECT id, title, artist, year,
	video, duration_label, duration
	FROM songs
	WHERE id > :last_value
	ORDER BY id ASC
	Limit :limit`

	var song Song
	rows, err := db.NamedQuery(query, queryParams)
	if err != nil {
		log.Fatalln(err)
	}
	for rows.Next() {
		err := rows.StructScan(&song)
		if err != nil {
			log.Fatalln(err)
		}

		songs = append(songs, song)
	}

	return &songs
}

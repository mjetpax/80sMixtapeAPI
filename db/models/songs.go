package models

import (
	"github.com/mjetpax/80sMixtapeAPI/config"
)

// Song is a struct for housing song data.
type Song struct {
	ID            int    `db:"id"`
	Title         string `db:"title"`
	Artist        string `db:"artist"`
	Year          int    `db:"year"`
	Video         string `db:"video"`
	DurationLabel string `db:"duration_label"`
	Duration      int    `db:"duration"`
}

// FetchSongs retrieves all songs from the database.
func FetchSongs() []Song {
	songs := []Song{}
	db := config.Env.DB
	db.Select(&songs, "select id, title, artist, year, video, duration_label, duration from songs")

	return songs
}

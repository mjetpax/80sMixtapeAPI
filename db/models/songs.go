package models

import (
	"database/sql"

	// pq is needed to communicate to postgresql.
	_ "github.com/lib/pq"
	"github.com/mjetpax/80sMixtapeAPI/config"
)

// Song is a struct for housing song data.
type Song struct {
	Artist        string
	Title         string
	Year          int
	Video         string
	DurationLabel string
	Duration      int
}

// GetSongs retrieves songs from the database.
func GetSongs() {
	db, err := sql.Open("postgres", config.Env.DatabaseConn)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}
}

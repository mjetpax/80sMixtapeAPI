package models

import (
	_ "github.com/lib/pq"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

// NewDatabase creates a new database connection and returns it.
func NewDatabase(dbConnString string) (*sqlx.DB, error) {
	db, err := sqlx.Connect("postgres", dbConnString)
	if err != nil {
		return nil, err
	}

	// Ping the db to test the connection.
	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}

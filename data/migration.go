package data

import (
	"database/sql"
	"log"

	// pq is needed to communicate to postgresql.
	_ "github.com/lib/pq"
	"github.com/mjetpax/80sMixtapeAPI/config"
	"github.com/rubenv/sql-migrate"
)

// MigrateDB runs the latest migrations against the database.
func MigrateDB() {

	migrations := &migrate.FileMigrationSource{
		Dir: "data/migrations",
	}

	// migrate needs database/sql to operate
	db, err := sql.Open("postgres", config.Env.DBConn)
	if err != nil {
		log.Panic(err)
	}
	defer db.Close()

	n, err := migrate.Exec(db, "postgres", migrations, migrate.Up)
	if err != nil {
		log.Printf("Error during database migration: %d \n", err)

		panic(err)
	}

	log.Printf("Applied %d migrations!\n", n)
}

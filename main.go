package main

import (
	"log"
	"net/http"

	"github.com/jmoiron/sqlx"
	"github.com/julienschmidt/httprouter"
	_ "github.com/lib/pq"
	"github.com/mjetpax/80sMixtapeAPI/api"
	"github.com/mjetpax/80sMixtapeAPI/config"
	"github.com/mjetpax/80sMixtapeAPI/data"
)

const port = ":8080"

var services config.Services

func init() {
	// Initialize environmental settings, including database connection.
	config.InitEnv()

	// Make database migrations as needed.
	data.MigrateDB()

	// Set default db connection.
	db, err := sqlx.Connect("postgres", config.Env.DBConn)
	if err != nil {
		log.Panic(err)
	}

	services = config.Services{DB: db}

	// Ready to go!
	log.Println("80's Mixtape API is listening on port " + port[1:] + "!")
}

func main() {
	// Close database connection when app exits.
	defer services.DB.Close()

	// Set up some routes
	router := httprouter.New()
	router.GET("/health", api.GetHealth)
	router.GET("/songs/:last_value/:limit", api.GetSongs(&services))
	router.GET("/cassette", api.GetCassette(&services))
	router.GET("/cassette/*cassette_type", api.GetCassette(&services))

	log.Fatal(http.ListenAndServe(port, router))
}

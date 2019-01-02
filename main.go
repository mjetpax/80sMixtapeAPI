package main

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/mjetpax/80sMixtapeAPI/api"
	"github.com/mjetpax/80sMixtapeAPI/config"
	"github.com/mjetpax/80sMixtapeAPI/db"
)

const port = ":8080"

func init() {
	// Initialize environmental settings, including database connection.
	config.InitEnv()

	// Make database migrations as needed.
	db.MigrateDB()

	// Ready to go!
	log.Println("80's Mixtape API is listening on port " + port[1:] + "!")
}

func main() {
	// Close database connection when app exits.
	defer config.Env.DB.Close()

	// Set up some routes
	router := httprouter.New()
	router.GET("/health", api.GetHealth)
	router.GET("/songs/:last_value/:limit", api.GetSongs)

	log.Fatal(http.ListenAndServe(port, router))
}

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
	config.LoadConfig()
	db.MigrateDB()
	log.Println("80's Mixtape API is listening on port " + port[1:] + "!")
}

func main() {

	// Set up some routes
	router := httprouter.New()
	router.GET("/health", api.GetHealth)

	log.Fatal(http.ListenAndServe(port, router))
}

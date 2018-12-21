package main

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/mjetpax/80sMixtapeAPI/api"
)

const port = ":8080"

func init() {
	log.Println("80's Mixtape API is listening to some rad jamz on port " + port + "!")
}

func main() {

	// Set up some routes
	router := httprouter.New()
	router.GET("/health", api.GetHealth)

	log.Fatal(http.ListenAndServe(port, router))
}

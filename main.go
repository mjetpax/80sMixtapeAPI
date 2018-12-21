package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

const port = ":8080"

func init() {
	log.Println("80's Mixtape API is listening to some rad jamz!")
}

// GetHealth returns the app's health.
func GetHealth(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "80's Mixtape API is running smooth!\n")
}

func main() {

	// Set up some routes
	router := httprouter.New()
	router.GET("/health", GetHealth)

	log.Fatal(http.ListenAndServe(port, router))
}

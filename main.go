package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/mjetpax/80sMixtapeAPI/api"
	"github.com/mjetpax/80sMixtapeAPI/config"
	"github.com/mjetpax/80sMixtapeAPI/models"
)

const port = ":8080"

func init() {
	config.LoadConfig()
	log.Println("80's Mixtape API is listening on port " + port[1:] + "!")
}

// Test endpoint
func DBConnTest(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	models.GetSongs()

	fmt.Fprintf(w, "hello, %s!\n", "Chad")

}

func main() {

	// Set up some routes
	router := httprouter.New()
	router.GET("/health", api.GetHealth)

	router.GET("/db-test", DBConnTest)

	log.Fatal(http.ListenAndServe(port, router))
}

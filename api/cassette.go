package api

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/mjetpax/80sMixtapeAPI/db/models"
)

// GetCassette gets mixed songs from model and returns them.
func GetCassette(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	cassetteType := ps.ByName("cassette_type")
	cassette := models.FetchCassette(cassetteType)

	err := json.NewEncoder(w).Encode(cassette)
	if err != nil {
		log.Println(err.Error())
	}

	return
}

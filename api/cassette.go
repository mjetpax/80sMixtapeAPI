package api

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/mjetpax/80sMixtapeAPI/config"
	"github.com/mjetpax/80sMixtapeAPI/data/models"
)

// GetCassette gets mixed songs from model and returns them.
func GetCassette(services *config.Services) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

		cassetteType := ps.ByName("cassette_type")
		cassette := models.FetchCassette(services.DB, cassetteType)

		err := json.NewEncoder(w).Encode(cassette)
		if err != nil {
			log.Println(err.Error())
		}

		return
	}
}

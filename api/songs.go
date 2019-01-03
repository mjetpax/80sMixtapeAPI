package api

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"github.com/mjetpax/80sMixtapeAPI/config"
	"github.com/mjetpax/80sMixtapeAPI/data/models"
)

// GetSongs gets songs from model and returns them.
func GetSongs(services *config.Services) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

		lastValue, _ := strconv.Atoi(ps.ByName("last_value"))
		limit, _ := strconv.Atoi(ps.ByName("limit"))

		songs := models.FetchSongs(services.DB, lastValue, limit)

		err := json.NewEncoder(w).Encode(songs)
		if err != nil {
			log.Println(err.Error())
		}
		return
	}
}

package api

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"github.com/mjetpax/80sMixtapeAPI/db/models"
)

// GetSongs gets songs from model and returns them.
func GetSongs(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	lastValue, _ := strconv.Atoi(ps.ByName("last_value"))
	limit, _ := strconv.Atoi(ps.ByName("limit"))

	songs := models.FetchSongs(lastValue, limit)

	err := json.NewEncoder(w).Encode(songs)
	if err != nil {
		log.Println(err.Error())
	}
	return
}

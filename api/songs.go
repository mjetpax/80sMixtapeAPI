package api

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/mjetpax/80sMixtapeAPI/db/models"
)

// GetSongs gets songs from model and returns them.
func GetSongs(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	songs := models.FetchSongs()

	err := json.NewEncoder(w).Encode(songs)
	if err != nil {
		log.Println(err.Error())
	}
	return
}

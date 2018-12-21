package api

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
)

var startTime = time.Now()

// Health check struct stores values for output to health check requests.
type Health struct {
	AppName   string `json:"application_name"`
	Message   string `json:"message"`
	StartDate string `json:"start_date"`
	UpTime    string `json:"up_time"`
}

// GetHealth returns the app's health.
func GetHealth(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	uptime := time.Since(startTime)
	health := Health{"80's Mixtape API", "80's Mixtape API is running smooth!", startTime.String(), uptime.String()}
	err := json.NewEncoder(w).Encode(health)
	if err != nil {
		log.Println(err.Error())
	}
	return
}

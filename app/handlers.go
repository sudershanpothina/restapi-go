package app

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func getTime(w http.ResponseWriter, r *http.Request){
	wtime := new(WorldTime)
	wtime.CurrentTime = time.Now()
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(wtime)	
}
func getTimeByLocation(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	wtime := new(WorldTime)
	if vars["tz"] == "Asia" {
		wtime.Asia = getLocalTime("Asia/Shanghai")
	}
	if vars["tz"] == "America" {
		wtime.America = getLocalTime("America/New_York")
	}
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(wtime)

}

func getLocalTime(location string) time.Time {
	loc, _ := time.LoadLocation(location)
	return time.Now().In(loc)
}

type WorldTime struct {
	CurrentTime time.Time `json:"current_time"`
	Asia time.Time `json:"Asia/Kolkata"`
	America time.Time `json:"America/New_York"`

}
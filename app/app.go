package app

import(
	"net/http"
	"log"
	"github.com/gorilla/mux"
)

func Start() {
	router := mux.NewRouter()
	router.HandleFunc("/api/time", getTime).Methods(http.MethodGet)
	router.HandleFunc("/api/time/{tz}", getTimeByLocation).Methods(http.MethodGet)
	log.Fatal(http.ListenAndServe(":8080", router))
}
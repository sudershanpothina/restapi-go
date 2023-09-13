package app

import(
	"net/http"
	"log"
)

func Start() {
	mux := http.NewServeMux()
	mux.HandleFunc("/hello", hello)
	mux.HandleFunc("/customers", getAllCustomers)
	log.Fatal(http.ListenAndServe(":8080", mux))
}
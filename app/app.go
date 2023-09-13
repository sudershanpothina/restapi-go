package app

import(
	"net/http"
	"log"
)

func Start() {
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/customers", getAllCustomers)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
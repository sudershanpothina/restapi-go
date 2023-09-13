package app

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func hello(w http.ResponseWriter, r *http.Request){
	fmt.Fprint(w, "Hello")
}
func getAllCustomers(w http.ResponseWriter, r *http.Request){
	customers := []Customer {
		{Name: "Subject1", City: "City1", ZipCode: "400040"},
		{Name: "Subject2", City: "City2", ZipCode: "500032"},
	}
	if r.Header.Get("Content-Type") == "application/xml" {
		w.Header().Add("Content-Type", "application/xml")
		xml.NewEncoder(w).Encode(customers)
	}else {
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(customers)
	}
	
}
func getCustomer(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	fmt.Fprint(w, vars["customer_id"])

}
func createCustomer(w http.ResponseWriter, r *http.Request){
	fmt.Fprint(w, "Post request received")
}
type Customer struct {
	Name string `json:"name" xml:"name"`
	City string `json:"city" xml:"city"`
	ZipCode string `json:"zip" xml:"zip"`
}
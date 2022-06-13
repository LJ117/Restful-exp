package app

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type Customer struct {
	Name    string `json:"full_name" xml:"xml_name"`
	City    string `json:"city" xml:"xml_city"`
	ZipCode string `json:"zip_code" xml:"xml_zip_code"`
}

func greet(w http.ResponseWriter, req *http.Request) {
	fprint, err := fmt.Fprint(w, "Hello World !!!")
	if err != nil {
		log.Printf("total %d bytes\n", fprint)
		return
	}
}
func getCustomer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Fprint(w, vars["customer_id"])
}
func getAllCustomers(w http.ResponseWriter, r *http.Request) {
	customers := []Customer{
		{"Seven", "CDC", "610061"},
		{"Bill", "GD", "710071"},
	}

	if r.Header.Get("Content-Type") == "application/xml" {
		w.Header().Set("Content-Type", "application/xml")
		if err := xml.NewEncoder(w).Encode(customers); err != nil {
			log.Fatal(err)
		}
	} else {
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(customers); err != nil {
			log.Fatal(err)
		}
	}
}

func createCustomer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "POST request")
}

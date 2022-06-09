package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Customer struct {
	Name    string `json:"full_name"`
	City    string `json:"city"`
	ZipCode string `json:"zip_code"`
}

func main() {

	http.HandleFunc("/greet", greet)
	http.HandleFunc("/customers", getAllCustomers)

	log.Fatal(http.ListenAndServe("localhost:8888", nil))
}

func greet(w http.ResponseWriter, req *http.Request) {
	fprint, err := fmt.Fprint(w, "Hello World !!!")
	if err != nil {
		log.Printf("total %d bytes\n", fprint)
		return
	}
}

func getAllCustomers(w http.ResponseWriter, r *http.Request) {
	customers := []Customer{
		{"Seven", "CDC", "610061"},
		{"Bill", "GD", "710071"},
	}
	w.Header().Set("Content-Type", "Application/json")
	if err := json.NewEncoder(w).Encode(customers); err != nil {
		log.Fatal(err)
	}

}

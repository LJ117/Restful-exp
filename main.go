package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Customer struct {
	Name    string
	City    string
	ZipCode string
}

func main() {

	http.HandleFunc("/greet", greet)
	http.HandleFunc("/customers", getAllCustomers)

	http.ListenAndServe("localhost:8888", nil)
}

func greet(w http.ResponseWriter, req *http.Request) {
	fmt.Fprint(w, "Hello World !!!")
}

func getAllCustomers(w http.ResponseWriter, r *http.Request) {
	customers := []Customer{
		{"Seven", "CDC", "610061"},
		{"Reaka", "GD", "710071"},
	}
	if err := json.NewEncoder(w).Encode(customers); err != nil {
		log.Fatal(err)
	}

}

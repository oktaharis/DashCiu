package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	monitoringUser "github.com/jeypc/homecontroller/monitoringuser"
)

func main() {
	r := mux.NewRouter()

	// Menggunakan variabel app dalam URL dan membiarkan page sebagai query parameter
	r.HandleFunc("/monitoring", monitoringUser.IndexMonuser).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", r))
}

package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jeypc/homecontroller/models"
	monitoringUser "github.com/jeypc/homecontroller/monitoringuser"
)

func main() {
	r := mux.NewRouter()
	models.ConnectDatabase("user")

	// Menggunakan variabel app dalam URL dan membiarkan page sebagai query parameter
	r.HandleFunc("/monitoring", monitoringUser.IndexMonuser).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", r))
}

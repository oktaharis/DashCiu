package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jeypc/homecontroller/controller/dashboardflexi"
)

func main() {
	r := mux.NewRouter()

	// Menggunakan variabel app dalam URL dan membiarkan page sebagai query parameter

	r.HandleFunc("/dashboard", dashboardflexi.IndexDashFlexi).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", r))
}

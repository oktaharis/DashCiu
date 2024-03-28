package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jeypc/homecontroller/controller/dashboardhandler"
)

func main() {
	r := mux.NewRouter()

	// Menggunakan variabel app dalam URL dan membiarkan page sebagai query parameter
	r.HandleFunc("/dashboard", dashboardhandler.IndexHandler).Methods("POST")

	log.Fatal(http.ListenAndServe(":8080", r))
}

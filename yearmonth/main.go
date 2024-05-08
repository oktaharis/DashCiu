package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	yearmonth "github.com/jeypc/homecontroller/controller"
)

func main() {
	r := mux.NewRouter()

	// Menggunakan variabel app dalam URL dan membiarkan page sebagai query parameter
	
	r.HandleFunc("/dashboard", yearmonth.IndexYear).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", r))
}

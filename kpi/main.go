package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jeypc/homecontroller/controller/explorespl"
	"github.com/jeypc/homecontroller/controller/filesredinesspl"
	"github.com/jeypc/homecontroller/controller/manageclaimspl"
)

func main() {
	r := mux.NewRouter()
	

	// Menggunakan variabel app dalam URL dan membiarkan page sebagai query parameter

	r.HandleFunc("/explorespl", explore.ExploreSpl).Methods("POST")
	r.HandleFunc("/filesredinesspl", filesredines.FilesSpl).Methods("POST")
	r.HandleFunc("/manageclaimspl", manageclaim.ClaimSpl).Methods("POST")

	log.Fatal(http.ListenAndServe(":8080", r))
}

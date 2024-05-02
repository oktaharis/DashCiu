package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	dashboardhandler "github.com/jeypc/homecontroller/controller/dashboardspl"
	"github.com/jeypc/homecontroller/controller/explorespl"
	"github.com/jeypc/homecontroller/controller/filesredinesspl"
	"github.com/jeypc/homecontroller/controller/manageclaimspl"
	managepolicyspl "github.com/jeypc/homecontroller/controller/managepolicy"
)

func main() {
	r := mux.NewRouter()

	// Menggunakan variabel app dalam URL dan membiarkan page sebagai query parameter

	r.HandleFunc("/explorespl", explore.ExploreSpl).Methods("POST")
	r.HandleFunc("/filesredinesspl", filesredines.FilesSpl).Methods("POST")
	r.HandleFunc("/manageclaimspl", manageclaim.ClaimSpl).Methods("POST")
	r.HandleFunc("/managepolicyspl", managepolicyspl.PolicySpl).Methods("POST")
	r.HandleFunc("/dashboardspl", dashboardhandler.IndexDashSpl).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", r))
}

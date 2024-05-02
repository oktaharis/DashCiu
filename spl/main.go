package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jeypc/homecontroller/controller/dashboardhandlerspl"
	"github.com/jeypc/homecontroller/controller/explorespl"
	"github.com/jeypc/homecontroller/controller/filesredinesspl"
	"github.com/jeypc/homecontroller/controller/manageclaimspl"
	"github.com/jeypc/homecontroller/controller/managepolicyspl"
	"github.com/jeypc/homecontroller/controller/subrogationspl"
)

func main() {
	r := mux.NewRouter()

	// Menggunakan variabel app dalam URL dan membiarkan page sebagai query parameter
	
	r.HandleFunc("/dashboardspl", dashboardhandler.DashboardSpl).Methods("POST")
	r.HandleFunc("/explorespl", explore.ExploreSpl).Methods("POST")
	r.HandleFunc("/filesredinesspl", filesredines.FilesSpl).Methods("POST")
	r.HandleFunc("/claimspl", manageclaim.ClaimSpl).Methods("POST")
	r.HandleFunc("/policyspl", managepolicy.PolicySpl).Methods("POST")
	r.HandleFunc("/subrogationspl", subrogation.SubrogationSpl).Methods("POST")

	log.Fatal(http.ListenAndServe(":8080", r))
}

package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	claimperiodafi "github.com/jeypc/homecontroller/controller/afi/claim"
	dashboardperiodafi "github.com/jeypc/homecontroller/controller/afi/dashboard"
	policyperiodafi "github.com/jeypc/homecontroller/controller/afi/policy"
)

func main() {
	r := mux.NewRouter()

	// Menggunakan variabel app dalam URL dan membiarkan page sebagai query parameter
	
	r.HandleFunc("/manageclaim", claimperiodafi.ClaimPeriodAfi).Methods("GET")
	r.HandleFunc("/dashboard", dashboardperiodafi.DashboardPeriodAfi).Methods("GET")
	r.HandleFunc("/managepolicy", policyperiodafi.PolicyPeriodAfi).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", r))
}

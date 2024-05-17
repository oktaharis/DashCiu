package main

import (
	"log"
	"net/http"

	"aficontroller/controller/dashboardafi"
	 "aficontroller/controller/filereadiness"
	 "aficontroller/controller/manageclaim"
	 "aficontroller/controller/managepolicy"
	"aficontroller/models"
	"aficontroller/controller/period/claim"
	"aficontroller/controller/period/dashboard"
	"aficontroller/controller/period/policy"


	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	// Menggunakan variabel app dalam URL dan membiarkan page sebagai query parameter
	models.ConnectDatabase("afi")
	r.HandleFunc("/dashboardafi", dashboardafi.IndexDashAfi).Methods("GET")
	r.HandleFunc("/managepolicyafi", managepolicyafi.IndexPolicyAfi).Methods("GET")
	r.HandleFunc("/manageclaimafi", manageclaimafi.IndexClaim).Methods("GET")
	r.HandleFunc("/fileredinesafi", filereadinessafi.IndexFilesAfi).Methods("GET")

	// ini period
	r.HandleFunc("/claimafi/period", claimperiodafi.ClaimPeriodAfi).Methods("GET")
	r.HandleFunc("/dashboardafi/period", dashboardperiodafi.DashboardPeriodAfi).Methods("GET")
	r.HandleFunc("/managepolicyafi/period", policyperiodafi.PolicyPeriodAfi).Methods("GET")


	log.Fatal(http.ListenAndServe(":8080", r))
}

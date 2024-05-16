package main

import (
	"log"
	"net/http"

	"kpicontroller/controller/dashboardkpi"
	filereadinesskpi "kpicontroller/controller/filereadiness"
	manageclaimkpi "kpicontroller/controller/manageclaim"
	managepolicykpi "kpicontroller/controller/managepolicy"
	claimperiodkpi "kpicontroller/controller/period/claim"
	dashboardperiodkpi "kpicontroller/controller/period/dashboard"
	policyperiodkpi "kpicontroller/controller/period/policy"
	"kpicontroller/models"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	// Menggunakan variabel app dalam URL dan membiarkan page sebagai query parameter
	models.ConnectDatabase()
	r.HandleFunc("/dashboard", dashboardkpi.IndexDashKpi).Methods("GET")
	r.HandleFunc("/managepolicy", managepolicykpi.IndexPolicyKpi).Methods("GET")
	r.HandleFunc("/manageclaim", manageclaimkpi.IndexClaimKpi).Methods("GET")
	r.HandleFunc("/fileredines", filereadinesskpi.IndexFilesAfi).Methods("GET")

	// ini period
	r.HandleFunc("/claim/period", claimperiodkpi.ClaimPeriodKpi).Methods("GET")
	r.HandleFunc("/dashboard/period", dashboardperiodkpi.DashboardPeriodKpi).Methods("GET")
	r.HandleFunc("/managepolicy/period", policyperiodkpi.PolicyPeriodKpi).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", r))
}

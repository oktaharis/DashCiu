package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	claimperiodkpi "github.com/jeypc/homecontroller/controller/kpi/claim"
	dashboardperiodkpi "github.com/jeypc/homecontroller/controller/kpi/dashboard"
	policyperiodkpi "github.com/jeypc/homecontroller/controller/kpi/policy"
	subrogationperiodkpi "github.com/jeypc/homecontroller/controller/kpi/subrogation"
)

func main() {
	r := mux.NewRouter()

	// Menggunakan variabel app dalam URL dan membiarkan page sebagai query parameter
	
	r.HandleFunc("/dashboard", dashboardperiodkpi.DashboardPeriodKpi).Methods("GET")
	r.HandleFunc("/managepolicy", policyperiodkpi.PolicyPeriodKpi).Methods("GET")
	r.HandleFunc("/manageclaim", claimperiodkpi.ClaimPeriodKpi).Methods("GET")
	r.HandleFunc("/subgrogation", subrogationperiodkpi.SubroPeriodKpi).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", r))
}

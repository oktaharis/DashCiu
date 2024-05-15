package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	claimperiodflexi "github.com/jeypc/homecontroller/controller/flexi/claim"
	dashboardperiodflexi "github.com/jeypc/homecontroller/controller/flexi/dashboard"
	policyperiodflexi "github.com/jeypc/homecontroller/controller/flexi/policy"
	subrogationperiodflexi "github.com/jeypc/homecontroller/controller/flexi/subrogation"
	uploadperiodflexi "github.com/jeypc/homecontroller/controller/flexi/upload"
)

func main() {
	r := mux.NewRouter()

	// Menggunakan variabel app dalam URL dan membiarkan page sebagai query parameter
	
	r.HandleFunc("/dashboard", dashboardperiodflexi.DashboardPeriodFlexi).Methods("GET")
	r.HandleFunc("/manageclaim", claimperiodflexi.ClaimPeriodFlexi).Methods("GET")
	r.HandleFunc("/managepolicy", policyperiodflexi.PolicyPeriodFlexi).Methods("GET")
	r.HandleFunc("/subgrogation", subrogationperiodflexi.SubroPeriodFlexi).Methods("GET")
	r.HandleFunc("/file-upload", uploadperiodflexi.UploadPeriodFlexi).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", r))
}

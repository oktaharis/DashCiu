package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	claimperiodspl "github.com/jeypc/homecontroller/controller/spl/claim"
	dashboardperiodspl "github.com/jeypc/homecontroller/controller/spl/dashboard"
	filereadinesperiodspl "github.com/jeypc/homecontroller/controller/spl/filereadines"
	policyperiodspl "github.com/jeypc/homecontroller/controller/spl/policy"
	subrogationperiodspl "github.com/jeypc/homecontroller/controller/spl/subrogation"
	uploadperiodspl "github.com/jeypc/homecontroller/controller/spl/upload"
)

func main() {
	r := mux.NewRouter()

	// Menggunakan variabel app dalam URL dan membiarkan page sebagai query parameter
	
	r.HandleFunc("/dashboard", dashboardperiodspl.DashboardPeriodSpl).Methods("GET")
	r.HandleFunc("/managepolicy", policyperiodspl.PolicyPeriodSpl).Methods("GET")
	r.HandleFunc("/manageclaim", claimperiodspl.ClaimPeriodSpl).Methods("GET")
	r.HandleFunc("/subgrogation", subrogationperiodspl.SubroPeriodSpl).Methods("GET")
	r.HandleFunc("/files-rediness", filereadinesperiodspl.FilePeriodSpl).Methods("GET")
	r.HandleFunc("/file-upload", uploadperiodspl.UploadPeriodSpl).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", r))
}

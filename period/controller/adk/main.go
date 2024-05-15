package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	claimperiodadk "github.com/jeypc/homecontroller/controller/adk/claim"
	dashboardperiodadk "github.com/jeypc/homecontroller/controller/adk/dashboard"
	filereadinesperiodadk "github.com/jeypc/homecontroller/controller/adk/filereadines"
	policyperiodadk "github.com/jeypc/homecontroller/controller/adk/policy"
	subrogationperiodadk "github.com/jeypc/homecontroller/controller/adk/subrogation"
	uploadperiodadk "github.com/jeypc/homecontroller/controller/adk/upload"
)

func main() {
	r := mux.NewRouter()

	// Menggunakan variabel app dalam URL dan membiarkan page sebagai query parameter
	
	r.HandleFunc("/dashboard", dashboardperiodadk.DashboardPeriodAdk).Methods("GET")
	r.HandleFunc("/managepolicy", policyperiodadk.PolicyPeriodAdk).Methods("GET")
	r.HandleFunc("/manageclaim", claimperiodadk.ClaimPeriodAdk).Methods("GET")
	r.HandleFunc("/subgrogation", subrogationperiodadk.SubroPeriodAdk).Methods("GET")
	r.HandleFunc("/files-redines", filereadinesperiodadk.FilePeriodAdk).Methods("GET")
	r.HandleFunc("/file-upload", uploadperiodadk.UploadPeriodadk).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", r))
}

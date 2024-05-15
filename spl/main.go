package main

import (
	"log"
	"net/http"

	dashboardhandler "splcontroller/controller/dashboardspl"
	"splcontroller/controller/explorespl"
	"splcontroller/controller/filesredinesspl"
	"splcontroller/controller/manageclaimspl"
	managepolicyspl "splcontroller/controller/managepolicy"
	"splcontroller/controller/subrogationspl"
	"splcontroller/controller/uploadspl"
	"splcontroller/controller/userspl"
	"splcontroller/models"
	"splcontroller/controller/period/dashboard"
	"splcontroller/controller/period/claim"
	"splcontroller/controller/period/filereadines"
	"splcontroller/controller/period/policy"
	"splcontroller/controller/period/subrogation"
	"splcontroller/controller/period/upload"


	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	models.ConnectDatabase("spl")

	// Menggunakan variabel app dalam URL dan membiarkan page sebagai query parameter
	
	r.HandleFunc("/explorespl", explore.ExploreSpl).Methods("GET")
	r.HandleFunc("/filesredinesspl", filesredines.FilesSpl).Methods("GET")
	r.HandleFunc("/manageclaimspl", manageclaim.ClaimSpl).Methods("GET")
	r.HandleFunc("/managepolicyspl", managepolicyspl.PolicySpl).Methods("GET")
	r.HandleFunc("/dashboardspl", dashboardhandler.IndexDashSpl).Methods("GET")
	r.HandleFunc("/subrogationspl", subrogation.SubrogationSpl).Methods("GET")
	r.HandleFunc("/uploadspl", upload.UploadSpl).Methods("GET")
	r.HandleFunc("/userspl", user.UserSpl).Methods("GET")

	// ini period

	r.HandleFunc("/dashboard/period", dashboardperiodspl.DashboardPeriodSpl).Methods("GET")
	r.HandleFunc("/managepolicy/period", policyperiodspl.PolicyPeriodSpl).Methods("GET")
	r.HandleFunc("/manageclaim/period", claimperiodspl.ClaimPeriodSpl).Methods("GET")
	r.HandleFunc("/subgrogation/period", subrogationperiodspl.SubroPeriodSpl).Methods("GET")
	r.HandleFunc("/files-rediness/period", filereadinesperiodspl.FilePeriodSpl).Methods("GET")
	r.HandleFunc("/file-upload/period", uploadperiodspl.UploadPeriodSpl).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", r))
}

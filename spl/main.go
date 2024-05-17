package main

import (
	"log"
	"net/http"

	dashboardhandler "splcontroller/controller/dashboardspl"
	"splcontroller/controller/explorespl"
	"splcontroller/controller/filesredinesspl"
	"splcontroller/controller/manageclaimspl"
	managepolicyspl "splcontroller/controller/managepolicy"
	"splcontroller/controller/period/claim"
	"splcontroller/controller/period/dashboard"
	"splcontroller/controller/period/filereadines"
	"splcontroller/controller/period/policy"
	"splcontroller/controller/period/subrogation"
	"splcontroller/controller/period/upload"
	"splcontroller/controller/subrogationspl"
	"splcontroller/controller/uploadspl"
	"splcontroller/models"

	"github.com/gorilla/mux"
) 

func main() {
    r := mux.NewRouter()
	// Koneksi ke database saat aplikasi dimulai
	models.ConnectDatabase()
	// Menggunakan variabel app dalam URL dan membiarkan page sebagai query parameter
	
	r.HandleFunc("/explorespl", explore.ExploreSpl).Methods("GET")
	r.HandleFunc("/filesredinesspl", filesredines.FilesSpl).Methods("GET")
	r.HandleFunc("/manageclaimspl", manageclaim.ClaimSpl).Methods("GET")
	r.HandleFunc("/managepolicyspl", managepolicyspl.PolicySpl).Methods("GET")
	r.HandleFunc("/dashboardspl", dashboardhandler.IndexDashSpl).Methods("GET")
	r.HandleFunc("/subrogationspl", subrogation.SubrogationSpl).Methods("GET")
	r.HandleFunc("/uploadspl", upload.UploadSpl).Methods("GET")

	// ini period

	r.HandleFunc("/dashboardspl/period", dashboardperiodspl.DashboardPeriodSpl).Methods("GET")
	r.HandleFunc("/managepolicyspl/period", policyperiodspl.PolicyPeriodSpl).Methods("GET")
	r.HandleFunc("/manageclaimspl/period", claimperiodspl.ClaimPeriodSpl).Methods("GET")
	r.HandleFunc("/subgrogationspl/period", subrogationperiodspl.SubroPeriodSpl).Methods("GET")
	r.HandleFunc("/files-redinessspl/period", filereadinesperiodspl.FilePeriodSpl).Methods("GET")
	r.HandleFunc("/file-uploadspl/period", uploadperiodspl.UploadPeriodSpl).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", r))
}

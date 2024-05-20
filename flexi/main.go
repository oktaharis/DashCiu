package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"flexicontroller/controller/dashboardflexi"
	"flexicontroller/controller/filesredinesflexi"
	"flexicontroller/controller/manageclaimflexi"
	"flexicontroller/controller/managepolicyflexi"
	"flexicontroller/controller/subrogationflexi"
	"flexicontroller/controller/uploadflexi"
	"flexicontroller/models"

	"flexicontroller/controller/period/dashboard"
	"flexicontroller/controller/period/claim"
	"flexicontroller/controller/period/policy"
	"flexicontroller/controller/period/subrogation"
	"flexicontroller/controller/period/upload"

)

func main() {
	r := mux.NewRouter()
	models.ConnectDatabase()

	// Menggunakan variabel app dalam URL dan membiarkan page sebagai query parameter
	r.HandleFunc("/filesredinesflexi", filesredines.FilesFlexi).Methods("GET")
	r.HandleFunc("/manageclaimflexi", manageclaim.ClaimFlexi).Methods("GET")
	r.HandleFunc("/managepolicyflexi", managepolicy.PolicyFlexi).Methods("GET")
	r.HandleFunc("/uploadflexi", upload.UploadFlexi).Methods("GET")
	r.HandleFunc("/dashboardflexi", dashboardflexi.IndexDashFlexi).Methods("GET")
	r.HandleFunc("/subrogationflexi", subrogation.SubrogationFlexi).Methods("GET")

	// ini adalah period
	
	r.HandleFunc("/dashboard/period", dashboardperiodflexi.DashboardPeriodFlexi).Methods("GET")
	r.HandleFunc("/manageclaim/period", claimperiodflexi.ClaimPeriodFlexi).Methods("GET")
	r.HandleFunc("/managepolicy/period", policyperiodflexi.PolicyPeriodFlexi).Methods("GET")
	r.HandleFunc("/subgrogation/period", subrogationperiodflexi.SubroPeriodFlexi).Methods("GET")
	r.HandleFunc("/file-upload/period", uploadperiodflexi.UploadPeriodFlexi).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", r))
}

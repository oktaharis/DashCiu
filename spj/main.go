package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	dashboardhandler "splcontroller/controller/dashboardspj"
	"splcontroller/controller/explorespj"
	"splcontroller/controller/filesredinesspj"
	"splcontroller/controller/manageclaimspj"
	managepolicyspj "splcontroller/controller/managepolicyspj"
	"splcontroller/controller/subrogationspj"
	"splcontroller/controller/uploadspj"
	"splcontroller/controller/userspj"
	"splcontroller/models"

	"splcontroller/controller/period/dashboard"
	"splcontroller/controller/period/claim"
	"splcontroller/controller/period/filereadines"
	"splcontroller/controller/period/policy"
	"splcontroller/controller/period/subrogation"
	"splcontroller/controller/period/upload"

)

func main() {
	r := mux.NewRouter()
	models.ConnectDatabase()
	// Menggunakan variabel app dalam URL dan membiarkan page sebagai query parameter
	
	r.HandleFunc("/explorespj", explore.ExploreSpj).Methods("GET")
	r.HandleFunc("/filesredinesspj", filesredines.FilesSpj).Methods("GET")
	r.HandleFunc("/manageclaimspj", manageclaim.ClaimSpj).Methods("GET")
	r.HandleFunc("/managepolicyspj", managepolicyspj.PolicySpj).Methods("GET")
	r.HandleFunc("/dashboardspj", dashboardhandler.IndexDashSpj).Methods("GET")
	r.HandleFunc("/subrogationspj", subrogation.SubrogationSpj).Methods("GET")
	r.HandleFunc("/uploadspj", upload.UploadSpj).Methods("GET")
	r.HandleFunc("/userspj", user.UserSpj).Methods("GET")

	// ini adalah period
		
	r.HandleFunc("/dashboard", dashboardperiodspj.DashboardPeriodSpj).Methods("GET")
	r.HandleFunc("/managepolicy", policyperiodspj.PolicyPeriodSpj).Methods("GET")
	r.HandleFunc("/manageclaim", claimperiodspj.ClaimPeriodSpj).Methods("GET")
	r.HandleFunc("/subgrogation", subrogationperiodspj.SubroPeriodSpj).Methods("GET")
	r.HandleFunc("/files-redines", filereadinesperiodspj.FilePeriodSpj).Methods("GET")
	r.HandleFunc("/file-upload", uploadperiodspj.UploadPeriodSpj).Methods("GET")


	log.Fatal(http.ListenAndServe(":8080", r))
}

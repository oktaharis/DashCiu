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

	log.Fatal(http.ListenAndServe(":8080", r))
}

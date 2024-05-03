package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	dashboardhandler "github.com/jeypc/homecontroller/controller/dashboardspl"
	"github.com/jeypc/homecontroller/controller/explorespl"
	"github.com/jeypc/homecontroller/controller/filesredinesspl"
	"github.com/jeypc/homecontroller/controller/manageclaimspl"
	managepolicyspl "github.com/jeypc/homecontroller/controller/managepolicy"
	"github.com/jeypc/homecontroller/controller/subrogationspl"
	"github.com/jeypc/homecontroller/controller/uploadspl"
	"github.com/jeypc/homecontroller/controller/userspl"
)

func main() {
	r := mux.NewRouter()

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

package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"adkcontroller/models"

	"adkcontroller/controller/dashboardadk"
	"adkcontroller/controller/managepolicyadk"
	"adkcontroller/controller/manageclaimadk"
)

func main() {
	r := mux.NewRouter()
	models.ConnectDatabase("adk")

	// Menggunakan variabel app dalam URL dan membiarkan page sebagai query parameter

	r.HandleFunc("/dashboardadk", dashboardhandler.IndexDashAdk).Methods("GET")
	r.HandleFunc("/managepolicyadk", managepolicy.IndexPolicyAdk).Methods("GET")
	r.HandleFunc("/manageclaimadk", manageclaim.IndexClaimAdk).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", r))
}

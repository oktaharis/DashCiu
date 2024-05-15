package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"adkcontroller/controller/manageclaimadk"
	"adkcontroller/controller/dashboardadk"
)

func main() {
	r := mux.NewRouter()
	models.ConnectDatabase()

	// Menggunakan variabel app dalam URL dan membiarkan page sebagai query parameter

	r.HandleFunc("/explorespl", explore.ExploreSpl).Methods("POST")
	r.HandleFunc("/filesredinesspl", filesredines.FilesSpl).Methods("POST")
	r.HandleFunc("/manageclaimspl", manageclaim.ClaimSpl).Methods("POST")

	log.Fatal(http.ListenAndServe(":8080", r))
}

package main

import (
	"log"
	"net/http"

	"aficontroller/controller/dashboardafi"
	filereadinessafi "aficontroller/controller/filereadiness"
	manageclaimafi "aficontroller/controller/manageclaim"
	managepolicyafi "aficontroller/controller/managepolicy"
	"aficontroller/models"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	// Menggunakan variabel app dalam URL dan membiarkan page sebagai query parameter
	models.ConnectDatabase("afi")
	r.HandleFunc("/dashboard", dashboardafi.IndexDashAfi).Methods("GET")
	r.HandleFunc("/managepolicy", managepolicyafi.IndexPolicyAfi).Methods("GET")
	r.HandleFunc("/manageclaim", manageclaimafi.IndexClaim).Methods("GET")
	r.HandleFunc("/fileredines", filereadinessafi.IndexFilesAfi).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", r))
}

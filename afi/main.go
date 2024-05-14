package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"homecontroller/controller/dashboardafi"
	filereadinessafi "homecontroller/controller/filereadiness"
	manageclaimafi "homecontroller/controller/manageclaim"
	managepolicyafi "homecontroller/controller/managepolicy"
)

func main() {
	r := mux.NewRouter()

	// Menggunakan variabel app dalam URL dan membiarkan page sebagai query parameter

	r.HandleFunc("/dashboard", dashboardafi.IndexDashAfi).Methods("GET")
	r.HandleFunc("/managepolicy", managepolicyafi.IndexPolicyAfi).Methods("GET")
	r.HandleFunc("/manageclaim", manageclaimafi.IndexClaim).Methods("GET")
	r.HandleFunc("/fileredines", filereadinessafi.IndexFilesAfi).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", r))
}

package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jeypc/homecontroller/controller/dashboardafi"
	filereadinessafi "github.com/jeypc/homecontroller/controller/filereadiness"
	manageclaimafi "github.com/jeypc/homecontroller/controller/manageclaim"
	managepolicyafi "github.com/jeypc/homecontroller/controller/managepolicy"
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

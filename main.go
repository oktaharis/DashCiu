package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jeypc/homecontroller/controller/dashboardhandler"
	"github.com/jeypc/homecontroller/controller/managepolicy"
	"github.com/jeypc/homecontroller/controller/subrogation"
)

func main() {
	r := mux.NewRouter()

	// Menggunakan variabel app dalam URL dan membiarkan page sebagai query parameter
	r.HandleFunc("/dashboard", dashboardhandler.IndexDash).Methods("POST")
	r.HandleFunc("/managepolicy", managepolicy.IndexPolicy).Methods("POST")
	r.HandleFunc("/subrogation", subrogation.IndexSub).Methods("POST")

	log.Fatal(http.ListenAndServe(":8080", r))
}

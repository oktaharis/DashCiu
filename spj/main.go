package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	dashboardhandler "github.com/jeypc/homecontroller/controller/dashboardspj"
	"github.com/jeypc/homecontroller/controller/explorespj"
	"github.com/jeypc/homecontroller/controller/filesredinesspj"
	"github.com/jeypc/homecontroller/controller/manageclaimspj"
	managepolicyspj "github.com/jeypc/homecontroller/controller/managepolicyspj"
	"github.com/jeypc/homecontroller/controller/subrogationspj"
	"github.com/jeypc/homecontroller/controller/uploadspj"
	"github.com/jeypc/homecontroller/controller/userspj"
)

func main() {
	r := mux.NewRouter()

	// Menggunakan variabel app dalam URL dan membiarkan page sebagai query parameter
	
	r.HandleFunc("/explorespj", explore.ExploreSpj).Methods("GET")
	r.HandleFunc("/filesredinesspj", filesredines.FilesSpj).Methods("GET")
	r.HandleFunc("/manageclaimspj", manageclaim.ClaimSpj).Methods("GET")
	r.HandleFunc("/managepolicyspj", managepolicyspj.PolicySpj).Methods("GET")
	r.HandleFunc("/dashboardspj", dashboardhandler.IndexDashSpj).Methods("GET")
	r.HandleFunc("/subrogationspj", subrogation.SubrogationSpj).Methods("GET")
	r.HandleFunc("/uploadspj", upload.UploadSpj).Methods("GET")
	r.HandleFunc("/userspj", user.UserSpj).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", r))
}

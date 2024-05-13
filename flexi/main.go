package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jeypc/homecontroller/controller/dashboardflexi"
	"github.com/jeypc/homecontroller/controller/filesredinesflexi"
	"github.com/jeypc/homecontroller/controller/manageclaimflexi"
	"github.com/jeypc/homecontroller/controller/managepolicyflexi"
	"github.com/jeypc/homecontroller/controller/userflexi"
	"github.com/jeypc/homecontroller/controller/uploadflexi"
	"github.com/jeypc/homecontroller/controller/subrogationflexi"
)

func main() {
	r := mux.NewRouter()

	// Menggunakan variabel app dalam URL dan membiarkan page sebagai query parameter
	r.HandleFunc("/filesredinesflexi", filesredines.FilesFlexi).Methods("GET")
	r.HandleFunc("/manageclaimflexi", manageclaim.ClaimFlexi).Methods("GET")
	r.HandleFunc("/managepolicyflexi", managepolicy.PolicyFlexi).Methods("GET")
	r.HandleFunc("/userflexi", user.UserFlexi).Methods("GET")
	r.HandleFunc("/uploadflexi", upload.UploadFlexi).Methods("GET")
	r.HandleFunc("/dashboardflexi", dashboardflexi.IndexDashFlexi).Methods("GET")
	r.HandleFunc("/subrogationflexi", subrogation.SubrogationFlexi).Methods("GET")


	log.Fatal(http.ListenAndServe(":8080", r))
}

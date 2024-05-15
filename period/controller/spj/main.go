package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	claimperiodspj "github.com/jeypc/homecontroller/controller/spj/claim"
	dashboardperiodspj "github.com/jeypc/homecontroller/controller/spj/dashboard"
	filereadinesperiodspj "github.com/jeypc/homecontroller/controller/spj/filereadines"
	policyperiodspj "github.com/jeypc/homecontroller/controller/spj/policy"
	subrogationperiodspj "github.com/jeypc/homecontroller/controller/spj/subrogation"
	uploadperiodspj "github.com/jeypc/homecontroller/controller/spj/upload"
)

func main() {
	r := mux.NewRouter()

	// Menggunakan variabel app dalam URL dan membiarkan page sebagai query parameter
	
	r.HandleFunc("/dashboard", dashboardperiodspj.DashboardPeriodSpj).Methods("GET")
	r.HandleFunc("/managepolicy", policyperiodspj.PolicyPeriodSpj).Methods("GET")
	r.HandleFunc("/manageclaim", claimperiodspj.ClaimPeriodSpj).Methods("GET")
	r.HandleFunc("/subgrogation", subrogationperiodspj.SubroPeriodSpj).Methods("GET")
	r.HandleFunc("/files-redines", filereadinesperiodspj.FilePeriodSpj).Methods("GET")
	r.HandleFunc("/file-upload", uploadperiodspj.UploadPeriodSpj).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", r))
}

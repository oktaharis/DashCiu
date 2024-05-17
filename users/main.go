package main

import (
	"log"
	"net/http"

	"userscontroller/models"
	monitoring "userscontroller/monitoringuser"
	"userscontroller/users"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	models.ConnectDatabase()

	// Menggunakan variabel app dalam URL dan membiarkan page sebagai query parameter
	r.HandleFunc("/users", users.Users).Methods("GET")
	r.HandleFunc("/monitoring", monitoring.IndexMonuser).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", r))
}

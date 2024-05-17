package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"userscontroller/models"
	"userscontroller/users"
)

func main() {
	r := mux.NewRouter()
	models.ConnectDatabase("users")

	// Menggunakan variabel app dalam URL dan membiarkan page sebagai query parameter
	r.HandleFunc("/users", users.Users).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", r))
}

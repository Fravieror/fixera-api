package main

import (
	"fixera-api/service"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// Init router
	r := mux.NewRouter()

	serv := service.NewUserService()
	// Route handles & endpoints
	r.HandleFunc("/user", serv.GetUsers).Methods("GET")
	r.HandleFunc("/user/{name}", serv.GetUser).Methods("GET")
	r.HandleFunc("/user", serv.CreateUser).Methods("POST")
	r.HandleFunc("/user/{name}", serv.UpdateUser).Methods("PUT")
	r.HandleFunc("/user/{name}", serv.DeleteUser).Methods("DELETE")

	// Start server
	log.Fatal(http.ListenAndServe(":4200", r))
}

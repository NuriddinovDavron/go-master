package main

import (
	storage "go-master/controllers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	route := mux.NewRouter()
	s := route.PathPrefix("/api").Subrouter() //Base Path

	//Routes

	s.HandleFunc("/createProfile", storage.CreateProfile).Methods("POST")
	s.HandleFunc("/getAllUsers",storage.GetAllUsers).Methods("GET")
	s.HandleFunc("/getUserProfile",storage.GetUserProfile).Methods("GET")
	s.HandleFunc("/updateProfile",storage.UpdateProfile).Methods("PUT")
	s.HandleFunc("/deleteProfile/{id}",storage.DeleteProfile).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", s)) // Run Server
}

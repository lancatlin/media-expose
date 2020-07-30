package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
	r := mux.NewRouter()
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	r.HandleFunc("/", page("index"))

	r.HandleFunc("/media", newMedia).Methods("POST")
	r.HandleFunc("/media/{id}", getMedia).Methods("GET")

	r.HandleFunc("/companies", NewCompany).Methods("POST")
	r.HandleFunc("/companies/{id}", GetCompany).Methods("GET")
	return r
}

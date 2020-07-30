package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
	r := mux.NewRouter()
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	// Pages
	r.HandleFunc("/", page("index"))
	r.HandleFunc("/new", page("new"))

	r.HandleFunc("/media", MediaPage).Methods("GET")

	r.HandleFunc("/companies", CompaniesPage).Methods("GET")

	// API
	r.HandleFunc("/api/media", newMedia).Methods("POST")
	r.HandleFunc("/api/media/{id}", getMedia).Methods("GET")

	r.HandleFunc("/api/companies", NewCompany).Methods("POST")
	r.HandleFunc("/api/companies/{id}", GetCompany).Methods("GET")
	r.HandleFunc("/api/search", searchMedia).Methods("GET")
	return r
}

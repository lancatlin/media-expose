package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
	r := mux.NewRouter()
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	// Pages
	r.HandleFunc("/", Index)
	r.HandleFunc("/new", page("new"))

	r.HandleFunc("/media", MediaPage).Methods("GET")

	// API
	r.HandleFunc("/api/media", searchMedia).Methods("GET")
	r.HandleFunc("/api/media", newMedia).Methods("POST")
	r.HandleFunc("/api/media/{id}", getMedia).Methods("GET")

	r.HandleFunc("/api/companies", NewCompany).Methods("POST")
	r.HandleFunc("/api/companies/{id}", GetCompany).Methods("GET")
	return r
}

func Index(w http.ResponseWriter, r *http.Request) {
	var media []Media
	err := db.Limit(3).Preload("Company").Find(&media).Error
	if err != nil {
		panic(err)
	}
	HTML(w, r, "index", media)
}

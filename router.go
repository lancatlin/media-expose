package main

import (
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
	r := mux.NewRouter()
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static", preventDir(http.FileServer(http.Dir("static")))))

	// Pages
	r.HandleFunc("/", Index)
	r.HandleFunc("/new", page("new"))

	r.HandleFunc("/media", MediaPage).Methods("GET")
	r.HandleFunc("/media/{id}", MediaInfo).Methods("GET")

	api := r.PathPrefix("/api").Subrouter()
	api.Use(enableCors)
	// API
	api.HandleFunc("/media", searchMedia).Methods("GET")
	api.HandleFunc("/media", newMedia).Methods("POST")
	api.HandleFunc("/media/{id}", getMedia).Methods("GET")

	api.HandleFunc("/companies", NewCompany).Methods("POST")
	api.HandleFunc("/companies/{id}", GetCompany).Methods("GET")
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

func preventDir(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.HasSuffix(r.URL.Path, "/") {
			http.NotFound(w, r)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func enableCors(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		next.ServeHTTP(w, r)
	})
}

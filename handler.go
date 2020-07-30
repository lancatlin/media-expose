package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

var (
	db *gorm.DB

	tpls map[string]*template.Template
)

func NewRouter() *mux.Router {
	r := mux.NewRouter()
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	r.HandleFunc("/", page("index"))
	r.HandleFunc("/media", NewMedia).Methods("POST")
	r.HandleFunc("/companies", NewCompanies).Methods("POST")
	return r
}

func NewMedia(w http.ResponseWriter, r *http.Request) {
	dec := json.NewDecoder(r.Body)
	var media Media
	if err := dec.Decode(&media); err != nil {
		http.Error(w, err.Error(), 500)
	}
	log.Println(media)
}

func NewCompanies(w http.ResponseWriter, r *http.Request) {
	dec := json.NewDecoder(r.Body)
	var company Company
	if err := dec.Decode(&company); err != nil {
		http.Error(w, err.Error(), 500)
	}
	log.Println(company)
	if err := company.Verify(); err != nil {
		http.Error(w, err.Error(), 403)
		return
	}
	if err := db.Create(&company).Error; err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	fmt.Fprintln(w, company.ID)
}

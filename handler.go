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

type Handler struct {
	db     *gorm.DB
	router *mux.Router

	tpls map[string]*template.Template
}

func NewHandler(db *gorm.DB) Handler {
	h := Handler{
		db:     db,
		router: mux.NewRouter(),
	}
	h.loadTemplates()
	h.router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	h.HandleFunc("/", h.page("index"))
	h.HandleFunc("/media", h.NewMedia).Methods("POST")
	h.HandleFunc("/companies", h.NewCompanies).Methods("POST")
	return h
}

func (h Handler) HandleFunc(path string, f http.HandlerFunc) *mux.Route {
	return h.router.Handle(path, f)
}

func (h Handler) NewMedia(w http.ResponseWriter, r *http.Request) {
	dec := json.NewDecoder(r.Body)
	var media Media
	if err := dec.Decode(&media); err != nil {
		http.Error(w, err.Error(), 500)
	}
	log.Println(media)
}

func (h Handler) NewCompanies(w http.ResponseWriter, r *http.Request) {
	dec := json.NewDecoder(r.Body)
	var company Company
	if err := dec.Decode(&company); err != nil {
		http.Error(w, err.Error(), 500)
	}
	log.Println(company)
	if err := h.VerifyCompany(company); err != nil {
		http.Error(w, err.Error(), 403)
		return
	}
	if err := h.db.Create(&company).Error; err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	fmt.Fprintln(w, company.ID)
}

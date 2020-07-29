package main

import (
	"encoding/json"
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
	h.router.HandleFunc("/", h.page("index"))
	h.router.HandleFunc("/media", h.NewMedia).Methods("POST")
	return h
}

func (h Handler) NewMedia(w http.ResponseWriter, r *http.Request) {
	dec := json.NewDecoder(r.Body)
	var media Media
	if err := dec.Decode(&media); err != nil {
		http.Error(w, err.Error(), 500)
	}
	log.Println(media)
}

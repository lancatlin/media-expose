package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

type Handler struct {
	db     *gorm.DB
	router *mux.Router
}

func NewHandler(db *gorm.DB) Handler {
	handler := Handler{
		db:     db,
		router: mux.NewRouter(),
	}
	handler.router.HandleFunc("/media", handler.NewMedia).Methods("POST")
	return handler
}

func (h Handler) NewMedia(w http.ResponseWriter, r *http.Request) {
	dec := json.NewDecoder(r.Body)
	var media Media
	if err := dec.Decode(&media); err != nil {
		http.Error(w, err.Error(), 500)
	}
	log.Println(media)
}

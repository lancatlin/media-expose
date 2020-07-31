package main

import (
	"encoding/json"
	"net/http"
)

func searchMedia(w http.ResponseWriter, r *http.Request) {
	tx := db.Begin()
	defer tx.Close()

	domain := r.FormValue("domain")
	if domain != "" {
		tx = tx.Where("domain = ?", domain)
	}

	var media []Media
	err := tx.Preload("Company").Find(&media).Error
	if err != nil {
		panic(err)
	}

	enc := json.NewEncoder(w)
	err = enc.Encode(media)
	if err != nil {
		panic(err)
	}
}

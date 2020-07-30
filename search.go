package main

import (
	"net/http"
)

func searchMedia(w http.ResponseWriter, r *http.Request) {
	tx := db.Begin()

	domain := r.FormValue("domain")
	if domain != "" {
		tx = tx.Where("domain = ?", domain)
	}

	var media Media
	err := tx.Preload("Company").First(&media).Error
	if NotFound(err) {
		http.Error(w, "not found", 404)
	} else if err != nil {
		panic(err)
	}

	media.JSON(w)
}

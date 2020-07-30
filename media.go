package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

var (
	ErrCompanyNotExist error = errors.New("company not exist")
)

type Media struct {
	Meta

	Domain  string
	Country string

	Company   Company
	CompanyID uint `json:"company_id"`
}

func newMedia(w http.ResponseWriter, r *http.Request) {
	dec := json.NewDecoder(r.Body)
	var media Media
	if err := dec.Decode(&media); err != nil {
		http.Error(w, err.Error(), 500)
	}
	log.Println(media)

	if err := media.Verify(); err != nil {
		http.Error(w, err.Error(), 403)
		return
	}
	if err := db.Create(&media).Error; err != nil {
		panic(err)
	}
	fmt.Fprintln(w, media.ID)
}

func (m Media) Verify() (err error) {
	err = db.Where("id = ?", m.CompanyID).First(&m.Company).Error
	if NotFound(err) {
		return ErrCompanyNotExist
	} else if err != nil {
		panic(err)
	}

	if m.duplicateName() {
		return errors.New("name already exist")
	}

	if m.domainNotExist() {
		return errors.New("domain not exist")
	}
	return nil
}

func (m Media) duplicateName() bool {
	return !gorm.IsRecordNotFoundError(db.Where("name = ?", m.Name).First(&Media{}).Error)
}

func (m Media) domainNotExist() bool {
	_, err := net.LookupIP(m.Domain)
	return err != nil
}

func getMedia(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	var media Media
	err := db.Where("id = ?", id).Preload("Company").First(&media).Error
	if NotFound(err) {
		http.Error(w, "not found", 404)
		return
	} else if err != nil {
		panic(err)
	}

	media.JSON(w)
}

func MediaPage(w http.ResponseWriter, r *http.Request) {
	var media []Media
	err := db.Preload("Company").Find(&media).Error
	if err != nil {
		panic(err)
	}
	HTML(w, r, "media", media)
}

func (m Media) JSON(w io.Writer) {
	enc := json.NewEncoder(w)
	if err := enc.Encode(m); err != nil {
		panic(err)
	}
}

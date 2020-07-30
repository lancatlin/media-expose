package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

type Meta struct {
	ID        uint
	CreatedAt time.Time
	Name      string
	Source    string
	Note      string
}

type Company struct {
	Meta
	Country         string
	Owner           string
	Shareholders    string
	InvestedByChina bool `json:"invested_by_china"`
}

func NewCompany(w http.ResponseWriter, r *http.Request) {
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

func GetCompany(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	var company Company
	err := db.Where("id = ?", id).First(&company).Error
	if NotFound(err) {
		http.Error(w, "not found", 404)
		return
	} else if err != nil {
		panic(err)
	}
	company.Marshal(w)
}

func (c Company) Marshal(w io.Writer) {
	enc := json.NewEncoder(w)
	if err := enc.Encode(c); err != nil {
		panic(err)
	}
}

func NotFound(err error) bool {
	return gorm.IsRecordNotFoundError(err)
}

func (c Company) Verify() (err error) {
	if c.ID != 0 {
		return errors.New("id should be 0")
	}

	if c.duplicateName() {
		return errors.New("company name already exist")
	}

	if c.TooLong() {
		return errors.New("some data is too long")
	}

	if c.IsEmpty() {
		return errors.New("some required data is empty")
	}
	return nil
}

func (c Company) duplicateName() bool {
	return !gorm.IsRecordNotFoundError(db.Where("name = ?", c.Name).First(&Company{}).Error)
}

func (c Company) TooLong() bool {
	return c.Meta.TooLong() || len(c.Owner) > 10 || len(c.Shareholders) > 30
}

func (m Meta) TooLong() bool {
	return len(m.Name) > 30 || len(m.Source) > 200 || len(m.Note) > 200
}

func (m Meta) IsEmpty() bool {
	return len(m.Name) == 0
}

package main

import (
	"errors"
	"time"

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

type CompanyChecker struct {
	company Company
	db      *gorm.DB
}

func (h Handler) VerifyCompany(company Company) error {
	checker := CompanyChecker{company, h.db}
	return checker.Verify()
}

func (c CompanyChecker) Verify() (err error) {
	if c.company.ID != 0 {
		return errors.New("id should be 0")
	}

	if c.DuplicateName() {
		return errors.New("company name already exist")
	}

	if c.TooLong() {
		return errors.New("some data is too long")
	}

	if c.company.IsEmpty() {
		return errors.New("some required data is empty")
	}
	return nil
}

func (c CompanyChecker) DuplicateName() bool {
	return !gorm.IsRecordNotFoundError(c.db.Where("name = ?", c.company.Name).First(&Company{}).Error)
}

func (c CompanyChecker) TooLong() bool {
	return c.company.TooLong() || len(c.company.Owner) > 10 || len(c.company.Shareholders) > 30
}

func (m Meta) TooLong() bool {
	return len(m.Name) > 30 || len(m.Source) > 200 || len(m.Note) > 200
}

func (m Meta) IsEmpty() bool {
	return len(m.Name) == 0
}

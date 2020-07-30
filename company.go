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

func (c Company) Verify() (err error) {
	if c.ID != 0 {
		return errors.New("id should be 0")
	}

	if c.DuplicateName() {
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

func (c Company) DuplicateName() bool {
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

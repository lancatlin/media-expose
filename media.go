package main

import "time"

type Media struct {
	ID        uint
	CreatedAt time.Time

	Name    string
	Domain  string
	Country string
	Source  string
	Note    string

	Company   Company
	CompanyID uint `json:"company_id"`
}
